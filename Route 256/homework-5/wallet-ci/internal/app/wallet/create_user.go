package wallet

import (
	"context"
	"fmt"
	"regexp"

	"gitlab.ozon.dev/route256/wallet/internal/config"
	"gitlab.ozon.dev/route256/wallet/internal/pkg/domain"
	"gitlab.ozon.dev/route256/wallet/internal/pkg/utils"
	desc "gitlab.ozon.dev/route256/wallet/pkg/api/wallet"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var phonePattern = regexp.MustCompile(`8\d{10}`)

func (i *Implementation) CreateUser(ctx context.Context, req *desc.CreateUserRequest) (*desc.CreateUserResponse, error) {
	ok := phonePattern.MatchString(req.Phone)
	if !ok {
		return nil, status.Error(codes.InvalidArgument, "phone not match pattern 8xxxxxxxxxx")
	}

	user := &domain.User{
		Name:         req.Name,
		Lastname:     req.Lastname,
		Age:          req.Age,
		Phone:        req.Phone,
		PasswordHash: utils.GetPasswordHash(req.Password),
		Level:        config.ANON,
	}

	err := i.store.CreateUser(ctx, user)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("CreateUser err: %v", err))
	}

	return &desc.CreateUserResponse{
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
