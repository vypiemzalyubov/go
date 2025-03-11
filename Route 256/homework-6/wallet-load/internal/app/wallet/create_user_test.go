package wallet

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/ddosify/go-faker/faker"
	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/route256/wallet/internal/pkg/mocks"
	desc "gitlab.ozon.dev/route256/wallet/pkg/api/wallet"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestCreateUser_Positive(t *testing.T) {
	t.Run("Simple create user", func(t *testing.T) {
		fake := faker.NewFaker()
		impl := Implementation{store: store}

		randomPhone := strings.Replace("8"+fake.RandomPhoneNumber(), "-", "", -1)

		req := &desc.CreateUserRequest{
			Name:     "Fake",
			Lastname: "User",
			Age:      20,
			Phone:    randomPhone,
			Password: "strong123",
		}

		resp, err := impl.CreateUser(context.Background(), req)
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, req.Name, resp.Info.Name)
		require.Equal(t, req.Lastname, resp.Info.Lastname)
		require.Equal(t, req.Phone, resp.Info.Phone)
	})
}

func TestCreateUser_Negative(t *testing.T) {
	t.Run("Create user with invalid phone format", func(t *testing.T) {
		impl := Implementation{store: store}

		req := &desc.CreateUserRequest{
			Name:     "Fake",
			Lastname: "Fakovich",
			Age:      30,
			Phone:    "7001234567",
			Password: "strong456",
		}

		_, err := impl.CreateUser(context.Background(), req)
		require.Error(t, err)
		require.Equal(t, codes.InvalidArgument, status.Code(err))
		require.Contains(t, err.Error(), "phone not match pattern 8xxxxxxxxxx")
	})
	t.Run("Internal error when creating user", func(t *testing.T) {
		mockedStore := mocks.NewMinimockStorage(t)
		impl := Implementation{store: mockedStore}

		mockedStore.CreateUserMock.Return(fmt.Errorf("CreateUser err"))

		req := &desc.CreateUserRequest{
			Name:     "Mark",
			Lastname: "Dacascos",
			Age:      60,
			Phone:    "89992221122",
			Password: "verystrong",
		}

		_, err := impl.CreateUser(context.Background(), req)
		require.Error(t, err)
		require.Equal(t, codes.Internal, status.Code(err))
		require.Contains(t, err.Error(), "CreateUser err")
	})
}
