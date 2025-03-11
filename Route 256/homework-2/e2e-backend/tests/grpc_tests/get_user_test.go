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

type WalletTestSuiteGetUser struct {
	suite.Suite
	client    wallet.WalletClient
	conn      *grpc.ClientConn
	UserSteps *grpcsteps.UserSteps
	UserID    string
	Token     string
}

func TestWalletTestSuiteGetUser(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(WalletTestSuiteGetUser))
}

func (suite *WalletTestSuiteGetUser) SetupSuite() {
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

func (suite *WalletTestSuiteGetUser) TearDownSuite() {
	suite.conn.Close()
}

func (suite *WalletTestSuiteGetUser) TestGetUser() {
	ctxAuth := utils.ToCtx(context.Background(), suite.Token)

	gettingUser, err := suite.UserSteps.GetUser(ctxAuth, suite.T(), suite.UserID)

	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), gettingUser)
	require.Equal(suite.T(), suite.UserID, gettingUser.ID)
}

func (suite *WalletTestSuiteGetUser) TestGetUserNegative() {
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
		tc := tc
		suite.T().Run(tc.name, func(t *testing.T) {
			ctxAuth := utils.ToCtx(context.Background(), tc.token)

			gettingUser, err := suite.UserSteps.GetUser(ctxAuth, t, tc.userId)

			require.Nil(t, gettingUser)
			require.NotNil(t, err)
			st, ok := status.FromError(err)
			require.True(t, ok, "expected gRPC status error")
			require.Equal(t, codes.Unauthenticated, st.Code())
			require.Contains(t, st.Message(), tc.expectedErr)
		})
	}
}

func (suite *WalletTestSuiteGetUser) TestGetUserInfoByAnotherUserNegative() {
	_, token := utils.CreateUserHelper(suite.UserSteps, suite.T())
	ctxAuth := utils.ToCtx(context.Background(), token)

	gettingUser, err := suite.UserSteps.GetUser(ctxAuth, suite.T(), suite.UserID)

	require.Nil(suite.T(), gettingUser)
	require.NotNil(suite.T(), err)

	st, ok := status.FromError(err)
	require.True(suite.T(), ok, "expected gRPC status error")
	require.Equal(suite.T(), codes.Unauthenticated, st.Code())
	require.Contains(suite.T(), st.Message(), "not authorized")
}
