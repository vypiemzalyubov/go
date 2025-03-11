package wallet

import (
	"context"

	"gitlab.ozon.dev/route256/wallet/internal/pkg/worker"
	desc "gitlab.ozon.dev/route256/wallet/pkg/api/wallet"
)

// ForceCollectOperation Подсчет колличества операций
func (i *Implementation) ForceCollectOperation(_ context.Context, _ *desc.OperationCountRequest) (*desc.ForceOperationCountResponse, error) {
	worker.ForceSignal <- struct{}{}
	return &desc.ForceOperationCountResponse{}, nil
}
