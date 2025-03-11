package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"gitlab.ozon.dev/route256/wallet/internal/config"
	desc "gitlab.ozon.dev/route256/wallet/pkg/api/wallet"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func createGatewayServer(ctx context.Context, cfg *config.AppConfig) (*http.Server, error) {
	conn, err := grpc.NewClient(
		cfg.Grpc.Addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to dial server: %w", err)
	}

	mux := runtime.NewServeMux()

	if err := desc.RegisterWalletHandler(ctx, mux, conn); err != nil {
		return nil, fmt.Errorf("failed registration handler: %w", err)
	}

	//nolint:gosec
	gatewayServer := &http.Server{
		Addr:    cfg.Rest.Addr,
		Handler: mux,
	}

	return gatewayServer, nil
}
