package domain

import (
	"time"

	"github.com/google/uuid"
)

// User информация о счете
type User struct {
	ID           uuid.UUID `db:"id"`
	Name         string    `db:"name"`
	Lastname     string    `db:"lastname"`
	Age          int32     `db:"age"`
	Phone        string    `db:"phone"`
	PasswordHash string    `db:"password_hash"`
	Level        string    `db:"level"`
	CreatedAt    time.Time `db:"created_at"`
}

type UserWithAccounts struct {
	User
	RawPassword string
	Accounts    []*Account
}

func (u *UserWithAccounts) GetAccount(accountID string) (*Account, bool) {
	for _, account := range u.Accounts {
		if account.AccountID == accountID {
			return account, true
		}
	}
	return nil, false
}
