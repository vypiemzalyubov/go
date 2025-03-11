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

// GetBalance получить инфу по балансу
func TestGetAccountBalance(t *testing.T) {
	t.Run("Positive cases", func(t *testing.T) {
		for _, tt := range []struct {
			name    string
			balance int32
		}{
			{
				name:    "positive balance",
				balance: 100,
			}, {
				name:    "zero balance",
				balance: 0,
			},
		} {
			t.Run(tt.name, func(t *testing.T) {
				ctx, user, err := testutils.CreateUserAndAuth(context.Background(), store, &domain.Account{

					AccountID:   uuid.NewString(),
					Amount:      tt.balance,
					Description: "TestCredit",
				})
				require.NoError(t, err)

				account := user.Accounts[0]

				balance, err := walletClient.GetAccountBalance(ctx, &desc.GetAccountBalanceRequest{
					UserId:    user.ID.String(),
					AccountId: account.AccountID,
				})
				require.NoError(t, err)
				require.NotNil(t, balance)
				require.Equal(t, account.AccountID, balance.AccountId)
				require.Equal(t, account.Amount, balance.Amount)
			})
		}

	})

	t.Run("Negative cases", func(t *testing.T) {
		t.Run("Account doesn't exists", func(t *testing.T) {
			ctx, user, err := testutils.CreateUserAndAuth(context.Background(), store)
			require.NoError(t, err)

			balance, err := walletClient.GetAccountBalance(ctx, &desc.GetAccountBalanceRequest{
				UserId:    user.ID.String(),
				AccountId: uuid.NewString(),
			})
			require.EqualError(t, err, status.Error(codes.NotFound, "account not found").Error())
			require.Nil(t, balance)
		})
	})
}
