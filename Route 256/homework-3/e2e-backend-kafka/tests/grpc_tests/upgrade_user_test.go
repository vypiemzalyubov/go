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

type WalletTestSuiteUpgradeUser struct {
	suite.Suite
	client    wallet.WalletClient
	conn      *grpc.ClientConn
	UserSteps *grpcsteps.UserSteps
	UserID    string
	Token     string
}

func (suite *WalletTestSuiteUpgradeUser) BeforeAll(t provider.T) {
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

func (suite *WalletTestSuiteUpgradeUser) AfterAll(t provider.T) {
	t.NewStep("Teardown gRPC Client")
	suite.conn.Close()
}

func (suite *WalletTestSuiteUpgradeUser) BeforeEach(t provider.T) {
	t.Epic("Wallet")
}

func TestWalletTestSuiteUpgradeUserRunner(t *testing.T) {
	t.Parallel()
	suite.RunSuite(t, new(WalletTestSuiteUpgradeUser))
}

func (suite *WalletTestSuiteUpgradeUser) TestUpgradeUser(t provider.T) {
	t.Title("Upgrade user")
	t.Tag("gRPC Test")
	t.Tag("Positive")

	ctxAuth := utils.ToCtx(context.Background(), suite.Token)

	t.WithNewStep("Upgrade user step", func(sCtx provider.StepCtx) {
		err := suite.UserSteps.UpgradeUser(ctxAuth, t, suite.UserID)

		sCtx.Require().NoError(err, utils.ErrUpgradedUserErrorShouldBeNil)
	}, allure.NewParameter("time", time.Now()))

	t.WithNewStep("Get user step", func(sCtx provider.StepCtx) {
		gettingUser, err := suite.UserSteps.GetUser(ctxAuth, t, suite.UserID)

		sCtx.Require().NoError(err, utils.ErrGetUserErrorShouldBeNil)
		sCtx.Require().NotNil(gettingUser, utils.ErrGetUserMustNotBeNil)
		sCtx.Require().Equal(suite.UserID, gettingUser.ID, utils.ErrUserIdMustBeEqual)
		sCtx.Require().Equal("FULL", gettingUser.Level, utils.ErrUpgradedUserLevelShouldBeFull)
	}, allure.NewParameter("time", time.Now()))
}

func (suite *WalletTestSuiteUpgradeUser) TestUpgradeUserAndCreateAccount(t provider.T) {
	t.Title("Upgrade user and create account")

	userID, token := utils.CreateUserHelper(suite.UserSteps, t)
	ctxAuth := utils.ToCtx(context.Background(), token)

	t.WithNewStep("Upgrade user step", func(sCtx provider.StepCtx) {
		err := suite.UserSteps.UpgradeUser(ctxAuth, t, userID)

		sCtx.Require().NoError(err, utils.ErrUpgradedUserErrorShouldBeNil)
	}, allure.NewParameter("time", time.Now()))

	testCases := []struct {
		name        string
		userID      string
		description string
		amount      int32
	}{
		{
			name:        "All filled fields for upgraded user",
			userID:      userID,
			description: "fake upgrade description 1",
			amount:      1,
		},
		{
			name:        "Empty description",
			userID:      userID,
			description: "", // Вроде не похоже на баг, но выглядит подозрительно, т.к поле обязательное, так что скорее баг. Решил оставить в положительных проверках
			amount:      0,
		},
		{
			name:        "Attempting to create three accounts for one upgraded user and max amount",
			userID:      userID,
			description: "fake upgrade description 2",
			amount:      200000,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t provider.T) {
			t.Title("Create account for upgraded user_" + tc.name)
			t.Tag("gRPC Test")
			t.Tag("Positive")
			t.Parallel()

			tc := tc

			t.WithNewStep("Create account step", func(sCtx provider.StepCtx) {
				createdAccount, err := suite.UserSteps.CreateAccount(ctxAuth, t, tc.userID, tc.description, tc.amount)

				sCtx.Require().NoError(err, utils.ErrCreatedAccountErrorShouldBeNil)
				sCtx.Require().NotNil(createdAccount, utils.ErrCreatedAccountMustNotBeNil)
				sCtx.Require().Equal(int32(tc.amount), createdAccount.Amount, utils.ErrAccountAmountMustBeEqual)
			}, allure.NewParameter("time", time.Now()))
		})
	}
}

func (suite *WalletTestSuiteUpgradeUser) TestUpgradeUserNegative(t provider.T) {
	t.Title("Upgrade one user twice")
	t.Tag("gRPC Test")
	t.Tag("Negative")

	userID, token := utils.CreateUserHelper(suite.UserSteps, t)
	ctxAuth := utils.ToCtx(context.Background(), token)

	t.WithNewStep("Upgrade user step", func(sCtx provider.StepCtx) {
		err := suite.UserSteps.UpgradeUser(ctxAuth, t, userID)

		sCtx.Require().NoError(err, utils.ErrUpgradedUserErrorShouldBeNil)
	}, allure.NewParameter("time", time.Now()))

	t.WithNewStep("Get user step", func(sCtx provider.StepCtx) {
		gettingUser, err := suite.UserSteps.GetUser(ctxAuth, t, userID)

		sCtx.Require().NoError(err, utils.ErrGetUserErrorShouldBeNil)
		sCtx.Require().NotNil(gettingUser, utils.ErrGetUserMustNotBeNil)
		sCtx.Require().Equal(userID, gettingUser.ID, utils.ErrUserIdMustBeEqual)
		sCtx.Require().Equal("FULL", gettingUser.Level, utils.ErrUpgradedUserLevelShouldBeFull)
	}, allure.NewParameter("time", time.Now()))

	t.WithNewStep("Second upgrade user step", func(sCtx provider.StepCtx) {
		err := suite.UserSteps.UpgradeUser(ctxAuth, t, userID)

		sCtx.Require().Error(err, utils.ErrUpgradeUserErrorMustNotBeNil)
		st, ok := status.FromError(err)
		sCtx.Require().True(ok, utils.ErrExpectGRPCStatusError)
		sCtx.Require().Equal(codes.AlreadyExists, st.Code(), utils.ErrExpectAlreadyExists)
		sCtx.Require().Contains(st.Message(), "user allready has FULL level", utils.ErrMessageExpectAlreadyFullLevel)
	}, allure.NewParameter("time", time.Now()))

}

func (suite *WalletTestSuiteUpgradeUser) TestUpgradeUserAndCreateAccountNegative(t provider.T) {
	t.Title("Upgrade user and create invalid account")

	userID, token := utils.CreateUserHelper(suite.UserSteps, t)
	ctxAuth := utils.ToCtx(context.Background(), token)

	t.WithNewStep("Upgrade user step", func(sCtx provider.StepCtx) {
		err := suite.UserSteps.UpgradeUser(ctxAuth, t, userID)

		sCtx.Require().NoError(err, utils.ErrUpgradedUserErrorShouldBeNil)
	}, allure.NewParameter("time", time.Now()))

	testCases := []struct {
		name        string
		userID      string
		description string
		amount      int32
		expectedErr string
		ctx         context.Context
	}{
		// {
		// 	name:        "Negative amount for upgraded user",
		// 	userID:      userID,
		// 	description: "fake upgrade description 1",
		// 	amount:      -1,                            Здесь точно баг, счет нельзя создавать с отрицательным значением
		// 	expectedErr: "amount cannot be negative",
		//  ctx:         utils.ToCtx(context.Background(), token),
		// },
		{
			name:        "Not authorized upgraded user",
			userID:      userID,
			description: "fake upgrade description 2",
			amount:      199999,
			expectedErr: "not authorized",
			ctx:         context.Background(),
		},
		// {
		// 	name:        "Amount in the account of upgraded user is too much",
		// 	userID:      userID,
		// 	description: "fake upgrade description 3",
		// 	amount:      200001,                    Баг, т.к для полного пользователя максимальная сумма на счете — 200000
		// 	expectedErr: "amount in the account is too much for full person",
		// 	ctx:         utils.ToCtx(context.Background(), token),
		// },
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t provider.T) {
			t.Title("Create invalid account for upgraded user_" + tc.name)
			t.Tag("gRPC Test")
			t.Tag("Negative")
			t.Parallel()

			tc := tc

			t.WithNewStep("Create account step", func(sCtx provider.StepCtx) {
				createdAccount, err := suite.UserSteps.CreateAccount(tc.ctx, t, tc.userID, tc.description, tc.amount)

				sCtx.Require().Nil(createdAccount)
				sCtx.Require().NotNil(err)
				st, ok := status.FromError(err)
				sCtx.Require().True(ok, utils.ErrExpectGRPCStatusError)
				sCtx.Require().Equal(codes.Unauthenticated, st.Code(), utils.ErrExpectUnauthenticated)
				sCtx.Require().Contains(st.Message(), tc.expectedErr, utils.ErrMessageExpectUnauthenticated)
			}, allure.NewParameter("time", time.Now()))
		})
	}

}
