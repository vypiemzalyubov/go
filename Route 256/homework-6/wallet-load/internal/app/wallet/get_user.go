package wallet

import (
	"context"
	"errors"
	"fmt"

	desc "gitlab.ozon.dev/route256/wallet/pkg/api/wallet"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) GetUser(ctx context.Context, req *desc.GetUserRequest) (*desc.GetUserResponse, error) {
	if err := i.CheckToken(ctx, req.UserId); err != nil {
		return nil, err
	}

	user, err := i.store.GetUser(ctx, req.UserId)
	if errors.Is(err, fmt.Errorf("account not found")) {
		return nil, status.Error(codes.NotFound, "user not found")
	}

	return &desc.GetUserResponse{
		Info: &desc.User{
			Id:                  user.ID.String(),
			Name:                user.Name,
			Lastname:            user.Lastname,
			Age:                 user.Age,
			Phone:               user.Phone,
			IdentificationLevel: desc.IdentificationLevel(desc.IdentificationLevel_value[user.Level]),
		},
	}, nil
}
