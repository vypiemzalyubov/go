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

type WalletTestSuiteGetAccountBalance struct {
	suite.Suite
	client    wallet.WalletClient
	conn      *grpc.ClientConn
	UserSteps *grpcsteps.UserSteps
	UserID    string
	accountID string
	Token     string
}

func (suite *WalletTestSuiteGetAccountBalance) BeforeAll(t provider.T) {
	t.NewStep("Setup gRPC Client, Create User, Login User, Get Token, Create Account")

	var err error
	suite.client, suite.conn, err = grpccli.NewWalletClient()
	t.Require().NoError(err, utils.ErrCreatedGrpcClient)

	suite.UserSteps = grpcsteps.NewUserSteps(suite.client)

	ctx := context.Background()
	createdUser, _ := suite.UserSteps.CreateUser(ctx, t, nil)
	suite.UserID = createdUser.ID

	loginToken, _ := suite.UserSteps.LogIn(ctx, t, createdUser.Phone, createdUser.Password)
	suite.Token = loginToken.Token

	ctxAuth := utils.ToCtx(context.Background(), suite.Token)
	createdAccount, _ := suite.UserSteps.CreateAccount(ctxAuth, t, createdUser.ID, "description", 1000)
	suite.accountID = createdAccount.AccountId
}

func (suite *WalletTestSuiteGetAccountBalance) AfterAll(t provider.T) {
	t.NewStep("Teardown gRPC Client")
	suite.conn.Close()
}

func (suite *WalletTestSuiteGetAccountBalance) BeforeEach(t provider.T) {
	t.Epic("Wallet")
}

func TestWalletTestSuiteGetAccountBalanceRunner(t *testing.T) {
	t.Parallel()
	suite.RunSuite(t, new(WalletTestSuiteGetAccountBalance))
}

func (suite *WalletTestSuiteGetAccountBalance) TestGetAccountBalance(t provider.T) {
	t.Title("Get account balance")
	t.Tag("gRPC Test")
	t.Tag("Positive")

	ctxAuth := utils.ToCtx(context.Background(), suite.Token)

	t.WithNewStep("Get account balance step", func(sCtx provider.StepCtx) {
		gettingBalance, err := suite.UserSteps.GetAccountBalance(ctxAuth, t, suite.UserID, suite.accountID)

		sCtx.Require().NoError(err, utils.ErrGetBalanceErrorShouldBeNil)
		sCtx.Require().NotNil(gettingBalance, utils.ErrGetBalanceMustNotBeNil)
		sCtx.Require().Equal(int32(1000), gettingBalance.Amount, utils.ErrBalanceAmountMustBeEqual)
	}, allure.NewParameter("time", time.Now()))

}

func (suite *WalletTestSuiteGetAccountBalance) TestGetAccountBalanceNegative(t provider.T) {
	t.Title("Get invalid account balance")

	testCases := []struct {
		name        string
		userID      string
		accountID   string
		codes       codes.Code
		expectedErr string
		cnstCode    string
		cnstMessage string
		ctx         context.Context
	}{
		{
			name:        "Empty userID",
			userID:      "",
			accountID:   suite.accountID,
			expectedErr: "not authorized",
			codes:       codes.Unauthenticated,
			cnstCode:    utils.ErrExpectUnauthenticated,
			cnstMessage: utils.ErrMessageExpectUnauthenticated,
			ctx:         utils.ToCtx(context.Background(), suite.Token),
		},
		{
			name:        "Invalid accountID",
			userID:      suite.UserID,
			accountID:   "fake account",
			expectedErr: "account not found",
			codes:       codes.NotFound,
			cnstCode:    utils.ErrExpectNotFound,
			cnstMessage: utils.ErrMessageExpectNotFound,
			ctx:         utils.ToCtx(context.Background(), suite.Token),
		},
		{
			name:        "Invalid userID and accountID",
			userID:      "fake id",
			accountID:   "fake account",
			expectedErr: "not authorized",
			codes:       codes.Unauthenticated,
			cnstCode:    utils.ErrExpectUnauthenticated,
			cnstMessage: utils.ErrMessageExpectUnauthenticated,
			ctx:         utils.ToCtx(context.Background(), suite.Token),
		},
		{
			name:        "Without authorization",
			userID:      suite.UserID,
			accountID:   suite.accountID,
			expectedErr: "not authorized",
			codes:       codes.Unauthenticated,
			cnstCode:    utils.ErrExpectUnauthenticated,
			cnstMessage: utils.ErrMessageExpectUnauthenticated,
			ctx:         context.Background(),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t provider.T) {
			t.Title("Get invalid account balance_" + tc.name)
			t.Tag("gRPC Test")
			t.Tag("Negative")
			t.Parallel()

			tc := tc

			t.WithNewStep("Get account balance step", func(sCtx provider.StepCtx) {
				gettingBalance, err := suite.UserSteps.GetAccountBalance(tc.ctx, t, tc.userID, tc.accountID)

				sCtx.Require().Nil(gettingBalance, utils.ErrGetBalanceShouldBeNil)
				sCtx.Require().NotNil(err, utils.ErrGetBalanceErrorMustNotBeNil)
				st, ok := status.FromError(err)
				sCtx.Require().True(ok, utils.ErrExpectGRPCStatusError)
				sCtx.Require().Equal(tc.codes, st.Code(), tc.cnstCode)
				sCtx.Require().Contains(st.Message(), tc.expectedErr, tc.cnstMessage)
			}, allure.NewParameter("time", time.Now()))
		})
	}

}

func (suite *WalletTestSuiteGetAccountBalance) TestGetAccountBalanceByAnotherUserNegative(t provider.T) {
	t.Title("Get account balance by another user")
	t.Tag("gRPC Test")
	t.Tag("Negative")

	_, token := utils.CreateUserHelper(suite.UserSteps, t)
	ctxAuth := utils.ToCtx(context.Background(), token)

	t.WithNewStep("Get account balance step", func(sCtx provider.StepCtx) {
		gettingBalance, err := suite.UserSteps.GetAccountBalance(ctxAuth, t, suite.UserID, suite.accountID)

		sCtx.Require().Nil(gettingBalance, utils.ErrGetBalanceShouldBeNil)
		sCtx.Require().NotNil(err, utils.ErrGetBalanceErrorMustNotBeNil)
		st, ok := status.FromError(err)
		sCtx.Require().True(ok, utils.ErrExpectGRPCStatusError)
		sCtx.Require().Equal(codes.Unauthenticated, st.Code(), utils.ErrExpectUnauthenticated)
		sCtx.Require().Contains(st.Message(), "not authorized", utils.ErrMessageExpectUnauthenticated)
	}, allure.NewParameter("time", time.Now()))
}
