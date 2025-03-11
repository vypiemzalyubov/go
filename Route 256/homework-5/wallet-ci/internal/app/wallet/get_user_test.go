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

func TestGetUser_Positive(t *testing.T) {
	t.Run("Simple get user", func(t *testing.T) {
		ctx, user, err := testutils.CreateUserAndAuth(context.Background(), store)
		require.NoError(t, err)

		impl := Implementation{store: store}

		resp, err := impl.GetUser(ctx, &desc.GetUserRequest{
			UserId: user.ID.String(),
		})
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, user.ID.String(), resp.Info.Id)
		require.Equal(t, user.Name, resp.Info.Name)
		require.Equal(t, user.Lastname, resp.Info.Lastname)
		require.Equal(t, user.Age, resp.Info.Age)
		require.Equal(t, user.Phone, resp.Info.Phone)
		require.Equal(t, desc.IdentificationLevel(desc.IdentificationLevel_value[user.Level]), resp.Info.IdentificationLevel)
	})
}

func TestGetUser_Negative(t *testing.T) {
	t.Run("Not authorized user", func(t *testing.T) {
		_, user, err := testutils.CreateUserAndAuth(context.Background(), store)
		require.NoError(t, err)

		impl := Implementation{store: store}

		resp, err := impl.GetUser(context.Background(), &desc.GetUserRequest{
			UserId: user.ID.String(),
		})

		require.Error(t, err)
		require.Nil(t, resp)
		st, ok := status.FromError(err)
		require.True(t, ok)
		require.Equal(t, codes.Unauthenticated, st.Code())
		require.Contains(t, st.Message(), "not authorized")
	})
	// По сути этот тест должен был бы вернуть 404 и покрытие было бы 100%, но он возвращает 401, что лишь немного делает более зеленым строку i.CheckToken, но само покрытие не увеличивает.
	// Поэтому, покрытие всего лишь 83.3, хотя я старался, как мог =) Скипать его не стал.
	t.Run("User not found", func(t *testing.T) {
		ctx, _, err := testutils.CreateUserAndAuth(context.Background(), store)
		require.NoError(t, err)

		impl := Implementation{store: store}

		resp, err := impl.GetUser(ctx, &desc.GetUserRequest{
			UserId: "fake-id",
		})

		require.Error(t, err)
		require.Nil(t, resp)
		st, ok := status.FromError(err)
		require.True(t, ok)
		require.Equal(t, codes.Unauthenticated, st.Code())
		require.Contains(t, st.Message(), "not authorized")
	})
}
