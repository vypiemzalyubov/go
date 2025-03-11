package wallet

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/go-jet/jet/v2/postgres"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/route256/wallet/internal/pkg/domain"
	"gitlab.ozon.dev/route256/wallet/internal/pkg/mocks"
	"gitlab.ozon.dev/route256/wallet/internal/pkg/storage/gen/wallet/public/table"
	"gitlab.ozon.dev/route256/wallet/internal/pkg/testutils"
	desc "gitlab.ozon.dev/route256/wallet/pkg/api/wallet"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func TestUpgradeUser_Positive(t *testing.T) {
	t.Run("Simple upgrade user", func(t *testing.T) {
		ctx, user, err := testutils.CreateUserAndAuth(context.Background(), store)
		require.NoError(t, err)

		impl := Implementation{store: store}

		req := &desc.UpgradeUserRequest{
			UserId: user.ID.String(),
		}

		_, err = impl.UpgradeUser(ctx, req)
		require.NoError(t, err)
	})
}

func TestUpgradeUser_Negative(t *testing.T) {
	t.Run("Upgrade not authorized user", func(t *testing.T) {
		_, user, err := testutils.CreateUserAndAuth(context.Background(), store)
		require.NoError(t, err)

		impl := Implementation{store: store}

		req := &desc.UpgradeUserRequest{
			UserId: user.ID.String(),
		}

		resp, err := impl.UpgradeUser(context.Background(), req)
		require.Error(t, err)
		st, _ := status.FromError(err)
		require.Equal(t, codes.Unauthenticated, st.Code())
		require.Contains(t, st.Message(), "not authorized")
		require.Nil(t, resp)
	})
	t.Run("Upgrade an already upgraded user", func(t *testing.T) {
		ctx, user, err := testutils.CreateUserAndAuth(context.Background(), store)
		require.NoError(t, err)

		impl := Implementation{store: store}

		req := &desc.UpgradeUserRequest{
			UserId: user.ID.String(),
		}

		_, err = impl.UpgradeUser(ctx, req)
		require.NoError(t, err)

		req = &desc.UpgradeUserRequest{
			UserId: user.ID.String(),
		}

		resp, err := impl.UpgradeUser(ctx, req)
		require.Error(t, err)
		st, _ := status.FromError(err)
		require.Equal(t, codes.AlreadyExists, st.Code())
		require.Contains(t, st.Message(), "user allready has FULL level")
		require.Nil(t, resp)
	})
	t.Run("Internal error during upgrade user", func(t *testing.T) {
		_, user, err := testutils.CreateUserAndAuth(context.Background(), store)
		require.NoError(t, err)

		mockedStore := mocks.NewMinimockStorage(t)
		impl := Implementation{store: mockedStore}

		session := uuid.NewString()
		mockedStore.GetUserSessionsMock.Return([]string{session}, nil)
		mockedStore.GetUserMock.Return(&domain.UserWithAccounts{}, nil)
		mockedStore.UpgradeUserMock.Return(fmt.Errorf("UpgradeUser err"))
		ctx := metadata.NewIncomingContext(context.Background(), metadata.MD{"authorization": []string{session}})

		req := &desc.UpgradeUserRequest{
			UserId: user.ID.String(),
		}

		_, err = impl.UpgradeUser(ctx, req)
		require.Error(t, err)
		require.Equal(t, codes.Internal, status.Code(err))
		require.Contains(t, err.Error(), "UpgradeUser err")
	})
}

func TestUpgradeUser_OldUser(t *testing.T) {
	t.Run("Upgrade old user", func(t *testing.T) {
		ctx, user, err := testutils.CreateUserAndAuth(context.Background(), store)
		require.NoError(t, err)

		createdAt := user.CreatedAt
		newCreatedAt := time.Now().AddDate(-1, 0, 0)
		require.NotEqual(t, createdAt, newCreatedAt, "User creation dates are not the same")

		year, month, day := newCreatedAt.Date()
		hour, minute, second := newCreatedAt.Clock()
		_, offset := newCreatedAt.Zone()
		zoneDuration := time.Duration(offset) * time.Second

		stmt := table.Users.
			UPDATE().
			SET(table.Users.CreatedAt.SET(postgres.Timestampz(year, month, day, hour, minute, second, zoneDuration, "UTC"))).
			WHERE(table.Users.ID.EQ(postgres.UUID(user.ID)))

		query, args := stmt.Sql()

		_, err = store.Exec(ctx, query, args...)
		require.NoError(t, err)

		impl := &Implementation{store: store}

		req := &desc.UpgradeUserRequest{
			UserId: user.ID.String(),
		}

		resp, err := impl.UpgradeUser(ctx, req)
		require.NoError(t, err)
		require.NotNil(t, resp)

		updatedUser, err := store.GetUser(ctx, user.ID.String())
		require.NoError(t, err)
		require.Equal(t, "FULL", updatedUser.Level)
	})
}
