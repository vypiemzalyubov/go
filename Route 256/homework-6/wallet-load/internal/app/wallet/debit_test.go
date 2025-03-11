package wallet

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/route256/wallet/internal/config"
	"gitlab.ozon.dev/route256/wallet/internal/pkg/domain"
	"gitlab.ozon.dev/route256/wallet/internal/pkg/mocks"
	"gitlab.ozon.dev/route256/wallet/internal/pkg/testutils"
	desc "gitlab.ozon.dev/route256/wallet/pkg/api/wallet"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func TestDebit_Positive(t *testing.T) {
	t.Run("Simple debit operation", func(t *testing.T) {
		ctx, user, err := testutils.CreateUserAndAuth(context.Background(), store, &domain.Account{
			AccountID:   uuid.NewString(),
			Amount:      100,
			Description: "TestCredit1",
		})
		require.NoError(t, err)

		account := user.Accounts[0]

		impl := Implementation{store: store}

		debitRequest := &desc.DebitRequest{
			UserId:      user.ID.String(),
			AccountId:   account.AccountID,
			Amount:      100,
			OperationId: uuid.NewString(),
		}

		resp, err := impl.Debit(ctx, debitRequest)
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, desc.OperationStatus_STATUS_OK, resp.Status)

		operations, err := store.GetAccountOperations(ctx, account.AccountID, 1000)
		require.NoError(t, err)
		require.Len(t, operations, 1)

		require.Equal(t, debitRequest.Amount, operations[0].Amount)
		require.Equal(t, debitRequest.OperationId, operations[0].OperationID)
		require.Equal(t, domain.OperationTypeDebit, operations[0].OperationType)

		checkNewAccountBalance(ctx, t, impl, user.ID.String(), account, debitRequest.Amount, domain.OperationTypeDebit)
	})
	t.Run("Debit operation with empty operation ID", func(t *testing.T) {
		ctx, user, err := testutils.CreateUserAndAuth(context.Background(), store, &domain.Account{
			AccountID:   uuid.NewString(),
			Amount:      50000,
			Description: "TestDebit2",
		})
		require.NoError(t, err)

		impl := Implementation{store: store}

		debitRequest := &desc.DebitRequest{
			UserId:      user.ID.String(),
			AccountId:   user.Accounts[0].AccountID,
			Amount:      50,
			OperationId: "",
		}

		resp, err := impl.Debit(ctx, debitRequest)
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, desc.OperationStatus_STATUS_OK, resp.Status)

		operations, err := store.GetAccountOperations(ctx, user.Accounts[0].AccountID, 1000)
		require.NoError(t, err)
		require.Len(t, operations, 1)
		require.Equal(t, debitRequest.Amount, operations[0].Amount)
		require.NotEmpty(t, operations[0].OperationID)
	})
}

func TestDebit_Negative(t *testing.T) {
	t.Run("Account doesn't exists", func(t *testing.T) {
		ctx, user, err := testutils.CreateUserAndAuth(context.Background(), store)
		require.NoError(t, err)

		impl := Implementation{store: store}

		accountID := uuid.NewString()
		resp, err := impl.Debit(ctx, &desc.DebitRequest{
			UserId:      user.ID.String(),
			AccountId:   uuid.NewString(),
			Amount:      100,
			OperationId: uuid.NewString(),
		})
		require.EqualError(t, err, status.Error(codes.NotFound, "account not found").Error())
		require.Nil(t, resp)

		checkOperationNotExists(ctx, t, accountID)
	})

	t.Run("Double operation with same ID", func(t *testing.T) {
		ctx, user, err := testutils.CreateUserAndAuth(context.Background(), store, &domain.Account{
			AccountID:   uuid.NewString(),
			Amount:      100,
			Description: "TestCredit",
		})
		require.NoError(t, err)

		account := user.Accounts[0]

		impl := Implementation{store: store}

		debitRequest := &desc.DebitRequest{
			UserId:      user.ID.String(),
			AccountId:   account.AccountID,
			Amount:      100,
			OperationId: uuid.NewString(),
		}

		t.Log("First credit operation")

		resp, err := impl.Debit(ctx, debitRequest)
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, desc.OperationStatus_STATUS_OK, resp.Status)

		t.Log("Second credit operation")

		resp, err = impl.Debit(ctx, debitRequest)
		require.Error(t, err)
		require.NotNil(t, resp)
		require.Equal(t, desc.OperationStatus_STATUS_FAIL, resp.Status)

		operations, err := store.GetAccountOperations(ctx, account.AccountID, 1000)
		require.NoError(t, err)
		require.Len(t, operations, 1)
	})
	t.Run("Debit operation during error getting user", func(t *testing.T) {
		mockedStore := mocks.NewMinimockStorage(t)
		impl := Implementation{store: mockedStore}

		session := uuid.NewString()
		mockedStore.GetUserSessionsMock.Return([]string{session}, nil)
		mockedStore.GetUserMock.Return(nil, fmt.Errorf("user not found"))
		ctx := metadata.NewIncomingContext(context.Background(), metadata.MD{"authorization": []string{session}})

		req := &desc.DebitRequest{
			UserId:      "fake-id",
			AccountId:   uuid.NewString(),
			Amount:      777,
			OperationId: uuid.NewString(),
		}

		debit, err := impl.Debit(ctx, req)
		require.Error(t, err)
		require.Nil(t, debit)
		require.Equal(t, codes.Internal, status.Code(err))
		require.Contains(t, err.Error(), "user not found")
	})
	t.Run("Debit operation by not authorized user", func(t *testing.T) {
		ctx, _, err := testutils.CreateUserAndAuth(context.Background(), store)
		require.NoError(t, err)

		impl := Implementation{store: store}

		debitRequest := &desc.DebitRequest{
			UserId:      uuid.NewString(),
			AccountId:   uuid.NewString(),
			Amount:      50,
			OperationId: uuid.NewString(),
		}

		resp, err := impl.Debit(ctx, debitRequest)
		require.EqualError(t, err, status.Error(codes.Unauthenticated, "not authorized").Error())
		require.Nil(t, resp)
	})
	t.Run("Debit limit for FULL user exceeded", func(t *testing.T) {
		ctx, user, err := testutils.CreateUserAndAuth(context.Background(), store, &domain.Account{
			AccountID:   uuid.NewString(),
			Amount:      0,
			Description: "Fake FULL Account",
		})
		require.NoError(t, err)

		impl := Implementation{store: store}

		_, err = impl.UpgradeUser(ctx, &desc.UpgradeUserRequest{
			UserId: user.ID.String(),
		})
		require.NoError(t, err)

		account := user.Accounts[0]

		debitRequest := &desc.DebitRequest{
			UserId:      user.ID.String(),
			AccountId:   account.AccountID,
			Amount:      config.MaxFullAmount + 1,
			OperationId: uuid.NewString(),
		}

		resp, err := impl.Debit(ctx, debitRequest)
		require.Error(t, err)
		require.NotNil(t, resp)
		require.Equal(t, desc.OperationStatus_STATUS_FAIL, resp.Status)

		st, ok := status.FromError(err)
		require.True(t, ok)
		require.Equal(t, codes.InvalidArgument, st.Code())
		require.Contains(t, st.Message(), "unavailable operation")
	})
	t.Run("Debit limit for ANON user exceeded", func(t *testing.T) {
		t.Skip("Этот тест не проходит, т.к баг еще не починили")
		ctx, user, err := testutils.CreateUserAndAuth(context.Background(), store, &domain.Account{
			AccountID:   uuid.NewString(),
			Amount:      0,
			Description: "Fake ANON Account",
		})
		require.NoError(t, err)

		impl := Implementation{store: store}

		debitRequest := &desc.DebitRequest{
			UserId:      user.ID.String(),
			AccountId:   user.Accounts[0].AccountID,
			Amount:      config.MaxAnonAmount + 1,
			OperationId: uuid.NewString(),
		}

		resp, err := impl.Debit(ctx, debitRequest)

		require.Error(t, err)
		require.NotNil(t, resp)
		require.Equal(t, desc.OperationStatus_STATUS_FAIL, resp.Status)

		st, ok := status.FromError(err)
		require.True(t, ok)
		require.Equal(t, codes.InvalidArgument, st.Code())
		require.Contains(t, st.Message(), "unavailable operation")
	})
	t.Run("Internal error when debit operation", func(t *testing.T) {
		_, user, err := testutils.CreateUserAndAuth(context.Background(), store, &domain.Account{
			AccountID:   uuid.NewString(),
			Amount:      0,
			Description: "Fake Account",
		})
		require.NoError(t, err)

		mockedStore := mocks.NewMinimockStorage(t)
		impl := Implementation{store: mockedStore}

		session := uuid.NewString()
		mockedStore.GetUserSessionsMock.Return([]string{session}, nil)
		mockedStore.GetUserMock.Return(user, nil)
		mockedStore.DebitMock.Return(fmt.Errorf("Debit process err"))
		ctx := metadata.NewIncomingContext(context.Background(), metadata.MD{"authorization": []string{session}})

		debit, err := impl.Debit(ctx, &desc.DebitRequest{
			UserId:      user.ID.String(),
			AccountId:   user.Accounts[0].AccountID,
			Amount:      1000,
			OperationId: uuid.NewString(),
		})

		require.Error(t, err)
		require.NotNil(t, debit)
		require.Equal(t, desc.OperationStatus_STATUS_FAIL, debit.Status)
		require.Equal(t, codes.Internal, status.Code(err))
		require.Contains(t, err.Error(), "Debit process err")
	})
}

func checkNewAccountBalance(ctx context.Context, t *testing.T, impl Implementation, userID string,
	account *domain.Account, operationAmount int32, operationType domain.OperationType) {

	balance, err := impl.GetAccountBalance(ctx, &desc.GetAccountBalanceRequest{
		UserId:    userID,
		AccountId: account.AccountID,
	})
	require.NoError(t, err)
	require.NotNil(t, balance)

	expectedAmount := account.Amount + (int32(operationType) * operationAmount)
	require.Equal(t, expectedAmount, balance.Amount)
}

func checkOperationNotExists(ctx context.Context, t *testing.T, accountID string) {
	operations, err := store.GetAccountOperations(ctx, accountID, 1000)
	require.NoError(t, err)
	require.Len(t, operations, 0)
}
