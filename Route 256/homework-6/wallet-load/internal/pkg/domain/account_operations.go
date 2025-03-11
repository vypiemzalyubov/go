package domain

import (
	"time"

	"github.com/google/uuid"
)

// OperationType типы операций
type OperationType int

const (
	// OperationTypeDebit поступление денег
	OperationTypeDebit OperationType = 1
	// OperationTypeCredit списание денег
	OperationTypeCredit OperationType = -1
)

// AccountOperation операции по счету
type AccountOperation struct {
	ID            uuid.UUID     `db:"id"`
	AccountID     string        `db:"account_id"`
	Amount        int32         `db:"amount"`
	OperationID   string        `db:"operation_id"`
	OperationType OperationType `db:"operation_type"`
	CreatedAt     time.Time     `db:"created_at"`
	UpdatedAt     time.Time     `db:"updated_at"`
}
