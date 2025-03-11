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

func TestCredit(t *testing.T) {
	t.Run("Positive cases", func(t *testing.T) {
		ctx, user, err := testutils.CreateUserAndAuth(context.Background(), store, &domain.Account{
			AccountID:   uuid.NewString(),
			Amount:      100,
			Description: "TestCredit",
		})
		require.NoError(t, err)

		account := user.Accounts[0]

		creditRequest := &desc.CreditRequest{
			UserId:      user.ID.String(),
			AccountId:   account.AccountID,
			Amount:      100,
			OperationId: uuid.NewString(),
		}

		resp, err := walletClient.Credit(ctx, creditRequest)
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, desc.OperationStatus_STATUS_OK, resp.Status)

		operations, err := store.GetAccountOperations(ctx, account.AccountID, 1000)
		require.NoError(t, err)
		require.Len(t, operations, 1)

		require.Equal(t, creditRequest.Amount, operations[0].Amount)
		require.Equal(t, creditRequest.OperationId, operations[0].OperationID)
		require.Equal(t, domain.OperationTypeCredit, operations[0].OperationType)

		checkNewAccountBalance(ctx, t, walletClient, user.ID.String(), account, creditRequest.Amount, domain.OperationTypeCredit)
	})

	t.Run("Negative cases", func(t *testing.T) {
		t.Run("Account doesn't exists", func(t *testing.T) {
			ctx, user, err := testutils.CreateUserAndAuth(context.Background(), store)
			require.NoError(t, err)

			accountID := uuid.NewString()
			resp, err := walletClient.Credit(ctx, &desc.CreditRequest{
				UserId:      user.ID.String(),
				AccountId:   accountID,
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

			creditRequest := &desc.CreditRequest{
				UserId:      user.ID.String(),
				AccountId:   account.AccountID,
				Amount:      100,
				OperationId: uuid.NewString(),
			}

			t.Log("First credit operation")

			resp, err := walletClient.Credit(ctx, creditRequest)
			require.NoError(t, err)
			require.NotNil(t, resp)
			require.Equal(t, desc.OperationStatus_STATUS_OK, resp.Status)

			t.Log("Second credit operation")

			resp, err = walletClient.Credit(ctx, creditRequest)

			require.Error(t, err)
			require.Nil(t, resp)

			operations, err := store.GetAccountOperations(ctx, account.AccountID, 1000)
			require.NoError(t, err)
			require.Len(t, operations, 1)
		})
	})
}
