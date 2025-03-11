//go:build grpctest
// +build grpctest

package grpc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/route256/wallet/internal/pkg/testutils"
	desc "gitlab.ozon.dev/route256/wallet/pkg/api/wallet"
)

func TestCreateAccount(t *testing.T) {
	t.Run("Positive cases", func(t *testing.T) {
		ctx, user, err := testutils.CreateUserAndAuth(context.Background(), store)
		require.NoError(t, err)

		req := &desc.CreateAccountRequest{
			UserId:      user.ID.String(),
			Amount:      100,
			Description: "some desc",
		}

		resp, err := walletClient.CreateAccount(ctx, req)
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, resp.GetAccountId())
		require.Equal(t, req.GetAmount(), resp.GetAmount())
		require.Equal(t, req.GetDescription(), resp.GetDescription())
	})
}
