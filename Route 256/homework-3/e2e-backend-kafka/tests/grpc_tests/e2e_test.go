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
)

type WalletTestSuiteE2E struct {
	suite.Suite
	client    wallet.WalletClient
	conn      *grpc.ClientConn
	UserSteps *grpcsteps.UserSteps
}

func (suite *WalletTestSuiteE2E) BeforeAll(t provider.T) {
	t.NewStep("Setup gRPC Client")

	var err error
	suite.client, suite.conn, err = grpccli.NewWalletClient()
	t.Require().NoError(err, utils.ErrCreatedGrpcClient)

	suite.UserSteps = grpcsteps.NewUserSteps(suite.client)
}

func (suite *WalletTestSuiteE2E) AfterAll(t provider.T) {
	t.NewStep("Teardown gRPC Client")
	suite.conn.Close()
}

func (suite *WalletTestSuiteE2E) BeforeEach(t provider.T) {
	t.Epic("Wallet")
}

func TestWalletTestSuiteE2ERunner(t *testing.T) {
	t.Parallel()
	suite.RunSuite(t, new(WalletTestSuiteE2E))
}

func (suite *WalletTestSuiteE2E) TestE2E(t provider.T) {
	t.Title("E2E test")
	t.Tag("gRPC Test")
	t.Tag("E2E")

	t.WithNewStep("E2E step", func(sCtx provider.StepCtx) {

		sCtx.NewStep("Create user")
		ctx := context.Background()
		createdUserAnon, err := suite.UserSteps.CreateUser(ctx, t, nil)
		sCtx.Require().NoError(err, utils.ErrCreatedUserErrorShouldBeNil)
		sCtx.Require().NotNil(createdUserAnon, utils.ErrCreatedUserMustNotBeNil)
		sCtx.Require().Equal("ANON", createdUserAnon.Level, utils.ErrCreatedUserLevelShouldBeAnon)

		sCtx.NewStep("Login user")
		loginToken, err := suite.UserSteps.LogIn(ctx, t, createdUserAnon.Phone, createdUserAnon.Password)
		sCtx.Require().NoError(err, utils.ErrLoginUserErrorShouldBeNil)
		sCtx.Require().NotNil(loginToken, utils.ErrLoginUserMustNotBeNil)

		sCtx.NewStep("Authorize user")
		ctxAuth := utils.ToCtx(context.Background(), loginToken.Token)

		sCtx.NewStep("Get user")
		gettingUserAnon, err := suite.UserSteps.GetUser(ctxAuth, t, createdUserAnon.ID)
		sCtx.Require().NoError(err, utils.ErrGetUserErrorShouldBeNil)
		sCtx.Require().NotNil(gettingUserAnon, utils.ErrGetUserMustNotBeNil)
		sCtx.Require().Equal(createdUserAnon.ID, gettingUserAnon.ID, utils.ErrUserIdMustBeEqual)

		sCtx.NewStep("Create account")
		createdAccountAnon, err := suite.UserSteps.CreateAccount(ctxAuth, t, createdUserAnon.ID, "e2e description", int32(25000))
		sCtx.Require().NoError(err, utils.ErrCreatedAccountErrorShouldBeNil)
		sCtx.Require().NotNil(createdAccountAnon, utils.ErrCreatedAccountMustNotBeNil)
		sCtx.Require().Equal(int32(25000), createdAccountAnon.Amount, utils.ErrAccountAmountMustBeEqual)

		sCtx.NewStep("Get account balance")
		gettingBalanceAnon, err := suite.UserSteps.GetAccountBalance(ctxAuth, t, createdUserAnon.ID, createdAccountAnon.AccountId)
		sCtx.Require().NoError(err, utils.ErrGetBalanceErrorShouldBeNil)
		sCtx.Require().NotNil(gettingBalanceAnon, utils.ErrGetBalanceMustNotBeNil)
		sCtx.Require().Equal(int32(25000), gettingBalanceAnon.Amount, utils.ErrBalanceAmountMustBeEqual)

		sCtx.NewStep("Debit to account balance")
		debitErrAnon := suite.UserSteps.Debit(ctxAuth, t, createdUserAnon.ID, createdAccountAnon.AccountId, int32(10000))
		sCtx.Require().NoError(debitErrAnon, utils.ErrDebitErrorShouldBeNil)

		sCtx.NewStep("Get account balance")
		gettingBalanceAfterDebitAnon, err := suite.UserSteps.GetAccountBalance(ctxAuth, t, createdUserAnon.ID, createdAccountAnon.AccountId)
		sCtx.Require().NoError(err, utils.ErrGetBalanceErrorShouldBeNil)
		sCtx.Require().NotNil(gettingBalanceAfterDebitAnon, utils.ErrGetBalanceMustNotBeNil)
		sCtx.Require().Equal(int32(35000), gettingBalanceAfterDebitAnon.Amount, utils.ErrBalanceAmountMustBeEqual)

		sCtx.NewStep("Credit from account balance")
		creditErrAnon := suite.UserSteps.Credit(ctxAuth, t, createdUserAnon.ID, createdAccountAnon.AccountId, int32(35000))
		sCtx.Require().NoError(creditErrAnon, utils.ErrCreditErrorShouldBeNil)

		sCtx.NewStep("Get account balance")
		gettingBalanceAfterCreditAnon, err := suite.UserSteps.GetAccountBalance(ctxAuth, t, createdUserAnon.ID, createdAccountAnon.AccountId)
		sCtx.Require().NoError(err, utils.ErrGetBalanceErrorShouldBeNil)
		sCtx.Require().NotNil(gettingBalanceAfterCreditAnon, utils.ErrGetBalanceMustNotBeNil)
		sCtx.Require().Equal(int32(0), gettingBalanceAfterCreditAnon.Amount, utils.ErrBalanceAmountMustBeEqual)

		sCtx.NewStep("Upgrade user")
		errUpgrade := suite.UserSteps.UpgradeUser(ctxAuth, t, createdUserAnon.ID)
		sCtx.Require().NoError(errUpgrade, utils.ErrUpgradedUserErrorShouldBeNil)

		sCtx.NewStep("Get upgraded user")
		gettingUserFull, err := suite.UserSteps.GetUser(ctxAuth, t, createdUserAnon.ID)
		sCtx.Require().NoError(err, utils.ErrGetUserErrorShouldBeNil)
		sCtx.Require().NotNil(gettingUserFull, utils.ErrGetUserMustNotBeNil)
		sCtx.Require().Equal(createdUserAnon.ID, gettingUserFull.ID, utils.ErrUserIdMustBeEqual)
		sCtx.Require().Equal("FULL", gettingUserFull.Level, utils.ErrUpgradedUserLevelShouldBeFull)

		sCtx.NewStep("Debit to account balance")
		debitErrFull := suite.UserSteps.Debit(ctxAuth, t, createdUserAnon.ID, createdAccountAnon.AccountId, int32(200000))
		sCtx.Require().NoError(debitErrFull, utils.ErrDebitErrorShouldBeNil)

		sCtx.NewStep("Get account balance")
		gettingBalanceAfterDebitFull, err := suite.UserSteps.GetAccountBalance(ctxAuth, t, createdUserAnon.ID, createdAccountAnon.AccountId)
		sCtx.Require().NoError(err, utils.ErrGetBalanceErrorShouldBeNil)
		sCtx.Require().NotNil(gettingBalanceAfterDebitFull, utils.ErrGetBalanceMustNotBeNil)
		sCtx.Require().Equal(int32(200000), gettingBalanceAfterDebitFull.Amount, utils.ErrBalanceAmountMustBeEqual)

		sCtx.NewStep("Credit from account balance")
		creditErrFull := suite.UserSteps.Credit(ctxAuth, t, createdUserAnon.ID, createdAccountAnon.AccountId, int32(199999))
		sCtx.Require().NoError(creditErrFull, utils.ErrCreditErrorShouldBeNil)

		sCtx.NewStep("Get account balance")
		gettingBalanceAfterCreditFull, err := suite.UserSteps.GetAccountBalance(ctxAuth, t, createdUserAnon.ID, createdAccountAnon.AccountId)
		sCtx.Require().NoError(err, utils.ErrGetBalanceErrorShouldBeNil)
		sCtx.Require().NotNil(gettingBalanceAfterCreditFull, utils.ErrGetBalanceMustNotBeNil)
		sCtx.Require().Equal(int32(1), gettingBalanceAfterCreditFull.Amount, utils.ErrBalanceAmountMustBeEqual)
	}, allure.NewParameter("time", time.Now()))
}
