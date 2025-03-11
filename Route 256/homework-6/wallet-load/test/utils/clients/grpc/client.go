package grpc

import (
	"github.com/rs/zerolog/log"
	desc "gitlab.ozon.dev/route256/wallet/pkg/api/wallet"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// NewServiceClient - create client for server with signInterceptor
func NewServiceClient() (desc.WalletClient, *grpc.ClientConn) {
	dialOpt := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.NewClient(":8002", dialOpt...)

	if err != nil {
		log.Fatal().Err(err).Msg("Failed to dial bufnet")
	}

	client := desc.NewWalletClient(conn)

	return client, conn
}
