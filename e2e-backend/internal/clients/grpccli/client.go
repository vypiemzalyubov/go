package grpccli

import (
	"e2e-backend/internal/pb/gitlab.ozon.dev/route256/wallet"

	"google.golang.org/grpc"
)

func NewWalletClient() (wallet.WalletClient, *grpc.ClientConn) {
	// TODO: необходимо реализовать коннект к приложению по grpc
	return nil, nil
}

