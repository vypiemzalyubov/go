package storage

import (
	"context"
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"gitlab.ozon.dev/route256/wallet/internal/pkg/domain"
)

// Storage - интерфейс для работы с бд
type Storage interface {
	CreateUser(ctx context.Context, user *domain.User) error
	UpgradeUser(ctx context.Context, userID string, level string) error
	GetUser(ctx context.Context, userID string) (*domain.UserWithAccounts, error)
	GetUserSessions(ctx context.Context, userID string) ([]string, error)
	GetUserByPhone(ctx context.Context, phone string) (*domain.User, error)
	LogIn(ctx context.Context, userID string, token string) error

	CollectOperation(ctx context.Context) (int64, error)
	AddAccount(ctx context.Context, account *domain.Account) error
	GetAccount(ctx context.Context, accountID string) (*domain.Account, error)
	Debit(ctx context.Context, accountID string, amount int32, operationID string) error
	Credit(ctx context.Context, accountID string, amount int32, operationID string) error
	GetAccountOperations(ctx context.Context, accountID string, limit uint64) ([]domain.AccountOperation, error)

	Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
}

type storage struct {
	db *sqlx.DB
}

// New .
func New(db *sqlx.DB) Storage {
	return &storage{
		db: db,
	}
}

// Builder вернет squirrel SQL Builder объект
func (s *storage) Builder() sq.StatementBuilderType {
	return sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
}

func (s *storage) DoInTransaction(ctx context.Context, db *sqlx.DB, fn func(ctx context.Context, tx *sqlx.Tx) error) (err error) {

	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("can't begin transaction: %w", err)
	}

	defer func() {
		if r := recover(); r != nil {
			rollbackTx(ctx, tx)
		} else if err != nil {
			rollbackTx(ctx, tx)
		} else {
			if err = tx.Commit(); err != nil {
				err = fmt.Errorf("cannot commit transaction: %w", err)
			}
		}
	}()

	err = fn(ctx, tx)

	return err
}

func rollbackTx(_ context.Context, tx *sqlx.Tx) {
	if err := tx.Rollback(); err != nil {
		log.Error().Err(err).Msg("sqltx/rollbackTx: cannot rollback transaction")
	}
}

func (s *storage) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return s.db.ExecContext(ctx, query, args...)
}
