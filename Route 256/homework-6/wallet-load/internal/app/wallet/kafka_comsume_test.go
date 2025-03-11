package wallet

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	messageStore "gitlab.ozon.dev/route256/wallet/internal/pkg/kafka"
	"gitlab.ozon.dev/route256/wallet/internal/pkg/mocks"
	desc "gitlab.ozon.dev/route256/wallet/pkg/api/wallet"
	"google.golang.org/protobuf/encoding/protojson"
)

func TestConsume(t *testing.T) {
	t.Run("Ok", func(t *testing.T) {
		msgStore := messageStore.NewMessageStore(100)

		operation := &desc.KafkaOperation{
			OperationType: desc.OperationType_TYPE_CREDIT,
			AccountId:     uuid.NewString(),
			ExternalId:    uuid.NewString(),
		}

		b, err := protojson.Marshal(operation)
		require.NoError(t, err)

		msgStore.Store(kafka.Message{Value: b})

		msgs := msgStore.Get(100)
		require.Len(t, msgs, 1)

		actual, err := convertKafkaMessagesToProto(msgs[0])
		require.NoError(t, err)
		require.Equal(t, operation, actual)
	})
	t.Run("Consume success", func(t *testing.T) {
		mockedKafka := mocks.NewKafkaClientMock(t)
		mockedStore := mocks.NewMinimockStorage(t)
		impl := Implementation{store: mockedStore, kafka: mockedKafka}

		testMessage := kafka.Message{
			Value: []byte(`{"operationType": 1, "amount": 100, "accountId": "fakeid123", "externalId": "fake123"}`),
		}

		msgStore := messageStore.NewMessageStore(100)
		msgStore.Store(testMessage)

		mockedKafka.GetConsumedMessagesMock.Return(msgStore)

		req := &desc.ConsumeRequest{Count: 1}

		resp, err := impl.Consume(context.Background(), req)
		require.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Len(t, resp.Operations, 1)
		assert.Equal(t, "fake123", resp.Operations[0].ExternalId)
	})
	t.Run("Empty message", func(t *testing.T) {
		msgStore := messageStore.NewMessageStore(100)

		msgStore.Store(kafka.Message{Value: nil})

		msgs := msgStore.Get(100)
		require.Len(t, msgs, 1)

		_, err := convertKafkaMessagesToProto(msgs[0])
		require.Error(t, err)
	})
}
