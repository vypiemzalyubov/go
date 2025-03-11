package wallet

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/route256/wallet/internal/pkg/testutils"
	desc "gitlab.ozon.dev/route256/wallet/pkg/api/wallet"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestLogIn_Positive(t *testing.T) {
	t.Run("Simple login user", func(t *testing.T) {
		ctxWithToken, user, err := testutils.CreateUserAndAuth(context.Background(), store)
		require.NoError(t, err)

		impl := Implementation{store: store}

		req := &desc.LogInRequest{
			Phone:    user.Phone,
			Password: user.RawPassword,
		}

		resp, err := impl.LogIn(ctxWithToken, req)
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, resp.Token)
	})
}

func TestLogIn_Negative(t *testing.T) {
	t.Run("Login with invalid phone number", func(t *testing.T) {
		ctx := context.Background()
		_, user, err := testutils.CreateUserAndAuth(ctx, store)
		require.NoError(t, err)

		impl := Implementation{store: store}

		req := &desc.LogInRequest{
			Phone:    "79992223322",
			Password: user.RawPassword,
		}

		_, err = impl.LogIn(ctx, req)
		require.Error(t, err)
		st, _ := status.FromError(err)
		require.Equal(t, codes.InvalidArgument, st.Code())
		require.Contains(t, st.Message(), "phone not match pattern 8xxxxxxxxxx")
	})
	t.Run("Login with invalid password", func(t *testing.T) {
		ctx := context.Background()
		_, user, err := testutils.CreateUserAndAuth(ctx, store)
		require.NoError(t, err)

		impl := Implementation{store: store}

		req := &desc.LogInRequest{
			Phone:    user.Phone,
			Password: "fakepass",
		}

		_, err = impl.LogIn(ctx, req)
		require.Error(t, err)
		st, _ := status.FromError(err)
		require.Equal(t, codes.InvalidArgument, st.Code())
		require.Contains(t, st.Message(), "ivalid password")
	})
	t.Run("Cannot login during login user", func(t *testing.T) {
		ctxWithToken, user, err := testutils.CreateUserAndAuth(context.Background(), store)
		require.NoError(t, err)

		faultyStore := testutils.FaultyStore{Storage: store}
		impl := Implementation{store: faultyStore}

		req := &desc.LogInRequest{
			Phone:    user.Phone,
			Password: user.RawPassword,
		}

		_, err = impl.LogIn(ctxWithToken, req)
		require.Error(t, err)
		st, _ := status.FromError(err)
		require.Equal(t, codes.Internal, st.Code())
		require.Contains(t, st.Message(), "cannot login")
	})
}
