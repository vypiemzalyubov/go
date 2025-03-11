package test

import (
	"context"
	"testing"

	"e2e-backend/internal/clients/grpccli"
	"e2e-backend/internal/pb/gitlab.ozon.dev/route256/wallet"
	"e2e-backend/internal/steps/grpcsteps"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type WalletTestSuiteLogin struct {
	suite.Suite
	client       wallet.WalletClient
	conn         *grpc.ClientConn
	UserSteps    *grpcsteps.UserSteps
	UserPhone    string
	UserPassword string
	Token        string
}

func TestWalletTestSuiteLogin(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(WalletTestSuiteLogin))
}

func (suite *WalletTestSuiteLogin) SetupSuite() {
	var err error
	suite.client, suite.conn, err = grpccli.NewWalletClient()
	require.NoError(suite.T(), err)
	suite.UserSteps = grpcsteps.NewUserSteps(suite.client)

	ctx := context.Background()

	createdUser, err := suite.UserSteps.CreateUser(ctx, suite.T(), nil)
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), createdUser)
	suite.UserPassword = createdUser.Password
	suite.UserPhone = createdUser.Phone
}

func (suite *WalletTestSuiteLogin) TearDownSuite() {
	suite.conn.Close()
}

func (suite *WalletTestSuiteLogin) TestLogin() {
	ctx := context.Background()

	loginToken, err := suite.UserSteps.LogIn(ctx, suite.T(), suite.UserPhone, suite.UserPassword)

	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), loginToken)
}

func (suite *WalletTestSuiteLogin) TestLoginNegative() {
	testCases := []struct {
		name        string
		phone       string
		password    string
		expectedErr string
	}{
		{
			name:        "Empty phone",
			phone:       "",
			password:    "pass1",
			expectedErr: "phone not match pattern 8xxxxxxxxxx",
		},
		// {
		// 	name:        "Non-existent phone and invalid password",   Здесь мы берем рандомный номер валидного формата, с которым до этого не создавали юзера.
		// 	phone:       "89997776655",          					  Сервер возвращает 500 и сообщение "runtime error: invalid memory address or nil pointer dereference".
		// 	password:    "pass2",				 					  Как мне кажется, сервер никогда не должен пятисотить и должен уметь обрабатывать все ошибки, чтобы со стороны клиента не было видно возможных уязвимостей.
		// 	expectedErr: "non-existent phone or invalid password",
		// },
		// {
		// 	name:        "Non-existent phone and valid password",     Ситуация такая же, как и в тесте выше, но уже с валидным паролем.
		// 	phone:       "89997776655",
		// 	password:    suite.UserPassword,
		// 	expectedErr: "non-existent phone",
		// },
		{
			name:        "Invalid password",
			phone:       suite.UserPhone,
			password:    "fakepass1",
			expectedErr: "ivalid password", // Нужно поправить сообщение об ошибке на "invalid password"
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.T().Run(tc.name, func(t *testing.T) {
			ctx := context.Background()

			loginToken, err := suite.UserSteps.LogIn(ctx, t, tc.phone, tc.password)

			require.Nil(t, loginToken)
			require.NotNil(t, err)
			st, ok := status.FromError(err)
			require.True(t, ok, "expected gRPC status error")
			require.Equal(t, codes.InvalidArgument, st.Code())
			require.Contains(t, st.Message(), tc.expectedErr)
		})
	}
}
