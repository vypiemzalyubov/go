package wallet

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/route256/wallet/internal/pkg/mocks"
	desc "gitlab.ozon.dev/route256/wallet/pkg/api/wallet"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestProduce(t *testing.T) {
	t.Run("Produce success", func(t *testing.T) {
		mockedKafka := mocks.NewKafkaClientMock(t)
		impl := Implementation{kafka: mockedKafka}

		operation := &desc.KafkaOperation{
			OperationType: desc.OperationType_TYPE_DEBIT,
			Amount:        100,
			AccountId:     uuid.NewString(),
			ExternalId:    uuid.NewString(),
		}

		mockedKafka.ProduceMessagesMock.Return(nil)
		req := &desc.ProduceRequest{Operation: operation}

		resp, err := impl.Produce(context.Background(), req)
		require.NoError(t, err)
		require.NotNil(t, resp)
	})
	t.Run("Invalid operation", func(t *testing.T) {
		impl := Implementation{}

		req := &desc.ProduceRequest{Operation: nil}

		resp, err := impl.Produce(context.Background(), req)
		require.Error(t, err)
		require.Nil(t, resp)
		st, ok := status.FromError(err)
		require.True(t, ok)
		require.Equal(t, codes.InvalidArgument, st.Code())
		require.Contains(t, st.Message(), "operation must bu required")
	})
	t.Run("Invalid operation type", func(t *testing.T) {
		operation := &desc.KafkaOperation{
			OperationType: desc.OperationType_TYPE_UNDEFINED,
			Amount:        100,
			AccountId:     uuid.NewString(),
			ExternalId:    uuid.NewString(),
		}
		req := &desc.ProduceRequest{Operation: operation}

		impl := Implementation{}

		resp, err := impl.Produce(context.Background(), req)
		require.Error(t, err)
		require.Nil(t, resp)
		st, ok := status.FromError(err)
		require.True(t, ok)
		require.Equal(t, codes.InvalidArgument, st.Code())
		require.Contains(t, st.Message(), "operationType: TYPE_UNDEFINED should not be empty")
	})

	t.Run("Negative amount", func(t *testing.T) {
		operation := &desc.KafkaOperation{
			OperationType: desc.OperationType_TYPE_DEBIT,
			Amount:        -100,
			AccountId:     uuid.NewString(),
			ExternalId:    uuid.NewString(),
		}
		req := &desc.ProduceRequest{Operation: operation}

		impl := Implementation{}

		resp, err := impl.Produce(context.Background(), req)
		require.Error(t, err)
		require.Nil(t, resp)
		st, ok := status.FromError(err)
		require.True(t, ok)
		require.Equal(t, codes.InvalidArgument, st.Code())
		require.Contains(t, st.Message(), "amount: -100 should be more than 0")
	})

	t.Run("Invalid account ID", func(t *testing.T) {
		operation := &desc.KafkaOperation{
			OperationType: desc.OperationType_TYPE_DEBIT,
			Amount:        100,
			AccountId:     "invalid-uuid",
			ExternalId:    uuid.NewString(),
		}
		req := &desc.ProduceRequest{Operation: operation}

		impl := Implementation{}

		resp, err := impl.Produce(context.Background(), req)
		require.Error(t, err)
		require.Nil(t, resp)
		st, ok := status.FromError(err)
		require.True(t, ok)
		require.Equal(t, codes.InvalidArgument, st.Code())
		require.Contains(t, st.Message(), "accountID: invalid-uuid should have UUID format")
	})

	t.Run("Invalid external ID", func(t *testing.T) {
		operation := &desc.KafkaOperation{
			OperationType: desc.OperationType_TYPE_DEBIT,
			Amount:        100,
			AccountId:     uuid.NewString(),
			ExternalId:    "invalid-uuid",
		}
		req := &desc.ProduceRequest{Operation: operation}

		impl := Implementation{}

		resp, err := impl.Produce(context.Background(), req)
		require.Error(t, err)
		require.Nil(t, resp)
		st, ok := status.FromError(err)
		require.True(t, ok)
		require.Equal(t, codes.InvalidArgument, st.Code())
		require.Contains(t, st.Message(), "externalID: invalid-uuid should have UUID format")
	})
}
