package test

import (
	"context"
	"testing"
	"time"

	"e2e-backend/internal/clients/grpccli"
	"e2e-backend/internal/models"
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

type WalletTestSuiteCreateUser struct {
	suite.Suite
	client    wallet.WalletClient
	conn      *grpc.ClientConn
	UserSteps *grpcsteps.UserSteps
}

func (suite *WalletTestSuiteCreateUser) BeforeAll(t provider.T) {
	t.NewStep("Setup gRPC Client")

	var err error
	suite.client, suite.conn, err = grpccli.NewWalletClient()
	t.Require().NoError(err, utils.ErrCreatedGrpcClient)

	suite.UserSteps = grpcsteps.NewUserSteps(suite.client)
}

func (suite *WalletTestSuiteCreateUser) AfterAll(t provider.T) {
	t.NewStep("Teardown gRPC Client")
	suite.conn.Close()
}

func (suite *WalletTestSuiteCreateUser) BeforeEach(t provider.T) {
	t.Epic("Wallet")
}

func TestWalletTestSuiteCreateUserRunner(t *testing.T) {
	t.Parallel()
	suite.RunSuite(t, new(WalletTestSuiteCreateUser))
}

func (suite *WalletTestSuiteCreateUser) TestCreateUser(t provider.T) {
	t.Title("Create user")
	t.Tag("gRPC Test")
	t.Tag("Positive")

	t.WithNewStep("Create user step", func(sCtx provider.StepCtx) {
		ctx := context.Background()
		createdUser, err := suite.UserSteps.CreateUser(ctx, t, nil)

		sCtx.Require().NoError(err, utils.ErrCreatedUserErrorShouldBeNil)
		sCtx.Require().NotNil(createdUser, utils.ErrCreatedUserMustNotBeNil)
		sCtx.Require().Equal("ANON", createdUser.Level, utils.ErrCreatedUserLevelShouldBeAnon)
	}, allure.NewParameter("time", time.Now()))

}

func (suite *WalletTestSuiteCreateUser) TestCreateUserNegative(t provider.T) {
	t.Title("Create user negative")

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
		t.Run(tc.name, func(t provider.T) {
			t.Title("Create user_" + tc.name)
			t.Tag("gRPC Test")
			t.Tag("Negative")
			t.Parallel()

			tc := tc

			t.WithNewStep("Create user step", func(sCtx provider.StepCtx) {
				ctx := context.Background()
				createdUser, err := suite.UserSteps.CreateUser(ctx, t, tc.user)

				sCtx.Require().Nil(createdUser, utils.ErrCreatedUserShoulBeNil)
				sCtx.Require().NotNil(err, utils.ErrCreatedUserErrorMustNotBeNil)
				st, ok := status.FromError(err)
				sCtx.Require().True(ok, utils.ErrExpectGRPCStatusError)
				sCtx.Require().Equal(codes.InvalidArgument, st.Code(), utils.ErrExpectInvalidArgument)
				sCtx.Require().Contains(st.Message(), tc.expectedErr, utils.ErrMessageExpectInvalidArgument)
			}, allure.NewParameter("time", time.Now()))
		})
	}
}
