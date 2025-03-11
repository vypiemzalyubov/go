//go:build httptest
// +build httptest

package http

import (
	"context"
	"flag"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"gitlab.ozon.dev/route256/wallet/internal/config"
	"gitlab.ozon.dev/route256/wallet/internal/pkg/domain"
	"gitlab.ozon.dev/route256/wallet/internal/pkg/storage"
	"gitlab.ozon.dev/route256/wallet/test/utils/clients/http"
)

var (
	store    storage.Storage
	client   http.WalletClient
	basePaht = "http://localhost:8001"
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
	})
	if err != nil {
		log.Fatal().Err(err).Msg("Can't init connect to DB")
	}
	defer db.Close()

	store = storage.New(db)
	client = http.NewWalletClient(basePaht, time.Second)

	code := m.Run()
	os.Exit(code)
}

// nolint:unused
func mustCreateAccount(ctx context.Context, amount int32) (*domain.Account, error) {
	account := &domain.Account{
		AccountID:   uuid.NewString(),
		Amount:      amount,
		Description: "test account",
	}
	err := store.AddAccount(ctx, account)
	if err != nil {
		return nil, err
	}

	return account, nil
}
