package config

import (
	"context"

	_ "github.com/jackc/pgx/v4"        // nolint:blank-imports
	_ "github.com/jackc/pgx/v4/stdlib" // nolint:blank-imports
	"github.com/jmoiron/sqlx"
)

// ConnectToPostgres init connect to postgres
func ConnectToPostgres(ctx context.Context, cfg Database) (*sqlx.DB, error) {
	conn, err := sqlx.Open("pgx", cfg.Dsn)
	if err != nil {
		return nil, err
	}
	conn.SetConnMaxLifetime(cfg.ConnTimeout)
	conn.SetMaxOpenConns(cfg.MaxOpenConn)
	conn.SetMaxIdleConns(cfg.MasterMaxIdleConn)

	if err := conn.PingContext(ctx); err != nil {
		return nil, err
	}

	return conn, nil
}
