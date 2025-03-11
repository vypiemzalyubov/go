package utils

import (
	"context"
	"e2e-backend/internal/steps/grpcsteps"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/stretchr/testify/require"
)

// Решил сделать этот хелпер для теста TestCreateAccountTooManyAccountsNegative, т.к при запуске всех тестов он падал, но работал при одиночном запуске.
// Предполагаю, что проблема в распарралеливании тестов и айдишники шарятся не так, как я хотел, поэтому тест при общемпрогоне падает. А это, по сути, костыльный хотфикс=)
func CreateUserHelper(steps *grpcsteps.UserSteps, t provider.T) (userID, token string) {
	ctx := context.Background()

	createdUser, err := steps.CreateUser(ctx, t, nil)
	require.NoError(t, err)
	require.NotNil(t, createdUser)

	loginToken, err := steps.LogIn(ctx, t, createdUser.Phone, createdUser.Password)
	require.NoError(t, err)
	require.NotNil(t, loginToken)

	return createdUser.ID, loginToken.Token
}
