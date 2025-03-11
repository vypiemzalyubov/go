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

type WalletTestSuiteGetAccountBalance struct {
	suite.Suite
	client    wallet.WalletClient
	conn      *grpc.ClientConn
	UserSteps *grpcsteps.UserSteps
	UserID    string
	accountID string
	Token     string
}

func TestWalletTestSuiteGetAccountBalance(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(WalletTestSuiteGetAccountBalance))
}

func (suite *WalletTestSuiteGetAccountBalance) SetupSuite() {
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

	ctxAuth := utils.ToCtx(context.Background(), suite.Token)
	createdAccount, err := suite.UserSteps.CreateAccount(ctxAuth, suite.T(), createdUser.ID, "description", 1000)
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), createdAccount)
	suite.accountID = createdAccount.AccountId
}

func (suite *WalletTestSuiteGetAccountBalance) TearDownSuite() {
	suite.conn.Close()
}

func (suite *WalletTestSuiteGetAccountBalance) TestGetAccountBalance() {
	ctxAuth := utils.ToCtx(context.Background(), suite.Token)

	gettingBalance, err := suite.UserSteps.GetAccountBalance(ctxAuth, suite.T(), suite.UserID, suite.accountID)

	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), gettingBalance)
	require.Equal(suite.T(), int32(1000), gettingBalance.Amount)
}

func (suite *WalletTestSuiteGetAccountBalance) TestGetAccountBalanceNegative() {
	testCases := []struct {
		name        string
		userID      string
		accountID   string
		codes       codes.Code
		expectedErr string
		ctx         context.Context
	}{
		{
			name:        "Empty userID",
			userID:      "",
			accountID:   suite.accountID,
			expectedErr: "not authorized",
			codes:       codes.Unauthenticated,
			ctx:         utils.ToCtx(context.Background(), suite.Token),
		},
		{
			name:        "Invalid accountID",
			userID:      suite.UserID,
			accountID:   "fake account",
			expectedErr: "account not found",
			codes:       codes.NotFound,
			ctx:         utils.ToCtx(context.Background(), suite.Token),
		},
		{
			name:        "Invalid userID and accountID",
			userID:      "fake id",
			accountID:   "fake account",
			expectedErr: "not authorized",
			codes:       codes.Unauthenticated,
			ctx:         utils.ToCtx(context.Background(), suite.Token),
		},
		{
			name:        "Without authorization",
			userID:      suite.UserID,
			accountID:   suite.accountID,
			expectedErr: "not authorized",
			codes:       codes.Unauthenticated,
			ctx:         context.Background(),
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.T().Run(tc.name, func(t *testing.T) {
			gettingBalance, err := suite.UserSteps.GetAccountBalance(tc.ctx, t, tc.userID, tc.accountID)

			require.Nil(t, gettingBalance)
			require.NotNil(t, err)
			st, ok := status.FromError(err)
			require.True(t, ok, "expected gRPC status error")
			require.Equal(t, tc.codes, st.Code())
			require.Contains(t, st.Message(), tc.expectedErr)
		})
	}
}

func (suite *WalletTestSuiteGetAccountBalance) TestGetAccountBalanceByAnotherUserNegative() {
	_, token := utils.CreateUserHelper(suite.UserSteps, suite.T())
	ctxAuth := utils.ToCtx(context.Background(), token)

	gettingBalance, err := suite.UserSteps.GetAccountBalance(ctxAuth, suite.T(), suite.UserID, suite.accountID)

	require.Nil(suite.T(), gettingBalance)
	require.NotNil(suite.T(), err)
	st, ok := status.FromError(err)
	require.True(suite.T(), ok, "expected gRPC status error")
	require.Equal(suite.T(), codes.Unauthenticated, st.Code())
	require.Contains(suite.T(), st.Message(), "not authorized")
}
