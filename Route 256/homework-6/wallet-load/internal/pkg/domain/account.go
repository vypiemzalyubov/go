package domain

import (
	"time"

	"github.com/google/uuid"
)

// Account информация о счете
type Account struct {
	ID          uuid.UUID `db:"id"`
	UserID      uuid.UUID `db:"user_id"`
	AccountID   string    `db:"account_id"`
	Amount      int32     `db:"amount"`
	Description string    `db:"description"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
