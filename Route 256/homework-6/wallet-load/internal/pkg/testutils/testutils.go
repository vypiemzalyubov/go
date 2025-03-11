package testutils

import (
	"context"
	"strings"

	"github.com/ddosify/go-faker/faker"
	"github.com/google/uuid"
	"gitlab.ozon.dev/route256/wallet/internal/pkg/domain"
	"gitlab.ozon.dev/route256/wallet/internal/pkg/storage"
	"gitlab.ozon.dev/route256/wallet/internal/pkg/token"
	"gitlab.ozon.dev/route256/wallet/internal/pkg/utils"
)

func CreateUserAndAuth(ctx context.Context, store storage.Storage, accounts ...*domain.Account) (context.Context, *domain.UserWithAccounts, error) {
	fake := faker.NewFaker()

	randomPhone := strings.Replace("8"+fake.RandomPhoneNumber(), "-", "", -1)

	password := fake.RandomPassword()
	user := &domain.User{
		Name:         fake.RandomPersonFirstName(),
		Lastname:     fake.RandomPersonLastName(),
		Age:          int32(fake.RandomDigitNotNull()),
		Phone:        randomPhone,
		PasswordHash: utils.GetPasswordHash(password),
	}

	err := store.CreateUser(ctx, user)
	if err != nil {
		return ctx, nil, err
	}

	session := uuid.NewString()
	err = store.LogIn(ctx, user.ID.String(), session)
	if err != nil {
		return ctx, nil, err
	}

	for _, acc := range accounts {
		err = store.AddAccount(ctx, &domain.Account{
			UserID:      user.ID,
			AccountID:   acc.AccountID,
			Amount:      acc.Amount,
			Description: acc.Description,
		})
		if err != nil {
			return ctx, nil, err
		}
	}

	ctxWithToken := token.ToCtx(ctx, session)
	userDb, err := store.GetUser(ctx, user.ID.String())
	if err != nil {
		return ctx, nil, err
	}
	userDb.RawPassword = password

	return ctxWithToken, userDb, nil
}
