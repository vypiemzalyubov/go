package wallet

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/rs/zerolog/log"
	"gitlab.ozon.dev/route256/wallet/internal/config"
	"gitlab.ozon.dev/route256/wallet/internal/pkg/storage"
)

var store storage.Storage

func TestMain(m *testing.M) {
	db, err := config.ConnectToPostgres(context.Background(), config.Database{
		Dsn:               os.Getenv(config.DSNKey),
		ConnTimeout:       time.Duration(10),
		MaxOpenConn:       50,
		MasterMaxIdleConn: 50,
		Driver:            "pgx",
		Migrations:        "migrations",
	})
	defer func() {
		if err = db.Close(); err != nil {
			panic(err)
		}
	}()

	store = storage.New(db)

	if err != nil {
		log.Fatal().Err(err).Msg("Can't init connect to DB")
	}

	code := m.Run()
	os.Exit(code)
}
