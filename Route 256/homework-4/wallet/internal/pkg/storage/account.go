package storage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"gitlab.ozon.dev/route256/wallet/internal/pkg/domain"
)

const AccountNotFound = "account not found"

func (s *storage) AddAccount(ctx context.Context, account *domain.Account) error {
	stmt := s.Builder().Insert("accounts").
		Columns("user_id", "account_id", "amount", "description").
		Values(account.UserID, account.AccountID, account.Amount, account.Description).
		Suffix("RETURNING id")

	req, args, err := stmt.ToSql()
	if err != nil {
		return err
	}

	err = s.db.GetContext(ctx, &account.ID, req, args...)

	return err
}

func (s *storage) GetAccount(ctx context.Context, accountID string) (*domain.Account, error) {
	var balance domain.Account

	stmt := s.Builder().Select("*").
		From("accounts").
		Where("account_id = ?", accountID)

	req, args, err := stmt.ToSql()
	if err != nil {
		return nil, err
	}

	err = s.db.GetContext(ctx, &balance, req, args...)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf(AccountNotFound)
	}
	if err != nil {
		return nil, err
	}

	return &balance, nil
}

func (s *storage) GetUserAccounts(ctx context.Context, userID string) ([]*domain.Account, error) {
	var accounts []*domain.Account

	stmt := s.Builder().Select("*").
		From("accounts").
		Where("user_id = ?", userID)

	req, args, err := stmt.ToSql()
	if err != nil {
		return nil, err
	}

	err = s.db.SelectContext(ctx, &accounts, req, args...)

	if err == sql.ErrNoRows {
		return []*domain.Account{}, nil
	}
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

func (s *storage) changeAccountAmount(ctx context.Context, db *sqlx.Tx, accountID string, newAmount int32) error {
	stmt := s.Builder().
		Update("accounts").
		Set("amount", newAmount).
		Where("account_id = ?", accountID)

	req, args, err := stmt.ToSql()
	if err != nil {
		return err
	}
	_, err = db.ExecContext(ctx, req, args...)

	return err
}
