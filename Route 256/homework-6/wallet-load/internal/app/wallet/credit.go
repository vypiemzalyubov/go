package wallet

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	desc "gitlab.ozon.dev/route256/wallet/pkg/api/wallet"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Credit списание денег
func (i *Implementation) Credit(ctx context.Context, req *desc.CreditRequest) (*desc.CreditResponse, error) {
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

	if account.Amount-req.GetAmount() < 0 {
		return &desc.CreditResponse{
			Status: desc.OperationStatus_STATUS_FAIL,
		}, status.Error(codes.InvalidArgument, "unavailable operation")
	}

	operationID := req.GetOperationId()
	if operationID == "" {
		operationID = uuid.NewString()
	}

	err = i.store.Credit(ctx, account.AccountID, req.GetAmount(), operationID)
	if err != nil {
		return &desc.CreditResponse{
			Status: desc.OperationStatus_STATUS_FAIL,
		}, status.Error(codes.Internal, fmt.Sprintf("Credit process err: %v", err))
	}

	return &desc.CreditResponse{
		Status: desc.OperationStatus_STATUS_OK,
	}, nil
}
