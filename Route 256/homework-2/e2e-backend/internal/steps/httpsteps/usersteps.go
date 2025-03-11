package httpsteps

import (
	"e2e-backend/internal/clients/httpcli"

	"github.com/ddosify/go-faker/faker"
)

type UserSteps struct {
	cli  httpcli.WalletClient
	fake faker.Faker
}

func NewUserSteps(cli httpcli.WalletClient) *UserSteps {
	return &UserSteps{
		cli:  cli,
		fake: faker.NewFaker(),
	}
}
