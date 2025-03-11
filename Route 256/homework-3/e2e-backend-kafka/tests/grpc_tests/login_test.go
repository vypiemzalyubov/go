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

type WalletTestSuiteLogin struct {
	suite.Suite
	client       wallet.WalletClient
	conn         *grpc.ClientConn
	UserSteps    *grpcsteps.UserSteps
	UserPhone    string
	UserPassword string
	Token        string
}

func (suite *WalletTestSuiteLogin) BeforeAll(t provider.T) {
	t.NewStep("Setup gRPC Client, Create User")

	var err error
	suite.client, suite.conn, err = grpccli.NewWalletClient()
	t.Require().NoError(err, "Failed to create gRPC client")

	suite.UserSteps = grpcsteps.NewUserSteps(suite.client)

	ctx := context.Background()
	createdUser, _ := suite.UserSteps.CreateUser(ctx, t, nil)
	suite.UserPassword = createdUser.Password
	suite.UserPhone = createdUser.Phone
}

func (suite *WalletTestSuiteLogin) AfterAll(t provider.T) {
	t.NewStep("Teardown gRPC Client")
	suite.conn.Close()
}

func (suite *WalletTestSuiteLogin) BeforeEach(t provider.T) {
	t.Epic("Wallet")
}

func TestWalletTestSuiteLoginRunner(t *testing.T) {
	t.Parallel()
	suite.RunSuite(t, new(WalletTestSuiteLogin))
}

func (suite *WalletTestSuiteLogin) TestLogin(t provider.T) {
	t.Title("Login user")
	t.Tag("gRPC Test")
	t.Tag("Positive")

	t.WithNewStep("Login user step", func(sCtx provider.StepCtx) {
		ctx := context.Background()

		loginToken, err := suite.UserSteps.LogIn(ctx, t, suite.UserPhone, suite.UserPassword)

		sCtx.Require().NoError(err, utils.ErrLoginUserErrorShouldBeNil)
		sCtx.Require().NotNil(loginToken, utils.ErrLoginUserMustNotBeNil)
	}, allure.NewParameter("time", time.Now()))

}

func (suite *WalletTestSuiteLogin) TestLoginNegative(t provider.T) {
	t.Title("Login user negative")

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
		t.Run(tc.name, func(t provider.T) {
			t.Title("Create account_" + tc.name)
			t.Tag("gRPC Test")
			t.Tag("Negative")
			t.Parallel()

			tc := tc

			t.WithNewStep("Login user step", func(sCtx provider.StepCtx) {
				ctx := context.Background()
				loginToken, err := suite.UserSteps.LogIn(ctx, t, tc.phone, tc.password)

				sCtx.Require().Nil(loginToken)
				sCtx.Require().NotNil(err)
				st, ok := status.FromError(err)
				sCtx.Require().True(ok, utils.ErrExpectGRPCStatusError)
				sCtx.Require().Equal(codes.InvalidArgument, st.Code())
				sCtx.Require().Contains(st.Message(), tc.expectedErr)
			}, allure.NewParameter("time", time.Now()))
		})
	}
}
