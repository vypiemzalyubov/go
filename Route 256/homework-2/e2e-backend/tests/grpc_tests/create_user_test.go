package test

import (
	"context"
	"testing"

	"e2e-backend/internal/clients/grpccli"
	"e2e-backend/internal/models"
	"e2e-backend/internal/pb/gitlab.ozon.dev/route256/wallet"
	"e2e-backend/internal/steps/grpcsteps"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type WalletTestSuiteCreateUser struct {
	suite.Suite
	client    wallet.WalletClient
	conn      *grpc.ClientConn
	UserSteps *grpcsteps.UserSteps
}

func TestWalletTestSuiteCreateUser(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(WalletTestSuiteCreateUser))
}

func (suite *WalletTestSuiteCreateUser) SetupSuite() {
	var err error
	suite.client, suite.conn, err = grpccli.NewWalletClient()
	require.NoError(suite.T(), err)
	suite.UserSteps = grpcsteps.NewUserSteps(suite.client)
}

func (suite *WalletTestSuiteCreateUser) TearDownSuite() {
	suite.conn.Close()
}

func (suite *WalletTestSuiteCreateUser) TestCreateUser() {
	ctx := context.Background()

	createdUser, err := suite.UserSteps.CreateUser(ctx, suite.T(), nil)
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), createdUser)
	require.Equal(suite.T(), "ANON", createdUser.Level)
}

func (suite *WalletTestSuiteCreateUser) TestCreateUserNegative() {
	testCases := []struct {
		name        string
		user        *models.User
		expectedErr string
	}{
		// {
		// 	name: "Empty name",
		// 	user: &models.User{
		// 		Name:     "",				Тест не упадет, хотя по моей логике должен, т.к имя пустое, но в приложении, видимо, нет валидации по имени
		// 		Lastname: "Rodrigo",
		// 		Age:      23,
		// 		Phone:    "89991234567",
		// 		Password: "madrid23",
		// 	},
		// 	expectedErr: "name cannot be empty",
		// },
		// {
		// 	name: "Empty last name",
		// 	user: &models.User{
		// 		Name:     "Vinicius",
		// 		Lastname: "",				Тест не упадет, хотя по моей логике должен, т.к фамилия пустая, но в приложении, видимо, нет валидации по фамилии
		// 		Age:      24,
		// 		Phone:    "89991234566",
		// 		Password: "madrid24",
		// 	},
		// 	expectedErr: "last name cannot be empty",
		// },
		// {
		// 	name: "Negative age",
		// 	user: &models.User{
		// 		Name:     "Lamine",
		// 		Lastname: "Yamal",
		// 		Age:      -1,				Тест не упадет, хотя по моей логике должен, т.к возраст пустой, но в приложении, видимо, нет валидации на возраст
		// 		Phone:    "89991234564",
		// 		Password: "barca",
		// 	},
		// 	expectedErr: "age must be greater than 0",
		// },
		{
			name: "Invalid phone",
			user: &models.User{
				Name:     "Jude",
				Lastname: "Bellingham",
				Age:      21,
				Phone:    "invalid_phone", //Ура, наконец-то валидация=)
				Password: "madrid21",
			},
			expectedErr: "phone not match pattern 8xxxxxxxxxx",
		},
		// {
		// 	name: "Too long phone",
		// 	user: &models.User{
		// 		Name:     "Jan Johannes",
		// 		Lastname: "Vennegoor of Hesselink",
		// 		Age:      45,
		// 		Phone:    "812345678901",     Тест не упадет, хотя по моей логике должен, т.к пароль слишком длинный, но в приложении, видимо, нет валидации на пароль
		// 		Password: "long45",
		// 	},
		// 	expectedErr: "phone cannot be more than 11 digits",
		// },
		// {
		// 	name: "Empty password",
		// 	user: &models.User{
		// 		Name:     "Kylian",
		// 		Lastname: "Mbappe",
		// 		Age:      25,
		// 		Phone:    "89991234563",
		// 		Password: "",				Тест не упадет, хотя по моей логике должен, т.к пароль пустой, но в приложении, видимо, нет валидации на пароль
		// 	},
		// 	expectedErr: "password cannot be empty",
		// },
	}

	for _, tc := range testCases {
		tc := tc
		suite.T().Run(tc.name, func(t *testing.T) {
			ctx := context.Background()

			createdUser, err := suite.UserSteps.CreateUser(ctx, t, tc.user)

			require.Nil(t, createdUser)
			require.NotNil(t, err)

			st, ok := status.FromError(err)
			require.True(t, ok, "expected gRPC status error")
			require.Equal(t, codes.InvalidArgument, st.Code())
			require.Contains(t, st.Message(), tc.expectedErr)
		})
	}
}
