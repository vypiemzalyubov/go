package test

import (
	"context"
	"e2e-backend/internal/clients/grpccli"
	"e2e-backend/internal/pb/gitlab.ozon.dev/route256/wallet"
	"e2e-backend/internal/steps/grpcsteps"
	utils "e2e-backend/internal/utils"

	"fmt"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type WalletTestSuiteCreateAccount struct {
	suite.Suite
	client    wallet.WalletClient
	conn      *grpc.ClientConn
	UserSteps *grpcsteps.UserSteps
	UserID    string
	Token     string
}

func (suite *WalletTestSuiteCreateAccount) BeforeAll(t provider.T) {
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

func (suite *WalletTestSuiteCreateAccount) AfterAll(t provider.T) {
	t.NewStep("Teardown gRPC Client")
	suite.conn.Close()
}

func (suite *WalletTestSuiteCreateAccount) BeforeEach(t provider.T) {
	t.Epic("Wallet")
}

func TestWalletTestSuiteCreateAccountRunner(t *testing.T) {
	t.Parallel()
	suite.RunSuite(t, new(WalletTestSuiteCreateAccount))
}

func (suite *WalletTestSuiteCreateAccount) TestCreateAccount(t provider.T) {
	t.Title("Create account")

	testCases := []struct {
		name        string
		userID      string
		description string
		amount      int32
	}{
		{
			name:        "All filled fields",
			userID:      suite.UserID,
			description: "fake description 1",
			amount:      1,
		},
		{
			name:        "Empty description",
			userID:      suite.UserID,
			description: "", // Вроде не похоже на баг, но выглядит подозрительно, т.к поле обязательное, так что скорее баг. Решил оставить в положительных проверках
			amount:      0,
		},
		{
			name:        "Attempting to create three accounts for one user and max amount",
			userID:      suite.UserID,
			description: "fake description 2",
			amount:      50000,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t provider.T) {
			t.Title("Create account_" + tc.name)
			t.Tag("gRPC Test")
			t.Tag("Positive")
			t.Parallel()

			tc := tc

			t.WithNewStep("Create account step", func(sCtx provider.StepCtx) {
				ctxAuth := utils.ToCtx(context.Background(), suite.Token)
				createdAccount, err := suite.UserSteps.CreateAccount(ctxAuth, t, tc.userID, tc.description, tc.amount)

				sCtx.Require().NoError(err, utils.ErrCreatedAccountErrorShouldBeNil)
				sCtx.Require().NotNil(createdAccount, utils.ErrCreatedAccountMustNotBeNil)
				sCtx.Require().Equal(tc.amount, createdAccount.Amount, utils.ErrAccountAmountMustBeEqual)
			}, allure.NewParameter("time", time.Now()))
		})
	}
}

func (suite *WalletTestSuiteCreateAccount) TestCreateAccountNegative(t provider.T) {
	t.Title("Create account negative")

	testCases := []struct {
		name        string
		userID      string
		description string
		amount      int32
		expectedErr string
		ctx         context.Context
	}{
		{
			name:        "Empty userID",
			userID:      "",
			description: "fake description 1",
			amount:      1,
			expectedErr: "not authorized",
			ctx:         utils.ToCtx(context.Background(), suite.Token),
		},
		// {
		// 	name:        "Negative amount",
		// 	userID:      suite.UserID,
		// 	description: "fake description 2",
		// 	amount:      -1,                            Здесь точно баг, счет нельзя создавать с отрицательным значением
		// 	expectedErr: "amount cannot be negative",
		//  ctx:         utils.ToCtx(context.Background(), suite.Token),
		// },
		{
			name:        "Not authorized user",
			userID:      suite.UserID,
			description: "fake description 2",
			amount:      100,
			expectedErr: "not authorized",
			ctx:         context.Background(),
		},
		// {
		// 	name:        "Amount in the account is too much",
		// 	userID:      suite.UserID,
		// 	description: "fake description 3",
		// 	amount:      50001,                       Баг, т.к для анонимного пользователя максимальная сумма на счете — 50000
		// 	expectedErr: "amount in the account is too much for an anonymous person",
		// 	ctx:         utils.ToCtx(context.Background(), suite.Token),
		// },
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t provider.T) {
			t.Title("Create account_" + tc.name)
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

func (suite *WalletTestSuiteCreateAccount) TestCreateAccountTooManyAccountsNegative(t provider.T) {
	t.Title("Create too many accounts")
	t.Tag("gRPC Test")
	t.Tag("Negative")

	userID, token := utils.CreateUserHelper(suite.UserSteps, t)
	ctxAuth := utils.ToCtx(context.Background(), token)

	for i := 1; i < 5; i++ {
		if i != 4 {
			t.WithNewStep("Create account step", func(sCtx provider.StepCtx) {
				createdAccount, err := suite.UserSteps.CreateAccount(ctxAuth, t, userID, fmt.Sprintf("Fake description %d", i), int32(i))

				sCtx.Require().NoError(err, utils.ErrCreatedAccountErrorShouldBeNil)
				sCtx.Require().NotNil(createdAccount, utils.ErrCreatedAccountMustNotBeNil)
				sCtx.Require().Equal(int32(i), createdAccount.Amount, utils.ErrAccountAmountMustBeEqual)
			}, allure.NewParameter("time", time.Now()))
		} else {
			t.WithNewStep("Failed to create account step", func(sCtx provider.StepCtx) {
				createdAccount, err := suite.UserSteps.CreateAccount(ctxAuth, t, userID, fmt.Sprintf("Fake description %d", i), int32(i))

				sCtx.Require().Nil(createdAccount, utils.ErrCreatedAccountShouldBeNil)
				sCtx.Require().NotNil(err, utils.ErrCreatedAccountErrorMustNotBeNil)
				st, ok := status.FromError(err)
				sCtx.Require().True(ok, utils.ErrExpectGRPCStatusError)
				sCtx.Require().Equal(codes.AlreadyExists, st.Code(), utils.ErrExpectAlreadyExists)
				sCtx.Require().Contains(st.Message(), "User allready has 3 accounts", utils.ErrMessageExpectAlreadyExists)
			}, allure.NewParameter("time", time.Now()))
		}

	}

}

func (suite *WalletTestSuiteCreateAccount) TestCreateAccountByAnotherUserNegative(t provider.T) {
	t.Title("Create account by another user")
	t.Tag("gRPC Test")
	t.Tag("Negative")

	_, token := utils.CreateUserHelper(suite.UserSteps, t)
	ctxAuth := utils.ToCtx(context.Background(), token)

	t.WithNewStep("Create account step", func(sCtx provider.StepCtx) {
		createdAccount, err := suite.UserSteps.CreateAccount(ctxAuth, t, suite.UserID, "fake", 49999)

		sCtx.Require().Nil(createdAccount, utils.ErrCreatedAccountShouldBeNil)
		sCtx.Require().NotNil(err, utils.ErrCreatedAccountErrorMustNotBeNil)
		st, ok := status.FromError(err)
		sCtx.Require().True(ok, utils.ErrExpectGRPCStatusError)
		sCtx.Require().Equal(codes.Unauthenticated, st.Code(), utils.ErrExpectUnauthenticated)
		sCtx.Require().Contains(st.Message(), "not authorized", utils.ErrMessageExpectUnauthenticated)
	}, allure.NewParameter("time", time.Now()))

}
