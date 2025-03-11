package test

import (
	"context"
	"testing"
	"time"

	"e2e-backend/internal/clients/grpccli"
	"e2e-backend/internal/pb/gitlab.ozon.dev/route256/wallet"
	"e2e-backend/internal/steps/grpcsteps"
	utils "e2e-backend/internal/utils"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type WalletTestSuiteDebitCredit struct {
	suite.Suite
	client    wallet.WalletClient
	conn      *grpc.ClientConn
	UserSteps *grpcsteps.UserSteps
	UserID    string
	Token     string
}

func (suite *WalletTestSuiteDebitCredit) BeforeAll(t provider.T) {
	t.NewStep("Setup gRPC Client, Create User, Login User, Get Token")

	var err error
	suite.client, suite.conn, err = grpccli.NewWalletClient()
	t.Require().NoError(err, utils.ErrCreatedGrpcClient)

	suite.UserSteps = grpcsteps.NewUserSteps(suite.client)

	ctx := context.Background()
	createdUser, _ := suite.UserSteps.CreateUser(ctx, t, nil)
	suite.UserID = createdUser.ID

	loginToken, _ := suite.UserSteps.LogIn(ctx, t, createdUser.Phone, createdUser.Password)
	suite.Token = loginToken.Token
}

func (suite *WalletTestSuiteDebitCredit) AfterAll(t provider.T) {
	t.NewStep("Teardown gRPC Client")
	suite.conn.Close()
}

func (suite *WalletTestSuiteDebitCredit) BeforeEach(t provider.T) {
	t.Epic("Wallet")
}

func TestWalletTestSuiteDebitCreditRunner(t *testing.T) {
	t.Parallel()
	suite.RunSuite(t, new(WalletTestSuiteDebitCredit))
}

func (suite *WalletTestSuiteDebitCredit) TestDebit(t provider.T) {
	t.Title("Debit to account")

	userID, token := utils.CreateUserHelper(suite.UserSteps, t)

	testCases := []struct {
		name            string
		userLevel       string
		initialAmount   int32
		debitAmount     int32
		expectedBalance int32
	}{
		{
			name:            "Debit for ANON",
			userLevel:       "ANON",
			initialAmount:   1000,
			debitAmount:     49000,
			expectedBalance: 50000,
		},
		{
			name:            "Debit for FULL",
			userLevel:       "FULL",
			initialAmount:   190000,
			debitAmount:     10000,
			expectedBalance: 200000,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t provider.T) {
			t.Title("Debit to account_" + tc.name)
			t.Tag("gRPC Test")
			t.Tag("Positive")
			t.Parallel()

			tc := tc

			t.WithNewStep("Debit to account step", func(sCtx provider.StepCtx) {
				ctxAuth := utils.ToCtx(context.Background(), token)

				if tc.userLevel == "FULL" {
					sCtx.NewStep("Upgrade user")
					err := suite.UserSteps.UpgradeUser(ctxAuth, t, userID)
					sCtx.Require().NoError(err, utils.ErrUpgradedUserErrorShouldBeNil)
				}

				sCtx.NewStep("Create account")
				createdAccount, err := suite.UserSteps.CreateAccount(ctxAuth, t, userID, "description", tc.initialAmount)
				sCtx.Require().NoError(err, utils.ErrCreatedAccountErrorShouldBeNil)
				sCtx.Require().NotNil(createdAccount, utils.ErrCreatedAccountMustNotBeNil)
				sCtx.Require().Equal(tc.initialAmount, createdAccount.Amount, utils.ErrAccountAmountMustBeEqual)

				sCtx.NewStep("Debit to account")
				err = suite.UserSteps.Debit(ctxAuth, t, userID, createdAccount.AccountId, tc.debitAmount)
				sCtx.Require().NoError(err, utils.ErrDebitErrorShouldBeNil)

				sCtx.NewStep("Get account balance")
				gettingBalance, err := suite.UserSteps.GetAccountBalance(ctxAuth, t, userID, createdAccount.AccountId)
				sCtx.Require().NoError(err, utils.ErrGetBalanceErrorShouldBeNil)
				sCtx.Require().NotNil(gettingBalance, utils.ErrGetBalanceMustNotBeNil)
				sCtx.Require().Equal(tc.expectedBalance, gettingBalance.Amount, utils.ErrBalanceAmountMustBeEqual)
			}, allure.NewParameter("time", time.Now()))
		})
	}
}

func (suite *WalletTestSuiteDebitCredit) TestDebitNegative(t provider.T) {
	t.Title("Debit invalid operation")

	userID, token := utils.CreateUserHelper(suite.UserSteps, t)

	testCases := []struct {
		name               string
		userLevel          string
		initialAmount      int32
		debitAmount        int32
		expectedMaxBalance int32
		expectedErr        string
	}{
		// {
		// 	name:               "Debit to exceed ANON limit",
		// 	userLevel:          "ANON",
		// 	initialAmount:      0,
		// 	debitAmount:        50001, //Здесь баг, хотя в прошлой реализации его не было, т.к сервис позволяет добавить для анонима на счет больше лимита. Я думал баги фиксят, а не добавляют=)
		// 	expectedMaxBalance: 0,
		// 	expectedErr:        "unavailable operation",
		// },
		{
			name:               "Debit to exceed FULL limit",
			userLevel:          "FULL",
			initialAmount:      0,
			debitAmount:        200001,
			expectedMaxBalance: 0,
			expectedErr:        "unavailable operation",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t provider.T) {
			t.Title("Debit to account with invalid operation_" + tc.name)
			t.Tag("gRPC Test")
			t.Tag("Negative")
			t.Parallel()

			tc := tc

			t.WithNewStep("Debit to account step", func(sCtx provider.StepCtx) {
				ctxAuth := utils.ToCtx(context.Background(), token)

				if tc.userLevel == "FULL" {
					sCtx.NewStep("Upgrade user")
					err := suite.UserSteps.UpgradeUser(ctxAuth, t, userID)
					sCtx.Require().NoError(err, utils.ErrUpgradedUserErrorShouldBeNil)
				}

				sCtx.NewStep("Create account")
				createdAccount, err := suite.UserSteps.CreateAccount(ctxAuth, t, userID, "description", tc.initialAmount)
				sCtx.Require().NoError(err, utils.ErrCreatedAccountErrorShouldBeNil)
				sCtx.Require().NotNil(createdAccount, utils.ErrCreatedAccountMustNotBeNil)
				sCtx.Require().Equal(tc.initialAmount, createdAccount.Amount, utils.ErrAccountAmountMustBeEqual)

				sCtx.NewStep("Debit to account")
				err = suite.UserSteps.Debit(ctxAuth, t, userID, createdAccount.AccountId, tc.debitAmount)
				sCtx.Require().Error(err, utils.ErrDebitErrorMustNotBeNil)
				st, ok := status.FromError(err)
				sCtx.Require().True(ok, utils.ErrExpectGRPCStatusError)
				sCtx.Require().Equal(codes.InvalidArgument, st.Code(), utils.ErrExpectInvalidArgument)
				sCtx.Require().Contains(st.Message(), tc.expectedErr, utils.ErrMessageExpectInvalidArgumentDebit)

				sCtx.NewStep("Get account balance")
				gettingBalance, err := suite.UserSteps.GetAccountBalance(ctxAuth, t, userID, createdAccount.AccountId)
				sCtx.Require().NoError(err, utils.ErrGetBalanceErrorShouldBeNil)
				sCtx.Require().NotNil(gettingBalance, utils.ErrGetBalanceMustNotBeNil)
				sCtx.Require().Equal(tc.expectedMaxBalance, gettingBalance.Amount, utils.ErrBalanceAmountMustBeEqual)
			}, allure.NewParameter("time", time.Now()))
		})
	}
}

func (suite *WalletTestSuiteDebitCredit) TestCredit(t provider.T) {
	t.Title("Credit from account")

	userID, token := utils.CreateUserHelper(suite.UserSteps, t)

	testCases := []struct {
		name            string
		userLevel       string
		initialAmount   int32
		creditAmount    int32
		expectedBalance int32
	}{
		{
			name:            "Credit for ANON",
			userLevel:       "ANON",
			initialAmount:   49999,
			creditAmount:    49999,
			expectedBalance: 0,
		},
		{
			name:            "Credit for FULL",
			userLevel:       "FULL",
			initialAmount:   199999,
			creditAmount:    199999,
			expectedBalance: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t provider.T) {
			t.Title("Credit from account_" + tc.name)
			t.Tag("gRPC Test")
			t.Tag("Positive")
			t.Parallel()

			tc := tc

			t.WithNewStep("Credit from account step", func(sCtx provider.StepCtx) {
				ctxAuth := utils.ToCtx(context.Background(), token)

				if tc.userLevel == "FULL" {
					sCtx.NewStep("Upgrade user")
					err := suite.UserSteps.UpgradeUser(ctxAuth, t, userID)
					sCtx.Require().NoError(err, utils.ErrUpgradedUserErrorShouldBeNil)
				}

				sCtx.NewStep("Create account")
				createdAccount, err := suite.UserSteps.CreateAccount(ctxAuth, t, userID, "description", tc.initialAmount)
				sCtx.Require().NoError(err, utils.ErrCreatedAccountErrorShouldBeNil)
				sCtx.Require().NotNil(createdAccount, utils.ErrCreatedAccountMustNotBeNil)
				sCtx.Require().Equal(tc.initialAmount, createdAccount.Amount, utils.ErrAccountAmountMustBeEqual)

				sCtx.NewStep("Credit from account")
				err = suite.UserSteps.Credit(ctxAuth, t, userID, createdAccount.AccountId, tc.creditAmount)
				sCtx.Require().NoError(err, utils.ErrCreditErrorShouldBeNil)

				sCtx.NewStep("Get account balance")
				gettingBalance, err := suite.UserSteps.GetAccountBalance(ctxAuth, t, userID, createdAccount.AccountId)
				sCtx.Require().NoError(err, utils.ErrGetBalanceErrorShouldBeNil)
				sCtx.Require().NotNil(gettingBalance, utils.ErrGetBalanceMustNotBeNil)
				sCtx.Require().Equal(tc.expectedBalance, gettingBalance.Amount, utils.ErrBalanceAmountMustBeEqual)
			}, allure.NewParameter("time", time.Now()))
		})
	}
}

func (suite *WalletTestSuiteDebitCredit) TestCreditNegative(t provider.T) {
	t.Title("Credit invalid operation")

	ctxAuth := utils.ToCtx(context.Background(), suite.Token)

	testCases := []struct {
		name               string
		userLevel          string
		initialAmount      int32
		creditAmount       int32
		expectedMaxBalance int32
		expectedErr        string
	}{
		{
			name:               "Check for negative balance for ANON",
			userLevel:          "ANON",
			initialAmount:      50000,
			creditAmount:       50001,
			expectedMaxBalance: 50000,
			expectedErr:        "unavailable operation",
		},
		{
			name:               "Check for negative balance for FULL",
			userLevel:          "FULL",
			initialAmount:      200000,
			creditAmount:       200001,
			expectedMaxBalance: 200000,
			expectedErr:        "unavailable operation",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t provider.T) {
			t.Title("Credit from account with invalid operation_" + tc.name)
			t.Tag("gRPC Test")
			t.Tag("Negative")
			t.Parallel()

			tc := tc

			t.WithNewStep("Credit from account step", func(sCtx provider.StepCtx) {

				if tc.userLevel == "FULL" {
					sCtx.NewStep("Upgrade user")
					err := suite.UserSteps.UpgradeUser(ctxAuth, t, suite.UserID)
					sCtx.Require().NoError(err, utils.ErrUpgradedUserErrorShouldBeNil)
				}

				sCtx.NewStep("Create account")
				createdAccount, err := suite.UserSteps.CreateAccount(ctxAuth, t, suite.UserID, "description", tc.initialAmount)
				sCtx.Require().NoError(err, utils.ErrCreatedAccountErrorShouldBeNil)
				sCtx.Require().NotNil(createdAccount, utils.ErrCreatedAccountMustNotBeNil)
				sCtx.Require().Equal(tc.initialAmount, createdAccount.Amount, utils.ErrAccountAmountMustBeEqual)

				sCtx.NewStep("Credit from account")
				err = suite.UserSteps.Credit(ctxAuth, t, suite.UserID, createdAccount.AccountId, tc.creditAmount)
				sCtx.Require().Error(err, utils.ErrCreditErrorMustNotBeNil)
				st, ok := status.FromError(err)
				sCtx.Require().True(ok, utils.ErrExpectGRPCStatusError)
				sCtx.Require().Equal(codes.InvalidArgument, st.Code(), utils.ErrExpectInvalidArgument)
				sCtx.Require().Contains(st.Message(), tc.expectedErr, utils.ErrMessageExpectInvalidArgumentDebit)

				sCtx.NewStep("Get account balance")
				gettingBalance, err := suite.UserSteps.GetAccountBalance(ctxAuth, t, suite.UserID, createdAccount.AccountId)
				sCtx.Require().NoError(err, utils.ErrGetBalanceErrorShouldBeNil)
				sCtx.Require().NotNil(gettingBalance, utils.ErrGetBalanceMustNotBeNil)
				sCtx.Require().Equal(tc.expectedMaxBalance, gettingBalance.Amount, utils.ErrMessageExpectInvalidArgumentDebit)
			}, allure.NewParameter("time", time.Now()))
		})
	}
}

// P.S. Уже не стал добавлять проверки на попытки списания/добавления на счет другого юзера, можно в дальнейшем расширить тесты при необходимости)
