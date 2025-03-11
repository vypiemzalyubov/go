package main

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose/v3"
	"github.com/rs/zerolog/log"
	"gitlab.ozon.dev/route256/wallet/internal/config"
	"gitlab.ozon.dev/route256/wallet/internal/server"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := config.InitAppConfig()

	var db *sqlx.DB
	err := config.RunWithRetry("ConnectToPostgres", func() error {
		locDB, err := config.ConnectToPostgres(ctx, cfg.Database)
		db = locDB
		return err
	})
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create DB connection")
	}

	defer db.Close()

	if cfg.Database.Migrations != "" {
		if err := goose.Up(db.DB, cfg.Database.Migrations); err != nil {
			log.Error().Err(err).Msg("migration failed")
		}
	}

	if err := server.NewServer(db).Start(cfg); err != nil {
		log.Error().Err(err).Msg("Failed creating gRPC server")
	}
}
