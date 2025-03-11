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

type WalletTestSuiteDebitCredit struct {
	suite.Suite
	client    wallet.WalletClient
	conn      *grpc.ClientConn
	UserSteps *grpcsteps.UserSteps
	UserID    string
	Token     string
}

func TestWalletTestSuiteDebitCredit(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(WalletTestSuiteDebitCredit))
}

func (suite *WalletTestSuiteDebitCredit) SetupSuite() {
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

func (suite *WalletTestSuiteDebitCredit) TearDownSuite() {
	suite.conn.Close()
}

func (suite *WalletTestSuiteDebitCredit) TestDebit() {
	userID, token := utils.CreateUserHelper(suite.UserSteps, suite.T())

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
		tc := tc
		suite.T().Run(tc.name, func(t *testing.T) {
			ctxAuth := utils.ToCtx(context.Background(), token)

			if tc.userLevel == "FULL" {
				err := suite.UserSteps.UpgradeUser(ctxAuth, t, userID)
				require.NoError(t, err)
			}

			createdAccount, err := suite.UserSteps.CreateAccount(ctxAuth, t, userID, "description", tc.initialAmount)

			require.NoError(suite.T(), err)
			require.NotNil(t, createdAccount)
			require.Equal(t, tc.initialAmount, createdAccount.Amount)

			err = suite.UserSteps.Debit(ctxAuth, t, userID, createdAccount.AccountId, tc.debitAmount)

			require.NoError(t, err)

			gettingBalance, err := suite.UserSteps.GetAccountBalance(ctxAuth, t, userID, createdAccount.AccountId)

			require.NoError(suite.T(), err)
			require.NotNil(suite.T(), gettingBalance)
			require.Equal(suite.T(), tc.expectedBalance, gettingBalance.Amount)
		})
	}
}

func (suite *WalletTestSuiteDebitCredit) TestDebitNegative() {
	userID, token := utils.CreateUserHelper(suite.UserSteps, suite.T())

	testCases := []struct {
		name               string
		userLevel          string
		initialAmount      int32
		debitAmount        int32
		expectedMaxBalance int32
		expectedErr        string
	}{
		{
			name:               "Debit to exceed ANON limit",
			userLevel:          "ANON",
			initialAmount:      0,
			debitAmount:        50001,
			expectedMaxBalance: 0,
			expectedErr:        "unavailable operation",
		},
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
		tc := tc
		suite.T().Run(tc.name, func(t *testing.T) {
			ctxAuth := utils.ToCtx(context.Background(), token)

			if tc.userLevel == "FULL" {
				err := suite.UserSteps.UpgradeUser(ctxAuth, t, userID)
				require.NoError(t, err)
			}

			createdAccount, err := suite.UserSteps.CreateAccount(ctxAuth, t, userID, "description", tc.initialAmount)

			require.NoError(suite.T(), err)
			require.NotNil(t, createdAccount)
			require.Equal(t, tc.initialAmount, createdAccount.Amount)

			err = suite.UserSteps.Debit(ctxAuth, t, userID, createdAccount.AccountId, tc.debitAmount)

			require.Error(t, err)
			st, ok := status.FromError(err)
			require.True(t, ok, "expected gRPC status error")
			require.Equal(t, codes.InvalidArgument, st.Code())
			require.Contains(t, st.Message(), tc.expectedErr)

			gettingBalance, err := suite.UserSteps.GetAccountBalance(ctxAuth, t, userID, createdAccount.AccountId)

			require.NoError(suite.T(), err)
			require.NotNil(suite.T(), gettingBalance)
			require.Equal(suite.T(), tc.expectedMaxBalance, gettingBalance.Amount)
		})
	}
}

func (suite *WalletTestSuiteDebitCredit) TestCredit() {
	userID, token := utils.CreateUserHelper(suite.UserSteps, suite.T())

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
		tc := tc
		suite.T().Run(tc.name, func(t *testing.T) {
			ctxAuth := utils.ToCtx(context.Background(), token)

			if tc.userLevel == "FULL" {
				err := suite.UserSteps.UpgradeUser(ctxAuth, t, userID)
				require.NoError(t, err)
			}

			createdAccount, err := suite.UserSteps.CreateAccount(ctxAuth, t, userID, "description", tc.initialAmount)

			require.NoError(suite.T(), err)
			require.NotNil(t, createdAccount)
			require.Equal(t, tc.initialAmount, createdAccount.Amount)

			err = suite.UserSteps.Credit(ctxAuth, t, userID, createdAccount.AccountId, tc.creditAmount)

			require.NoError(t, err)

			gettingBalance, err := suite.UserSteps.GetAccountBalance(ctxAuth, t, userID, createdAccount.AccountId)

			require.NoError(suite.T(), err)
			require.NotNil(suite.T(), gettingBalance)
			require.Equal(suite.T(), tc.expectedBalance, gettingBalance.Amount)
		})
	}
}

func (suite *WalletTestSuiteDebitCredit) TestCreditNegative() {
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
		tc := tc
		suite.T().Run(tc.name, func(t *testing.T) {

			if tc.userLevel == "FULL" {
				err := suite.UserSteps.UpgradeUser(ctxAuth, t, suite.UserID)
				require.NoError(t, err)
			}

			createdAccount, err := suite.UserSteps.CreateAccount(ctxAuth, t, suite.UserID, "description", tc.initialAmount)

			require.NoError(suite.T(), err)
			require.NotNil(t, createdAccount)
			require.Equal(t, tc.initialAmount, createdAccount.Amount)

			err = suite.UserSteps.Credit(ctxAuth, t, suite.UserID, createdAccount.AccountId, tc.creditAmount)

			require.Error(t, err)
			st, ok := status.FromError(err)
			require.True(t, ok, "expected gRPC status error")
			require.Equal(t, codes.InvalidArgument, st.Code())
			require.Contains(t, st.Message(), tc.expectedErr)

			gettingBalance, err := suite.UserSteps.GetAccountBalance(ctxAuth, t, suite.UserID, createdAccount.AccountId)

			require.NoError(suite.T(), err)
			require.NotNil(suite.T(), gettingBalance)
			require.Equal(suite.T(), tc.expectedMaxBalance, gettingBalance.Amount)
		})
	}
}

// P.S. Уже не стал добавлять проверки на попытки списания/добавления на счет другого юзера, можно оставить на следующую домашку)
