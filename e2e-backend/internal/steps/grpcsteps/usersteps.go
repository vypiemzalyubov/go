package grpcsteps

import (
	"context"
	"e2e-backend/internal/models"
	"e2e-backend/internal/pb/gitlab.ozon.dev/route256/wallet"
	"strings"
	"testing"

	"github.com/ddosify/go-faker/faker"
	"github.com/stretchr/testify/require"
)

type UserSteps struct {
	cli  wallet.WalletClient
	fake faker.Faker
}

func NewUserSteps(cli wallet.WalletClient) *UserSteps {
	return &UserSteps{
		cli:  cli,
		fake: faker.NewFaker(),
	}
}

func (s *UserSteps) CreateUser(ctx context.Context, t *testing.T) *models.User {
	user := &models.User{
		Name:     s.fake.RandomPersonFirstName(),
		Lastname: s.fake.RandomPersonLastName(),
		Age:      int32(s.fake.RandomDigitNotNull()),
		Phone:    strings.Replace("8"+s.fake.RandomPhoneNumber(), "-", "", -1),
		Password: s.fake.RandomPassword(),
	}

	resp, err := s.cli.CreateUser(ctx, &wallet.CreateUserRequest{
		Name:     user.Name,
		Lastname: user.Lastname,
		Age:      user.Age,
		Phone:    user.Phone,
		Password: user.Password,
	})
	require.NoError(t, err)
	require.NotNil(t, resp)
	user.ID = resp.Info.Id
	user.Level = resp.Info.GetIdentificationLevel().String()

	return user
}
