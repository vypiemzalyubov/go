package wallet

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/route256/wallet/internal/pkg/domain"
	"gitlab.ozon.dev/route256/wallet/internal/pkg/mocks"
	"gitlab.ozon.dev/route256/wallet/internal/pkg/storage"
	"gitlab.ozon.dev/route256/wallet/internal/pkg/testutils"
	desc "gitlab.ozon.dev/route256/wallet/pkg/api/wallet"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func TestCreateAccount_Positive(t *testing.T) {
	t.Run("Simple create account", func(t *testing.T) {
		impl := Implementation{store: store}

		ctx, user, err := testutils.CreateUserAndAuth(context.Background(), store)
		require.NoError(t, err)

		req := &desc.CreateAccountRequest{
			UserId:      user.ID.String(),
			Amount:      100,
			Description: "some desc 1",
		}

		resp, err := impl.CreateAccount(ctx, req)
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, resp.GetAccountId())
		require.Equal(t, req.GetAmount(), resp.GetAmount())
		require.Equal(t, req.GetDescription(), resp.GetDescription())
	})
	t.Run("Create account with max amount", func(t *testing.T) {
		impl := Implementation{store: store}

		ctx, user, err := testutils.CreateUserAndAuth(context.Background(), store)
		require.NoError(t, err)

		req := &desc.CreateAccountRequest{
			UserId:      user.ID.String(),
			Amount:      50000,
			Description: "some desc 2",
		}

		resp, err := impl.CreateAccount(ctx, req)
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, resp.GetAccountId())
		require.Equal(t, req.GetAmount(), resp.GetAmount())
		require.Equal(t, req.GetDescription(), resp.GetDescription())
	})
}

func TestCreateAccount_Negative(t *testing.T) {
	t.Run("Error fetching user", func(t *testing.T) {
		impl := Implementation{store: store}

		req := &desc.CreateAccountRequest{
			UserId:      "fake-userid",
			Amount:      300,
			Description: "some fake desc",
		}

		_, err := impl.CreateAccount(context.Background(), req)
		require.Error(t, err)
		require.Equal(t, codes.Unauthenticated, status.Code(err))
		require.Contains(t, err.Error(), "not authorized")
	})
	t.Run("User already has 3 accounts", func(t *testing.T) {
		impl := Implementation{store: store}

		ctx, user, err := testutils.CreateUserAndAuth(context.Background(), store)
		require.NoError(t, err)

		err = createThreeAccounts(ctx, store, user)
		require.NoError(t, err)

		req := &desc.CreateAccountRequest{
			UserId:      user.ID.String(),
			Amount:      100,
			Description: "some desc",
		}

		_, err = impl.CreateAccount(ctx, req)
		require.Error(t, err)
		require.Equal(t, codes.AlreadyExists, status.Code(err))
		require.Contains(t, err.Error(), "User allready has 3 accounts")
	})
	t.Run("Internal error when adding account", func(t *testing.T) {
		_, user, err := testutils.CreateUserAndAuth(context.Background(), store)
		require.NoError(t, err)

		mockedStore := mocks.NewMinimockStorage(t)
		impl := Implementation{store: mockedStore}

		session := uuid.NewString()
		mockedStore.GetUserSessionsMock.Return([]string{session}, nil)
		mockedStore.GetUserMock.Return(&domain.UserWithAccounts{}, nil)
		mockedStore.AddAccountMock.Return(fmt.Errorf("Create account err"))
		ctx := metadata.NewIncomingContext(context.Background(), metadata.MD{"authorization": []string{session}})

		req := &desc.CreateAccountRequest{
			UserId:      user.ID.String(),
			Amount:      1000,
			Description: "fake account",
		}

		_, err = impl.CreateAccount(ctx, req)
		require.Error(t, err)
		require.Equal(t, codes.Internal, status.Code(err))
		require.Contains(t, err.Error(), "Create account err")
	})
}

func createThreeAccounts(ctx context.Context, store storage.Storage, user *domain.UserWithAccounts) error {
	for i := 0; i < 3; i++ {
		acc := &domain.Account{
			UserID:    user.ID,
			AccountID: uuid.NewString(),
			Amount:    100,
		}
		err := store.AddAccount(ctx, acc)
		if err != nil {
			return err
		}
	}
	return nil
}
