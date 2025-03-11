package wallet

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"gitlab.ozon.dev/route256/wallet/internal/pkg/domain"
	desc "gitlab.ozon.dev/route256/wallet/pkg/api/wallet"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateAccount Создать счет
func (i *Implementation) CreateAccount(ctx context.Context, req *desc.CreateAccountRequest) (*desc.CreateAccountResponse, error) {
	if err := i.CheckToken(ctx, req.UserId); err != nil {
		return nil, err
	}

	user, err := i.User(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	if len(user.Accounts) == 3 {
		return nil, status.Error(codes.AlreadyExists, "User allready has 3 accounts")
	}

	acc := &domain.Account{
		UserID:      user.ID,
		AccountID:   uuid.NewString(),
		Amount:      req.GetAmount(),
		Description: req.GetDescription(),
	}
	err = i.store.AddAccount(ctx, acc)
	if err != nil {
		log.Error()
		return nil, status.Error(codes.Internal, fmt.Sprintf("Create account err: %v", err))
	}

	return &desc.CreateAccountResponse{
		AccountId:   acc.AccountID,
		Amount:      acc.Amount,
		Description: acc.Description,
	}, nil
}
