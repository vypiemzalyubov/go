//go:build grpctest
// +build grpctest

package grpc

import (
	"context"
	"flag"
	"os"
	"testing"
	"time"

	"github.com/rs/zerolog/log"
	"gitlab.ozon.dev/route256/wallet/internal/config"
	"gitlab.ozon.dev/route256/wallet/internal/pkg/storage"
	desc "gitlab.ozon.dev/route256/wallet/pkg/api/wallet"
	"gitlab.ozon.dev/route256/wallet/test/utils/clients/grpc"
)

var (
	store        storage.Storage
	walletClient desc.WalletClient
)

func TestMain(m *testing.M) {
	flag.Parse()

	ctx := context.Background()

	db, err := config.ConnectToPostgres(ctx, config.Database{
		Dsn:               os.Getenv(config.DSNKey),
		ConnTimeout:       time.Duration(10),
		MaxOpenConn:       50,
		MasterMaxIdleConn: 50,
		Driver:            "pgx",
		Migrations:        "migrations",
	})
	if err != nil {
		log.Fatal().Err(err).Msg("Can't init connect to DB")
	}
	defer db.Close()

	store = storage.New(db)

	client, clientConn := grpc.NewServiceClient()
	defer clientConn.Close()
	walletClient = client

	code := m.Run()
	os.Exit(code)
}
