//go:build grpctest
// +build grpctest

package grpc

import (
	"context"
	"strings"
	"testing"

	"github.com/ddosify/go-faker/faker"
	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/route256/wallet/internal/pkg/token"
	desc "gitlab.ozon.dev/route256/wallet/pkg/api/wallet"
)

func TestCreateUserAndLogIn(t *testing.T) {
	ctx := context.Background()
	fake := faker.NewFaker()

	randomPhone := strings.Replace("8"+fake.RandomPhoneNumber(), "-", "", -1)

	req := &desc.CreateUserRequest{
		Name:     fake.RandomPersonFirstName(),
		Lastname: fake.RandomPersonLastName(),
		Age:      int32(fake.RandomDigitNotNull()),
		Phone:    randomPhone,
		Password: fake.RandomPassword(),
	}

	resp, err := walletClient.CreateUser(ctx, req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, desc.IdentificationLevel_ANON, resp.Info.IdentificationLevel)

	loginResp, err := walletClient.LogIn(ctx, &desc.LogInRequest{
		Phone:    req.Phone,
		Password: req.Password,
	})
	require.NoError(t, err)
	require.NotNil(t, loginResp)
	require.NotEmpty(t, loginResp.Token)

	ctxWithToken := token.ToCtx(ctx, loginResp.Token)

	upgradeResp, err := walletClient.UpgradeUser(ctxWithToken, &desc.UpgradeUserRequest{
		UserId: resp.Info.Id,
	})
	require.NoError(t, err)
	require.NotNil(t, upgradeResp)

	getResp, err := walletClient.GetUser(ctxWithToken, &desc.GetUserRequest{
		UserId: resp.Info.Id,
	})
	require.NoError(t, err)
	require.NotNil(t, getResp)
	require.Equal(t, desc.IdentificationLevel_FULL, getResp.Info.IdentificationLevel)
}
