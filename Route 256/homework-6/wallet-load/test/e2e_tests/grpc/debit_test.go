//go:build grpctest
// +build grpctest

package grpc

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/route256/wallet/internal/pkg/domain"
	"gitlab.ozon.dev/route256/wallet/internal/pkg/testutils"
	desc "gitlab.ozon.dev/route256/wallet/pkg/api/wallet"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestDebit(t *testing.T) {
	t.Run("Positive cases", func(t *testing.T) {
		ctx, user, err := testutils.CreateUserAndAuth(context.Background(), store, &domain.Account{
			AccountID:   uuid.NewString(),
			Amount:      100,
			Description: "TestCredit",
		})
		require.NoError(t, err)

		account := user.Accounts[0]

		debitRequest := &desc.DebitRequest{
			UserId:      user.ID.String(),
			AccountId:   account.AccountID,
			Amount:      100,
			OperationId: uuid.NewString(),
		}

		resp, err := walletClient.Debit(ctx, debitRequest)
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, desc.OperationStatus_STATUS_OK, resp.Status)

		operations, err := store.GetAccountOperations(ctx, account.AccountID, 1000)
		require.NoError(t, err)
		require.Len(t, operations, 1)

		require.Equal(t, debitRequest.Amount, operations[0].Amount)
		require.Equal(t, debitRequest.OperationId, operations[0].OperationID)
		require.Equal(t, domain.OperationTypeDebit, operations[0].OperationType)

		checkNewAccountBalance(ctx, t, walletClient, user.ID.String(), account, debitRequest.Amount, domain.OperationTypeDebit)
	})

	t.Run("Negative cases", func(t *testing.T) {
		t.Run("Account doesn't exists", func(t *testing.T) {
			ctx, user, err := testutils.CreateUserAndAuth(context.Background(), store)
			require.NoError(t, err)

			accountID := uuid.NewString()
			resp, err := walletClient.Debit(ctx, &desc.DebitRequest{
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

			debitRequest := &desc.DebitRequest{
				UserId:      user.ID.String(),
				AccountId:   account.AccountID,
				Amount:      100,
				OperationId: uuid.NewString(),
			}

			t.Log("First credit operation")

			resp, err := walletClient.Debit(ctx, debitRequest)
			require.NoError(t, err)
			require.NotNil(t, resp)
			require.Equal(t, desc.OperationStatus_STATUS_OK, resp.Status)

			t.Log("Second credit operation")

			resp, err = walletClient.Debit(ctx, debitRequest)
			require.Error(t, err)
			require.Nil(t, resp)

			operations, err := store.GetAccountOperations(ctx, account.AccountID, 1000)
			require.NoError(t, err)
			require.Len(t, operations, 1)
		})

	})
}

func checkNewAccountBalance(ctx context.Context, t *testing.T, cl desc.WalletClient, userID string,
	account *domain.Account, operationAmount int32, operationType domain.OperationType) {

	balance, err := cl.GetAccountBalance(ctx, &desc.GetAccountBalanceRequest{
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
