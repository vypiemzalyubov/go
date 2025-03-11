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

type WalletTestSuiteGetUser struct {
	suite.Suite
	client    wallet.WalletClient
	conn      *grpc.ClientConn
	UserSteps *grpcsteps.UserSteps
	UserID    string
	Token     string
}

func (suite *WalletTestSuiteGetUser) BeforeAll(t provider.T) {
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

func (suite *WalletTestSuiteGetUser) AfterAll(t provider.T) {
	t.NewStep("Teardown gRPC Client")
	suite.conn.Close()
}

func (suite *WalletTestSuiteGetUser) BeforeEach(t provider.T) {
	t.Epic("Wallet")
}

func TestWalletTestSuiteGetUserRunner(t *testing.T) {
	t.Parallel()
	suite.RunSuite(t, new(WalletTestSuiteGetUser))
}

func (suite *WalletTestSuiteGetUser) TestGetUser(t provider.T) {
	t.Title("Get user")
	t.Tag("gRPC Test")
	t.Tag("Positive")

	ctxAuth := utils.ToCtx(context.Background(), suite.Token)

	t.WithNewStep("Get user step", func(sCtx provider.StepCtx) {
		gettingUser, err := suite.UserSteps.GetUser(ctxAuth, t, suite.UserID)

		sCtx.Require().NoError(err, utils.ErrGetUserErrorShouldBeNil)
		sCtx.Require().NotNil(gettingUser, utils.ErrGetUserMustNotBeNil)
		sCtx.Require().Equal(suite.UserID, gettingUser.ID, utils.ErrUserIdMustBeEqual)
	}, allure.NewParameter("time", time.Now()))

}

func (suite *WalletTestSuiteGetUser) TestGetUserNegative(t provider.T) {
	t.Title("Get invalid user")

	testCases := []struct {
		name        string
		userId      string
		token       string
		expectedErr string
	}{
		{
			name:        "Non-existent user without auth",
			userId:      "123",
			token:       "notoken1",
			expectedErr: "not authorized",
		},
		{
			name:        "Non-existent user with auth",
			userId:      "123",
			token:       suite.Token,
			expectedErr: "not authorized", // Мне кажется, здесь лучше подошло бы 404 (как и в тесте выше), хотя и 401 имеет место быть. Остается на откуп разрабов=)
		},
		{
			name:        "Existing user without auth",
			userId:      suite.UserID,
			token:       "notoken2",
			expectedErr: "not authorized",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t provider.T) {
			t.Title("Get invalid user_" + tc.name)
			t.Tag("gRPC Test")
			t.Tag("Negative")
			t.Parallel()

			tc := tc

			t.WithNewStep("Get user step", func(sCtx provider.StepCtx) {
				ctxAuth := utils.ToCtx(context.Background(), tc.token)
				gettingUser, err := suite.UserSteps.GetUser(ctxAuth, t, tc.userId)

				sCtx.Require().Nil(gettingUser, utils.ErrGetUserShouldBeNil)
				sCtx.Require().NotNil(err, utils.ErrGetUserErrorMustNotBeNil)
				st, ok := status.FromError(err)
				sCtx.Require().True(ok, utils.ErrExpectGRPCStatusError)
				sCtx.Require().Equal(codes.Unauthenticated, st.Code(), utils.ErrExpectUnauthenticated)
				sCtx.Require().Contains(st.Message(), tc.expectedErr, utils.ErrMessageExpectUnauthenticated)
			}, allure.NewParameter("time", time.Now()))
		})
	}

}

func (suite *WalletTestSuiteGetUser) TestGetUserInfoByAnotherUserNegative(t provider.T) {
	t.Title("Get user by another user")
	t.Tag("gRPC Test")
	t.Tag("Negative")

	_, token := utils.CreateUserHelper(suite.UserSteps, t)
	ctxAuth := utils.ToCtx(context.Background(), token)

	t.WithNewStep("Get user step", func(sCtx provider.StepCtx) {
		gettingUser, err := suite.UserSteps.GetUser(ctxAuth, t, suite.UserID)

		sCtx.Require().Nil(gettingUser, utils.ErrGetUserShouldBeNil)
		sCtx.Require().NotNil(err, utils.ErrGetUserErrorMustNotBeNil)
		st, ok := status.FromError(err)
		sCtx.Require().True(ok, utils.ErrExpectGRPCStatusError)
		sCtx.Require().Equal(codes.Unauthenticated, st.Code(), utils.ErrExpectUnauthenticated)
		sCtx.Require().Contains(st.Message(), "not authorized", utils.ErrMessageExpectUnauthenticated)
	}, allure.NewParameter("time", time.Now()))
}
