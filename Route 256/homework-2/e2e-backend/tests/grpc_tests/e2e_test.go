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
)

type WalletTestSuiteSmoke struct {
	suite.Suite
	client    wallet.WalletClient
	conn      *grpc.ClientConn
	UserSteps *grpcsteps.UserSteps
}

func TestWalletTestSuiteSmoke(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(WalletTestSuiteSmoke))
}

func (suite *WalletTestSuiteSmoke) SetupSuite() {
	var err error
	suite.client, suite.conn, err = grpccli.NewWalletClient()
	require.NoError(suite.T(), err)
	suite.UserSteps = grpcsteps.NewUserSteps(suite.client)
}

func (suite *WalletTestSuiteSmoke) TearDownSuite() {
	suite.conn.Close()
}

func (suite *WalletTestSuiteSmoke) TestSmoke() {
	ctx := context.Background()

	// 1. Создаем юзера
	createdUserAnon, err := suite.UserSteps.CreateUser(ctx, suite.T(), nil)
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), createdUserAnon)
	require.Equal(suite.T(), "ANON", createdUserAnon.Level)

	// 2. Логинимся
	loginToken, err := suite.UserSteps.LogIn(ctx, suite.T(), createdUserAnon.Phone, createdUserAnon.Password)
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), loginToken)

	// 3. Авторизуемся
	ctxAuth := utils.ToCtx(context.Background(), loginToken.Token)

	// 4. Получаем информацию о юзере
	gettingUserAnon, err := suite.UserSteps.GetUser(ctxAuth, suite.T(), createdUserAnon.ID)
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), gettingUserAnon)
	require.Equal(suite.T(), createdUserAnon.ID, gettingUserAnon.ID)

	// 5. Создаем счет юзера
	createdAccountAnon, err := suite.UserSteps.CreateAccount(ctxAuth, suite.T(), createdUserAnon.ID, " smoke description", int32(25000))
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), createdAccountAnon)
	require.Equal(suite.T(), int32(25000), createdAccountAnon.Amount)

	// 6. Получаем информацию о счете
	gettingBalanceAnon, err := suite.UserSteps.GetAccountBalance(ctxAuth, suite.T(), createdUserAnon.ID, createdAccountAnon.AccountId)
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), gettingBalanceAnon)
	require.Equal(suite.T(), int32(25000), gettingBalanceAnon.Amount)

	// 7. Кладем на счет сумму
	debitErrAnon := suite.UserSteps.Debit(ctxAuth, suite.T(), createdUserAnon.ID, createdAccountAnon.AccountId, int32(10000))
	require.NoError(suite.T(), debitErrAnon)

	// 8. Проверяем счет, после добавления суммы
	gettingBalanceAfterDebitAnon, err := suite.UserSteps.GetAccountBalance(ctxAuth, suite.T(), createdUserAnon.ID, createdAccountAnon.AccountId)
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), gettingBalanceAfterDebitAnon)
	require.Equal(suite.T(), int32(35000), gettingBalanceAfterDebitAnon.Amount)

	// 9. Списываем со счета сумму
	creditErrAnon := suite.UserSteps.Credit(ctxAuth, suite.T(), createdUserAnon.ID, createdAccountAnon.AccountId, int32(35000))
	require.NoError(suite.T(), creditErrAnon)

	// 10. Проверяем счет, после списывания суммы
	gettingBalanceAfterCreditAnon, err := suite.UserSteps.GetAccountBalance(ctxAuth, suite.T(), createdUserAnon.ID, createdAccountAnon.AccountId)
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), gettingBalanceAfterCreditAnon)
	require.Equal(suite.T(), int32(0), gettingBalanceAfterCreditAnon.Amount)

	// 11. Повышаем юзера до полного
	errUpgrade := suite.UserSteps.UpgradeUser(ctxAuth, suite.T(), createdUserAnon.ID)
	require.NoError(suite.T(), errUpgrade)

	// 12. Получаем информацию о полном юзере
	gettingUserFull, err := suite.UserSteps.GetUser(ctxAuth, suite.T(), createdUserAnon.ID)
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), gettingUserFull)
	require.Equal(suite.T(), createdUserAnon.ID, gettingUserFull.ID)
	require.Equal(suite.T(), "FULL", gettingUserFull.Level)

	// 13. Кладем на счет сумму
	debitErrFull := suite.UserSteps.Debit(ctxAuth, suite.T(), createdUserAnon.ID, createdAccountAnon.AccountId, int32(200000))
	require.NoError(suite.T(), debitErrFull)

	// 14. Проверяем счет, после добавления суммы
	gettingBalanceAfterDebitFull, err := suite.UserSteps.GetAccountBalance(ctxAuth, suite.T(), createdUserAnon.ID, createdAccountAnon.AccountId)
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), gettingBalanceAfterDebitFull)
	require.Equal(suite.T(), int32(200000), gettingBalanceAfterDebitFull.Amount)

	// 15. Списываем со счета сумму
	creditErrFull := suite.UserSteps.Credit(ctxAuth, suite.T(), createdUserAnon.ID, createdAccountAnon.AccountId, int32(199999))
	require.NoError(suite.T(), creditErrFull)

	// 16. Проверяем счет, после списывания суммы
	gettingBalanceAfterCreditFull, err := suite.UserSteps.GetAccountBalance(ctxAuth, suite.T(), createdUserAnon.ID, createdAccountAnon.AccountId)
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), gettingBalanceAfterCreditFull)
	require.Equal(suite.T(), int32(1), gettingBalanceAfterCreditFull.Amount)

}
