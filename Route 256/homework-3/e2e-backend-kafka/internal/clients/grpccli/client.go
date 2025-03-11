package grpccli

import (
	"e2e-backend/internal/pb/gitlab.ozon.dev/route256/wallet"
	"e2e-backend/internal/utils/config"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewWalletClient() (wallet.WalletClient, *grpc.ClientConn, error) {
	cfg, err := config.NewConfig()
	if err != nil {
		return nil, nil, err
	}

	address := fmt.Sprintf("%s:%d", cfg.Wallet.Host, cfg.Wallet.Port)

	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, err
	}

	client := wallet.NewWalletClient(conn)
	return client, conn, nil
}
