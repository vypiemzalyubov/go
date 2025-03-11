package wallet

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/route256/wallet/internal/pkg/domain"
	"gitlab.ozon.dev/route256/wallet/internal/pkg/mocks"
	"gitlab.ozon.dev/route256/wallet/internal/pkg/testutils"
	desc "gitlab.ozon.dev/route256/wallet/pkg/api/wallet"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func TestCredit_Positive(t *testing.T) {
	t.Run("Simple credit operation", func(t *testing.T) {
		ctx, user, err := testutils.CreateUserAndAuth(context.Background(), store, &domain.Account{
			AccountID:   uuid.NewString(),
			Amount:      100,
			Description: "TestCredit",
		})
		require.NoError(t, err)

		account := user.Accounts[0]

		impl := Implementation{store: store}

		creditRequest := &desc.CreditRequest{
			UserId:      user.ID.String(),
			AccountId:   account.AccountID,
			Amount:      100,
			OperationId: uuid.NewString(),
		}

		resp, err := impl.Credit(ctx, creditRequest)
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, desc.OperationStatus_STATUS_OK, resp.Status)

		operations, err := store.GetAccountOperations(ctx, account.AccountID, 1000)
		require.NoError(t, err)
		require.Len(t, operations, 1)

		require.Equal(t, creditRequest.Amount, operations[0].Amount)
		require.Equal(t, creditRequest.OperationId, operations[0].OperationID)
		require.Equal(t, domain.OperationTypeCredit, operations[0].OperationType)

		checkNewAccountBalance(ctx, t, impl, user.ID.String(), account, creditRequest.Amount, domain.OperationTypeCredit)
	})
}

func TestCredit_Negative(t *testing.T) {
	t.Run("Account doesn't exists", func(t *testing.T) {
		ctx, user, err := testutils.CreateUserAndAuth(context.Background(), store)
		require.NoError(t, err)

		impl := Implementation{store: store}

		accountID := uuid.NewString()

		resp, err := impl.Credit(ctx, &desc.CreditRequest{
			UserId:      user.ID.String(),
			AccountId:   uuid.NewString(),
			Amount:      100,
			OperationId: uuid.NewString(),
		})
		require.EqualError(t, err, status.Error(codes.NotFound, "account not found").Error())
		require.Nil(t, resp)

		checkOperationNotExists(ctx, t, accountID)
	})
	t.Run("Credit operation for not authorized user", func(t *testing.T) {
		ctx, user, err := testutils.CreateUserAndAuth(context.Background(), store)
		require.NoError(t, err)

		impl := Implementation{store: store}

		accountID := uuid.NewString()

		resp, err := impl.Credit(context.Background(), &desc.CreditRequest{
			UserId:      user.ID.String(),
			AccountId:   uuid.NewString(),
			Amount:      300,
			OperationId: uuid.NewString(),
		})
		require.EqualError(t, err, status.Error(codes.Unauthenticated, "not authorized").Error())
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

		creditRequest := &desc.CreditRequest{
			UserId:      user.ID.String(),
			AccountId:   account.AccountID,
			Amount:      100,
			OperationId: uuid.NewString(),
		}

		t.Log("First credit operation")

		resp, err := impl.Credit(ctx, creditRequest)
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, desc.OperationStatus_STATUS_OK, resp.Status)

		t.Log("Second credit operation")

		resp, err = impl.Credit(ctx, creditRequest)
		require.Error(t, err)
		require.NotNil(t, resp)
		require.Equal(t, desc.OperationStatus_STATUS_FAIL, resp.Status)

		operations, err := store.GetAccountOperations(ctx, account.AccountID, 1000)
		require.NoError(t, err)
		require.Len(t, operations, 1)
	})
	t.Run("Credit too much money", func(t *testing.T) {
		ctx, user, err := testutils.CreateUserAndAuth(context.Background(), store, &domain.Account{
			AccountID:   uuid.NewString(),
			Amount:      0,
			Description: "Fake too much money",
		})
		require.NoError(t, err)

		account := user.Accounts[0]

		impl := Implementation{store: store}

		creditRequest := &desc.CreditRequest{
			UserId:      user.ID.String(),
			AccountId:   account.AccountID,
			Amount:      1,
			OperationId: uuid.NewString(),
		}

		resp, err := impl.Credit(ctx, creditRequest)
		require.Error(t, err)
		require.EqualError(t, err, status.Error(codes.InvalidArgument, "unavailable operation").Error())

		require.NotNil(t, resp)
		require.Equal(t, desc.OperationStatus_STATUS_FAIL, resp.Status)

		operations, err := store.GetAccountOperations(ctx, account.AccountID, 1000)
		require.NoError(t, err)
		require.Len(t, operations, 0)

		checkNewAccountBalance(ctx, t, impl, user.ID.String(), account, 0, domain.OperationTypeCredit)
	})
	t.Run("Automatically generate OperationId", func(t *testing.T) {
		ctx, user, err := testutils.CreateUserAndAuth(context.Background(), store, &domain.Account{
			AccountID:   uuid.NewString(),
			Amount:      111,
			Description: "Fake OperationID",
		})
		require.NoError(t, err)

		account := user.Accounts[0]

		impl := Implementation{store: store}

		creditRequest := &desc.CreditRequest{
			UserId:    user.ID.String(),
			AccountId: account.AccountID,
			Amount:    50,
		}

		resp, err := impl.Credit(ctx, creditRequest)
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, desc.OperationStatus_STATUS_OK, resp.Status)

		operations, err := store.GetAccountOperations(ctx, account.AccountID, 1000)
		require.NoError(t, err)
		require.Len(t, operations, 1)

		require.NotEmpty(t, operations[0].OperationID)
		require.Equal(t, creditRequest.Amount, operations[0].Amount)
		require.Equal(t, domain.OperationTypeCredit, operations[0].OperationType)
	})
	t.Run("Internal error when credit operation", func(t *testing.T) {
		_, user, err := testutils.CreateUserAndAuth(context.Background(), store, &domain.Account{
			AccountID:   uuid.NewString(),
			Amount:      50000,
			Description: "Fake desc",
		})
		require.NoError(t, err)

		account := user.Accounts[0]

		mockedStore := mocks.NewMinimockStorage(t)
		impl := Implementation{store: mockedStore}

		session := uuid.NewString()
		mockedStore.GetUserSessionsMock.Return([]string{session}, nil)
		mockedStore.GetUserMock.Return(user, nil)
		mockedStore.CreditMock.Return(fmt.Errorf("Credit process err"))

		ctx := metadata.NewIncomingContext(context.Background(), metadata.MD{"authorization": []string{session}})

		creditRequest := &desc.CreditRequest{
			UserId:      user.ID.String(),
			AccountId:   account.AccountID,
			Amount:      5,
			OperationId: uuid.NewString(),
		}

		resp, err := impl.Credit(ctx, creditRequest)
		require.Error(t, err)
		require.NotNil(t, resp)
		require.Equal(t, desc.OperationStatus_STATUS_FAIL, resp.Status)
		require.Equal(t, codes.Internal, status.Code(err))
		require.Contains(t, err.Error(), "Credit process err")
	})
}
