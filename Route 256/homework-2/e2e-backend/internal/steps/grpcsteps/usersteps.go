package grpcsteps

import (
	"context"
	"e2e-backend/internal/models"
	"e2e-backend/internal/pb/gitlab.ozon.dev/route256/wallet"
	"e2e-backend/internal/utils/config"
	"e2e-backend/internal/utils/logger"
	"strings"
	"testing"

	"github.com/ddosify/go-faker/faker"
	"github.com/stretchr/testify/require"
)

type UserSteps struct {
	cli  wallet.WalletClient
	fake faker.Faker
	log  *logger.CtxLogger
}

func NewUserSteps(cli wallet.WalletClient) *UserSteps {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	logService, err := logger.NewLoggerService(cfg)
	if err != nil {
		panic(err)
	}

	log := logService.NewPrefix("UserSteps")

	return &UserSteps{
		cli:  cli,
		fake: faker.NewFaker(),
		log:  log,
	}
}

func (s *UserSteps) CreateUser(ctx context.Context, t *testing.T, user *models.User) (*models.User, error) {
	if user == nil {
		user = &models.User{
			Name:     s.fake.RandomPersonFirstName(),
			Lastname: s.fake.RandomPersonLastName(),
			Age:      int32(s.fake.RandomDigitNotNull()),
			Phone:    strings.Replace("8"+s.fake.RandomPhoneNumber(), "-", "", -1),
			Password: s.fake.RandomPassword(),
		}
	}

	req := &wallet.CreateUserRequest{
		Name:     user.Name,
		Lastname: user.Lastname,
		Age:      user.Age,
		Phone:    user.Phone,
		Password: user.Password,
	}

	s.log.InfofJSON("Sending CreateUser request", req)
	resp, err := s.cli.CreateUser(ctx, req)
	if err != nil {
		s.log.Errorf("Failed to create user: %v", err)
		return nil, err
	}

	s.log.InfofJSON("Received CreateUser response", resp)

	user.ID = resp.Info.Id
	user.Level = resp.Info.GetIdentificationLevel().String()

	return user, nil

}

func (s *UserSteps) LogIn(ctx context.Context, t *testing.T, phone string, password string) (*wallet.LogInResponse, error) {
	req := &wallet.LogInRequest{
		Phone:    phone,
		Password: password,
	}

	s.log.InfofJSON("Sending LoginUser request", req)
	resp, err := s.cli.LogIn(ctx, req)
	if err != nil {
		s.log.Errorf("Failed to login user: %v", err)
		return nil, err
	}

	s.log.InfofJSON("Received LoginUser response", resp)

	return resp, nil
}

func (s *UserSteps) GetUser(ctx context.Context, t *testing.T, userID string) (*models.User, error) {
	req := &wallet.GetUserRequest{
		UserId: userID,
	}

	s.log.InfofJSON("Sending GetUser request", req)
	resp, err := s.cli.GetUser(ctx, req)
	if err != nil {
		s.log.Errorf("Failed to get user: %v", err)
		return nil, err
	}

	s.log.InfofJSON("Received GetUser response", resp)

	user := &models.User{
		ID:       resp.Info.Id,
		Name:     resp.Info.Name,
		Lastname: resp.Info.Lastname,
		Age:      resp.Info.Age,
		Phone:    resp.Info.Phone,
		Level:    resp.Info.GetIdentificationLevel().String(),
	}

	return user, nil
}

func (s *UserSteps) CreateAccount(ctx context.Context, t *testing.T, userID string, description string, amount int32) (*models.Account, error) {
	req := &wallet.CreateAccountRequest{
		UserId:      userID,
		Description: description,
		Amount:      amount,
	}

	s.log.InfofJSON("Sending CreateAccount request", req)
	resp, err := s.cli.CreateAccount(ctx, req)
	if err != nil {
		s.log.Errorf("Failed to create account: %v", err)
		return nil, err
	}

	s.log.InfofJSON("Received CreateAccount response", resp)

	account := &models.Account{
		Description: resp.Description,
		AccountId:   resp.AccountId,
		Amount:      resp.Amount,
	}

	return account, nil
}

func (s *UserSteps) GetAccountBalance(ctx context.Context, t *testing.T, userID, accountID string) (*models.Balance, error) {
	req := &wallet.GetAccountBalanceRequest{
		UserId:    userID,
		AccountId: accountID,
	}

	s.log.InfofJSON("Sending GetAccountBalance request", req)
	resp, err := s.cli.GetAccountBalance(ctx, req)
	if err != nil {
		s.log.Errorf("Failed to create account: %v", err)
		return nil, err
	}

	s.log.InfofJSON("Received GetAccountBalance response", resp)

	balance := &models.Balance{
		AccountId: resp.AccountId,
		Amount:    resp.Amount,
	}

	return balance, nil
}

func (s *UserSteps) UpgradeUser(ctx context.Context, t *testing.T, userID string) error {
	req := &wallet.UpgradeUserRequest{
		UserId: userID,
	}

	s.log.InfofJSON("Sending UpgradeUser request", req)
	resp, err := s.cli.UpgradeUser(ctx, req)
	if err != nil {
		s.log.InfofJSON("Failed to upgrade user: %v", err)
		return err
	}

	s.log.InfofJSON("Received UpgradeUser response", resp)
	require.NotNil(t, resp)

	return nil
}

func (s *UserSteps) Credit(ctx context.Context, t *testing.T, userID, accountID string, amount int32) error {
	req := &wallet.CreditRequest{
		UserId:    userID,
		AccountId: accountID,
		Amount:    amount,
	}

	s.log.InfofJSON("Sending Credit request", req)
	resp, err := s.cli.Credit(ctx, req)
	if err != nil {
		s.log.Errorf("Failed to credit account: %v", err)
		return err
	}

	s.log.InfofJSON("Received Credit response", resp)
	require.NotNil(t, resp)

	return nil
}

func (s *UserSteps) Debit(ctx context.Context, t *testing.T, userID, accountID string, amount int32) error {
	req := &wallet.DebitRequest{
		UserId:    userID,
		AccountId: accountID,
		Amount:    amount,
	}

	s.log.InfofJSON("Sending Debit request", req)
	resp, err := s.cli.Debit(ctx, req)
	if err != nil {
		s.log.Errorf("Failed to debit account: %v", err)
		return err
	}

	s.log.InfofJSON("Received Debit response", resp)
	require.NotNil(t, resp)

	return nil
}
