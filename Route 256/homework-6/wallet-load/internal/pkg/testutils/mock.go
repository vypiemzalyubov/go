package testutils

import (
	"context"
	"fmt"

	"gitlab.ozon.dev/route256/wallet/internal/pkg/domain"
	"gitlab.ozon.dev/route256/wallet/internal/pkg/storage"
)

// Т.к мокать по заданию можно только базу, то для других случаев сделал кастомный мок

type FaultyStore struct {
	storage.Storage
}

func (f FaultyStore) LogIn(context.Context, string, string) error {
	return fmt.Errorf("cannot login")
}

func (f FaultyStore) GetUser(context.Context, string) (*domain.UserWithAccounts, error) {
	return nil, fmt.Errorf("user not found")
}
