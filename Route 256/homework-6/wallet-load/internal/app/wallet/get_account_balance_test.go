package wallet

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

func TestGetAccountBalance_Positive(t *testing.T) {
	t.Run("Simple get account balance", func(t *testing.T) {
		for _, tt := range []struct {
			name    string
			balance int64
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
					Amount:      int32(tt.balance),
					Description: "TestCredit",
				})
				require.NoError(t, err)

				account := user.Accounts[0]

				impl := Implementation{store: store}

				balance, err := impl.GetAccountBalance(ctx, &desc.GetAccountBalanceRequest{
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
}

func TestGetAccountBalance_Negative(t *testing.T) {
	t.Run("Account doesn't exists", func(t *testing.T) {
		ctx, user, err := testutils.CreateUserAndAuth(context.Background(), store)
		require.NoError(t, err)

		impl := Implementation{store: store}

		balance, err := impl.GetAccountBalance(ctx, &desc.GetAccountBalanceRequest{
			UserId:    user.ID.String(),
			AccountId: uuid.NewString(),
		})

		require.EqualError(t, err, status.Error(codes.NotFound, "account not found").Error())
		require.Nil(t, balance)
	})
	t.Run("Not authorized user during get account balance", func(t *testing.T) {
		_, user, err := testutils.CreateUserAndAuth(context.Background(), store)
		require.NoError(t, err)

		impl := Implementation{store: store}

		balance, err := impl.GetAccountBalance(context.Background(), &desc.GetAccountBalanceRequest{
			UserId:    user.ID.String(),
			AccountId: uuid.NewString(),
		})

		require.Error(t, err)
		require.Nil(t, balance)
		st, ok := status.FromError(err)
		require.True(t, ok)
		require.Equal(t, codes.Unauthenticated, st.Code())
		require.Contains(t, st.Message(), "not authorized")
	})
	t.Run("User not found during get account balance", func(t *testing.T) {
		ctx, user, err := testutils.CreateUserAndAuth(context.Background(), store)
		require.NoError(t, err)

		faultyStore := testutils.FaultyStore{Storage: store}
		impl := Implementation{store: faultyStore}

		balance, err := impl.GetAccountBalance(ctx, &desc.GetAccountBalanceRequest{
			UserId:    user.ID.String(),
			AccountId: uuid.NewString(),
		})

		require.Error(t, err)
		require.Nil(t, balance)
		require.Contains(t, err.Error(), "user not found")
	})
}
