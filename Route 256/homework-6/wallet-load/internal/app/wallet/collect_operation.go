package wallet

import (
	"context"

	"gitlab.ozon.dev/route256/wallet/internal/pkg/worker"
	desc "gitlab.ozon.dev/route256/wallet/pkg/api/wallet"
)

// CollectOperation Подсчет колличества операций
func (i *Implementation) CollectOperation(_ context.Context, _ *desc.OperationCountRequest) (*desc.OperationCountResponse, error) {
	return &desc.OperationCountResponse{OperationsCount: worker.Count.Load()}, nil
}
