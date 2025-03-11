package wallet

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/route256/wallet/internal/pkg/worker"
	desc "gitlab.ozon.dev/route256/wallet/pkg/api/wallet"
)

type MockWorkerForceCollectOps struct {
	ForceSignal chan struct{}
}

func (m *MockWorkerForceCollectOps) ForceCollect() {
	m.ForceSignal <- struct{}{}
}

func TestForceCollectOperation(t *testing.T) {
	t.Run("Simple force collect operation", func(t *testing.T) {
		forceSignal := make(chan struct{})

		mockWorker := &MockWorkerForceCollectOps{ForceSignal: forceSignal}

		originalForceSignal := worker.ForceSignal
		worker.ForceSignal = mockWorker.ForceSignal
		defer func() { worker.ForceSignal = originalForceSignal }()

		impl := &Implementation{}

		ctx := context.Background()
		req := &desc.OperationCountRequest{}

		signalReceived := make(chan struct{})
		go func() {
			<-forceSignal
			close(signalReceived)
		}()

		_, err := impl.ForceCollectOperation(ctx, req)
		require.NoError(t, err)

		select {
		case <-signalReceived:
		default:
			t.Fatal("Expected force collect signal not received")
		}
	})
}
