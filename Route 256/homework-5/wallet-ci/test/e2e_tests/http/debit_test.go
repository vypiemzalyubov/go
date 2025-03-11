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
	desc "gitlab.ozon.dev/route256/wallet/pkg/api/wallet"
	wallet "gitlab.ozon.dev/route256/wallet/test/utils/clients/http"
)

func TestDebit(t *testing.T) {
	t.Run("Positive cases", func(t *testing.T) {
		ctx, user, err := testutils.CreateUserAndAuth(context.Background(), store, &domain.Account{
			AccountID:   uuid.NewString(),
			Amount:      100,
			Description: "TestCredit",
		})
		require.NoError(t, err)

		account := user.Accounts[0]

		debitRequest := &wallet.DebitRequest{
			UserID:      user.ID.String(),
			AccountID:   account.AccountID,
			Amount:      100,
			OperationID: uuid.NewString(),
		}

		res, resp, err := client.Debit(ctx, debitRequest)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
		require.Equal(t, desc.OperationStatus_STATUS_OK.String(), res.Status)

		operations, err := store.GetAccountOperations(ctx, account.AccountID, 1000)
		require.NoError(t, err)
		require.Len(t, operations, 1)

		require.Equal(t, debitRequest.Amount, operations[0].Amount)
		require.Equal(t, debitRequest.OperationID, operations[0].OperationID)
		require.Equal(t, domain.OperationTypeDebit, operations[0].OperationType)
	})
}
