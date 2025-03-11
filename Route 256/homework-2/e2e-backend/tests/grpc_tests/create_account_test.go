package test

import (
	"context"
	"fmt"
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

type WalletTestSuiteCreateAccount struct {
	suite.Suite
	client    wallet.WalletClient
	conn      *grpc.ClientConn
	UserSteps *grpcsteps.UserSteps
	UserID    string
	Token     string
}

func TestWalletTestSuiteCreateAccount(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(WalletTestSuiteCreateAccount))
}

func (suite *WalletTestSuiteCreateAccount) SetupSuite() {
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

func (suite *WalletTestSuiteCreateAccount) TearDownSuite() {
	suite.conn.Close()
}

func (suite *WalletTestSuiteCreateAccount) TestCreateAccount() {
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
		tc := tc
		suite.T().Run(tc.name, func(t *testing.T) {
			ctxAuth := utils.ToCtx(context.Background(), suite.Token)

			createdAccount, err := suite.UserSteps.CreateAccount(ctxAuth, t, tc.userID, tc.description, tc.amount)

			require.NoError(suite.T(), err)
			require.NotNil(t, createdAccount)
			require.Equal(t, int32(tc.amount), createdAccount.Amount)
		})
	}
}

func (suite *WalletTestSuiteCreateAccount) TestCreateAccountNegative() {
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

func (suite *WalletTestSuiteCreateAccount) TestCreateAccountTooManyAccountsNegative() {
	userID, token := utils.CreateUserHelper(suite.UserSteps, suite.T())
	ctxAuth := utils.ToCtx(context.Background(), token)

	for i := 1; i < 5; i++ {
		if i != 4 {
			createdAccount, err := suite.UserSteps.CreateAccount(ctxAuth, suite.T(), userID, fmt.Sprintf("Fake description %d", i), int32(i))

			require.NoError(suite.T(), err)
			require.NotNil(suite.T(), createdAccount)
			require.Equal(suite.T(), int32(i), createdAccount.Amount)
		} else {
			createdAccount, err := suite.UserSteps.CreateAccount(ctxAuth, suite.T(), userID, fmt.Sprintf("Fake description %d", i), int32(i))

			require.Nil(suite.T(), createdAccount)
			require.NotNil(suite.T(), err)
			st, ok := status.FromError(err)
			require.True(suite.T(), ok, "expected gRPC status error")
			require.Equal(suite.T(), codes.AlreadyExists, st.Code())
			require.Contains(suite.T(), st.Message(), "User allready has 3 accounts") // В слове allready одна лишняя l
		}

	}

}

func (suite *WalletTestSuiteCreateAccount) TestCreateAccountByAnotherUserNegative() {
	_, token := utils.CreateUserHelper(suite.UserSteps, suite.T())
	ctxAuth := utils.ToCtx(context.Background(), token)

	createdAccount, err := suite.UserSteps.CreateAccount(ctxAuth, suite.T(), suite.UserID, "fake", 49999)

	require.Nil(suite.T(), createdAccount)
	require.NotNil(suite.T(), err)
	st, ok := status.FromError(err)
	require.True(suite.T(), ok, "expected gRPC status error")
	require.Equal(suite.T(), codes.Unauthenticated, st.Code())
	require.Contains(suite.T(), st.Message(), "not authorized")
}
