package wallet

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"gitlab.ozon.dev/route256/wallet/internal/pkg/utils"
	desc "gitlab.ozon.dev/route256/wallet/pkg/api/wallet"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) LogIn(ctx context.Context, req *desc.LogInRequest) (*desc.LogInResponse, error) {
	ok := phonePattern.MatchString(req.Phone)
	if !ok {
		return nil, status.Error(codes.InvalidArgument, "phone not match pattern 8xxxxxxxxxx")
	}

	user, err := i.store.GetUserByPhone(ctx, req.Phone)
	if errors.Is(err, fmt.Errorf("account not found")) {
		return nil, status.Error(codes.NotFound, "user not found")
	}

	if utils.GetPasswordHash(req.Password) != user.PasswordHash {
		return nil, status.Error(codes.InvalidArgument, "ivalid password") // Опечатка в invalid
	}

	token := uuid.NewString()
	err = i.store.LogIn(ctx, user.ID.String(), token)
	if err != nil {
		return nil, status.Error(codes.Internal, "cannot login")
	}

	return &desc.LogInResponse{
		Token: token,
	}, nil
}
