package wallet

import (
	context "context"
	"fmt"

	"github.com/google/uuid"
	"gitlab.ozon.dev/route256/wallet/internal/config"
	desc "gitlab.ozon.dev/route256/wallet/pkg/api/wallet"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Debit поступление денег
func (i *Implementation) Debit(ctx context.Context, req *desc.DebitRequest) (*desc.DebitResponse, error) {
	if err := i.CheckToken(ctx, req.UserId); err != nil {
		return nil, err
	}

	user, err := i.User(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	account, ok := user.GetAccount(req.GetAccountId())
	if !ok {
		return nil, status.Error(codes.NotFound, "account not found")
	}

	if user.Level == config.ANON && account.Amount+req.GetAmount() > config.MaxAnonAmount {
		return &desc.DebitResponse{
			Status: desc.OperationStatus_STATUS_FAIL,
		}, status.Error(codes.InvalidArgument, "unavailable operation")
	}

	if user.Level == config.FULL && account.Amount+req.GetAmount() > config.MaxFullAmount {
		return &desc.DebitResponse{
			Status: desc.OperationStatus_STATUS_FAIL,
		}, status.Error(codes.InvalidArgument, "unavailable operation")
	}

	operationID := req.GetOperationId()
	if operationID == "" {
		operationID = uuid.NewString()
	}
	err = i.store.Debit(ctx, req.GetAccountId(), req.GetAmount(), operationID)
	if err != nil {
		return &desc.DebitResponse{
			Status: desc.OperationStatus_STATUS_FAIL,
		}, status.Error(codes.Internal, fmt.Sprintf("Debit process err: %v", err))
	}

	return &desc.DebitResponse{
		Status: desc.OperationStatus_STATUS_OK,
	}, nil
}
