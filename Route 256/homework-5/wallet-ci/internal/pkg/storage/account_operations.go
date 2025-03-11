package storage

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"gitlab.ozon.dev/route256/wallet/internal/pkg/domain"
)

func (s *storage) Credit(ctx context.Context, accountID string, amount int32, operationID string) error {
	err := s.AddAccountOperation(ctx, &domain.AccountOperation{
		AccountID:     accountID,
		Amount:        amount,
		OperationID:   operationID,
		OperationType: domain.OperationTypeCredit,
	})
	return err
}

func (s *storage) Debit(ctx context.Context, accountID string, amount int32, operationID string) error {
	err := s.AddAccountOperation(ctx, &domain.AccountOperation{
		AccountID:     accountID,
		Amount:        amount,
		OperationID:   operationID,
		OperationType: domain.OperationTypeDebit,
	})
	return err
}

func (s *storage) AddAccountOperation(ctx context.Context, operation *domain.AccountOperation) error {
	balance, err := s.GetAccount(ctx, operation.AccountID)
	if err != nil {
		return fmt.Errorf("can't get account balance for change amount: %w", err)
	}

	err = s.DoInTransaction(ctx, s.db, func(ctx context.Context, tx *sqlx.Tx) error {
		err = s.addAccountOperation(ctx, tx, operation)
		if err != nil {
			return err
		}

		newAccountAmount := balance.Amount + (int32(operation.OperationType) * operation.Amount)
		err = s.changeAccountAmount(ctx, tx, operation.AccountID, newAccountAmount)
		if err != nil {
			return err
		}

		return nil
	})

	return err
}

func (s *storage) addAccountOperation(ctx context.Context, db *sqlx.Tx, operation *domain.AccountOperation) error {
	stmt := s.Builder().Insert("account_operations").
		Columns("account_id, amount, operation_id", "operation_type").
		Values(operation.AccountID, operation.Amount, operation.OperationID, operation.OperationType).
		Suffix("RETURNING id")

	req, args, err := stmt.ToSql()
	if err != nil {
		return err
	}

	err = db.GetContext(ctx, &operation.ID, req, args...)

	return err
}

func (s *storage) GetAccountOperations(ctx context.Context, accountID string, limit uint64) ([]domain.AccountOperation, error) {
	var operations []domain.AccountOperation

	stmt := s.Builder().Select("*").
		From("account_operations").
		Where("account_id = ?", accountID).
		OrderBy("created_at").
		Limit(limit)

	req, args, err := stmt.ToSql()
	if err != nil {
		return nil, err
	}

	err = s.db.SelectContext(ctx, &operations, req, args...)
	if err != nil {
		return nil, err
	}

	return operations, nil
}

func (s *storage) CollectOperation(ctx context.Context) (int64, error) {
	var count int64

	err := s.db.GetContext(ctx, &count, "SELECT COUNT(*)  FROM account_operations")

	if err != nil {
		log.Err(err).Msgf("error scanning account_operations table")
		return 0, err
	}

	return count, nil
}
