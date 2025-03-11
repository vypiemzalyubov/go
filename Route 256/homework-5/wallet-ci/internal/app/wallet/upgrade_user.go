package wallet

import (
	"context"
	"errors"
	"fmt"

	"gitlab.ozon.dev/route256/wallet/internal/config"
	desc "gitlab.ozon.dev/route256/wallet/pkg/api/wallet"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) UpgradeUser(ctx context.Context, req *desc.UpgradeUserRequest) (*desc.UpgradeUserResponse, error) {
	if err := i.CheckToken(ctx, req.UserId); err != nil {
		return nil, err
	}

	user, err := i.store.GetUser(ctx, req.UserId)
	if errors.Is(err, fmt.Errorf("account not found")) {
		return nil, status.Error(codes.NotFound, "user not found")
	}

	if user.Level == config.FULL {
		return nil, status.Error(codes.AlreadyExists, "user allready has FULL level")
	}
	err = i.store.UpgradeUser(ctx, req.UserId, "FULL")
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("UpgradeUser err: %v", err))
	}

	return &desc.UpgradeUserResponse{}, nil
}
