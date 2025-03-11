//go:build httptest
// +build httptest

package http

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/route256/wallet/internal/pkg/testutils"
	wallet "gitlab.ozon.dev/route256/wallet/test/utils/clients/http"
)

func TestCreateAccount(t *testing.T) {
	t.Run("Positive cases", func(t *testing.T) {
		ctx, user, err := testutils.CreateUserAndAuth(context.Background(), store)
		require.NoError(t, err)

		req := &wallet.CreateAccountRequest{
			UserID:      user.ID.String(),
			Amount:      100,
			Description: "some desc",
		}

		res, resp, err := client.CreateAccount(ctx, req)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
		require.NotNil(t, resp)
		require.NotEmpty(t, res.AccountID)
		require.Equal(t, res.Amount, req.Amount)
		require.Equal(t, res.Description, req.Description)
	})
}
