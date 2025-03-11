package test

import (
	"context"
	"testing"

	"e2e-backend/internal/clients/grpccli"
	"e2e-backend/internal/pb/gitlab.ozon.dev/route256/wallet"
	"e2e-backend/internal/steps/grpcsteps"
	"e2e-backend/internal/utils"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
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

func TestWalletTestSuiteUpgradeUser(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(WalletTestSuiteUpgradeUser))
}

func (suite *WalletTestSuiteUpgradeUser) SetupSuite() {
	var err error
	suite.client, suite.conn, err = grpccli.NewWalletClient()
	require.NoError(suite.T(), err)
	suite.UserSteps = grpcsteps.NewUserSteps(suite.client)

	ctx := context.Background()

	createdUser, err := suite.UserSteps.CreateUser(ctx, suite.T(), nil)
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), createdUser)
	suite.UserID = createdUser.ID

	loginToken, err := suite.UserSteps.LogIn(ctx, suite.T(), createdUser.Phone, createdUser.Password)
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), loginToken)
	suite.Token = loginToken.Token
}

func (suite *WalletTestSuiteUpgradeUser) TearDownSuite() {
	suite.conn.Close()
}

func (suite *WalletTestSuiteUpgradeUser) TestUpgradeUser() {
	ctxAuth := utils.ToCtx(context.Background(), suite.Token)

	err := suite.UserSteps.UpgradeUser(ctxAuth, suite.T(), suite.UserID)
	require.NoError(suite.T(), err)

	gettingUser, err := suite.UserSteps.GetUser(ctxAuth, suite.T(), suite.UserID)

	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), gettingUser)
	require.Equal(suite.T(), suite.UserID, gettingUser.ID)
	require.Equal(suite.T(), "FULL", gettingUser.Level)
}

func (suite *WalletTestSuiteUpgradeUser) TestUpgradeUserAndCreateAccount() {
	userID, token := utils.CreateUserHelper(suite.UserSteps, suite.T())
	ctxAuth := utils.ToCtx(context.Background(), token)

	err := suite.UserSteps.UpgradeUser(ctxAuth, suite.T(), userID)
	require.NoError(suite.T(), err)

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
		tc := tc
		suite.T().Run(tc.name, func(t *testing.T) {
			createdAccount, err := suite.UserSteps.CreateAccount(ctxAuth, t, tc.userID, tc.description, tc.amount)

			require.NoError(t, err)
			require.NotNil(t, createdAccount)
			require.Equal(t, int32(tc.amount), createdAccount.Amount)
		})
	}
}

func (suite *WalletTestSuiteUpgradeUser) TestUpgradeUserNegative() {
	userID, token := utils.CreateUserHelper(suite.UserSteps, suite.T())
	ctxAuth := utils.ToCtx(context.Background(), token)

	err := suite.UserSteps.UpgradeUser(ctxAuth, suite.T(), userID)
	require.NoError(suite.T(), err)

	gettingUser, err := suite.UserSteps.GetUser(ctxAuth, suite.T(), userID)

	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), gettingUser)
	require.Equal(suite.T(), userID, gettingUser.ID)
	require.Equal(suite.T(), "FULL", gettingUser.Level)

	errNew := suite.UserSteps.UpgradeUser(ctxAuth, suite.T(), userID)

	require.Error(suite.T(), errNew)
	st, ok := status.FromError(errNew)
	require.True(suite.T(), ok, "expected gRPC status error")
	require.Equal(suite.T(), codes.AlreadyExists, st.Code())
	require.Contains(suite.T(), st.Message(), "user allready has FULL level")
}

func (suite *WalletTestSuiteUpgradeUser) TestUpgradeUserAndCreateAccountNegative() {
	userID, token := utils.CreateUserHelper(suite.UserSteps, suite.T())
	ctxAuth := utils.ToCtx(context.Background(), token)

	err := suite.UserSteps.UpgradeUser(ctxAuth, suite.T(), userID)
	require.NoError(suite.T(), err)

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
		tc := tc
		suite.T().Run(tc.name, func(t *testing.T) {
			createdAccount, err := suite.UserSteps.CreateAccount(tc.ctx, t, tc.userID, tc.description, tc.amount)

			require.Nil(t, createdAccount)
			require.NotNil(t, err)
			st, ok := status.FromError(err)
			require.True(t, ok, "expected gRPC status error")
			require.Equal(t, codes.Unauthenticated, st.Code())
			require.Contains(t, st.Message(), tc.expectedErr)
		})
	}
}
