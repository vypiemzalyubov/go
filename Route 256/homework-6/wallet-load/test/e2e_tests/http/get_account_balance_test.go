//go:build httptest
// +build httptest

package http

import (
	"context"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/route256/wallet/internal/pkg/domain"
	"gitlab.ozon.dev/route256/wallet/internal/pkg/testutils"
	wallet "gitlab.ozon.dev/route256/wallet/test/utils/clients/http"
)

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

				res, resp, err := client.GetAccountBalance(ctx, &wallet.GetAccountBalanceRequest{
					UserID:    user.ID.String(),
					AccountID: account.AccountID,
				})
				require.NoError(t, err)
				require.Equal(t, http.StatusOK, resp.StatusCode)
				require.NotNil(t, resp)
				require.Equal(t, res.AccountID, account.AccountID)
				require.Equal(t, res.Amount, account.Amount)
			})
		}

	})
}
