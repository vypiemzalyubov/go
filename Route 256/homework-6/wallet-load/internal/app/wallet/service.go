package wallet

import (
	"context"
	"errors"
	"fmt"
	"os"

	"gitlab.ozon.dev/route256/wallet/internal/config"
	"gitlab.ozon.dev/route256/wallet/internal/pkg/cbr"
	"gitlab.ozon.dev/route256/wallet/internal/pkg/domain"
	"gitlab.ozon.dev/route256/wallet/internal/pkg/kafka"
	"gitlab.ozon.dev/route256/wallet/internal/pkg/storage"
	"gitlab.ozon.dev/route256/wallet/internal/pkg/token"
	desc "gitlab.ozon.dev/route256/wallet/pkg/api/wallet"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Implementation .
type Implementation struct {
	desc.UnimplementedWalletServer

	store     storage.Storage
	kafka     kafka.ClientInterface
	cbrClient *cbr.Client
}

// NewWallet return new instance of Implementation.
func NewWallet(store storage.Storage, kafka kafka.ClientInterface) *Implementation {
	return &Implementation{
		store:     store,
		kafka:     kafka,
		cbrClient: cbr.NewClient(os.Getenv(config.CbrURL)),
	}
}

func (i *Implementation) CheckToken(ctx context.Context, userID string) error {
	sessions, err := i.store.GetUserSessions(ctx, userID)
	if err != nil {
		return status.Error(codes.Internal, fmt.Sprintf("GetUserSession err: %v", err))
	}

	ok := token.Check(ctx, sessions)
	if !ok {
		return status.Error(codes.Unauthenticated, "not authorized")
	}
	return nil
}

func (i *Implementation) User(ctx context.Context, userID string) (*domain.UserWithAccounts, error) {
	user, err := i.store.GetUser(ctx, userID)
	if errors.Is(err, fmt.Errorf("user not found")) {
		return nil, status.Error(codes.NotFound, "user not found")
	}
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return user, nil
}
