package wallet

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/route256/wallet/internal/pkg/worker"
	"gitlab.ozon.dev/route256/wallet/pkg/api/wallet"
)

type MockWorkerCollectOps struct {
	Count int64
}

func (m *MockWorkerCollectOps) Load() int64 {
	return m.Count
}

func TestCollectOperation(t *testing.T) {
	t.Run("Simple collect operation", func(t *testing.T) {
		worker.Count.Store(10)

		impl := &Implementation{}

		ctx := context.Background()
		req := &wallet.OperationCountRequest{}

		resp, err := impl.CollectOperation(ctx, req)

		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, int64(10), resp.OperationsCount)
	})
}
