package grpc

// import (
// 	"context"
// 	"e2e-backend/internal/clients/grpccli"
// 	desc "e2e-backend/internal/pb/gitlab.ozon.dev/route256/wallet"
// 	"github.com/stretchr/testify/require"
// 	"testing"
// 	"time"
// )

// func TestWorker(t *testing.T) {
// 	ctx := context.Background()

// 	wallet := grpccli.NewWalletClient()

// 	t.Run("custom polling", func(t *testing.T) {
// 		resp, err := wallet.CollectOperation(ctx, &desc.OperationCountRequest{})
// 		require.NoError(t, err)
// 		require.NotNil(t, resp)
// 		t.Logf("Operations count: %d", resp.GetOperationsCount())

// 		makeCredit(ctx, t, wallet)

// 		ok := polling(ctx, t, wallet, resp.GetOperationsCount()+1)
// 		require.True(t, ok)
// 	})

// 	t.Run("testify polling", func(t *testing.T) {
// 		resp, err := wallet.CollectOperation(ctx, &desc.OperationCountRequest{})
// 		require.NoError(t, err)
// 		require.NotNil(t, resp)
// 		t.Logf("Operations count: %d", resp.GetOperationsCount())

// 		makeCredit(ctx, t, wallet)

// 		expectedOperations := resp.GetOperationsCount() + 1
// 		waitOperations(ctx, t, wallet, expectedOperations)
// 	})

// 	t.Run("ForceCollectOperation", func(t *testing.T) {

// 		respBefore, err := wallet.CollectOperation(ctx, &desc.OperationCountRequest{})
// 		require.NoError(t, err)
// 		require.NotNil(t, respBefore)
// 		t.Logf("Operations count: %d", respBefore.GetOperationsCount())

// 		makeCredit(ctx, t, wallet)

// 		respForce, err := wallet.ForceCollectOperation(ctx, &desc.OperationCountRequest{})
// 		require.NoError(t, err)
// 		require.NotNil(t, respForce)

// 		waitOperations(ctx, t, wallet, respBefore.GetOperationsCount()+1)
// 	})
// }

// func makeCredit(ctx context.Context, t *testing.T, wallet desc.WalletClient) {
// 	t.Helper()

// 	t.Fatal("тут ваша реализации создания операции debit или credit")
// }

// func polling(ctx context.Context, t *testing.T, wallet desc.WalletClient, greaterOrEqual int64) bool {
// 	t.Helper()

// 	ticker := time.NewTicker(500 * time.Millisecond)
// 	defer ticker.Stop()
// 	ctx, cancel := context.WithTimeout(ctx, 20*time.Second)
// 	defer cancel()

// 	for {
// 		select {
// 		case tick := <-ticker.C:
// 			resp, err := wallet.CollectOperation(ctx, &desc.OperationCountRequest{})
// 			require.NoError(t, err)
// 			require.NotNil(t, resp)
// 			t.Log(tick, "ping")

// 			if resp.GetOperationsCount() >= greaterOrEqual {
// 				return true
// 			}
// 		case <-ctx.Done():
// 			t.Log(ctx.Err())
// 			require.FailNow(t, ctx.Err().Error())
// 		}
// 	}
// }

// func waitOperations(ctx context.Context, t *testing.T, wallet desc.WalletClient, waitOperationCount int64) {
// 	t.Helper()

// 	require.Eventually(t, func() bool {
// 		resp, err := wallet.CollectOperation(ctx, &desc.OperationCountRequest{})
// 		require.NoError(t, err)
// 		require.NotNil(t, resp)

// 		t.Log(time.Now(), "ping")
// 		return resp.GetOperationsCount() >= waitOperationCount
// 	}, 20*time.Second, 500*time.Millisecond)
// }
