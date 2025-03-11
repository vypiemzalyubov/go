package grpc

// import (
// 	"context"
// 	"e2e-backend/internal/clients/grpccli"
// 	desc "e2e-backend/internal/pb/gitlab.ozon.dev/route256/wallet"
// 	"github.com/google/uuid"
// 	"github.com/segmentio/kafka-go"
// 	"github.com/stretchr/testify/require"
// 	"google.golang.org/protobuf/proto"
// 	"testing"
// 	"time"
// )

// func TestProduceConsume(t *testing.T) {
// 	ctx := context.Background()
// 	wallet := grpccli.NewWalletClient()

// 	operation := &desc.KafkaOperation{
// 		OperationType: desc.OperationType_TYPE_CREDIT,
// 		Amount:        1,
// 		AccountId:     uuid.NewString(),
// 		ExternalId:    uuid.NewString(),
// 	}

// 	resp, err := wallet.Produce(ctx, &desc.ProduceRequest{Operation: operation})
// 	require.NoError(t, err)
// 	require.NotNil(t, resp)

// 	require.Eventually(t, func() bool {
// 		resp, err := wallet.Consume(ctx, &desc.ConsumeRequest{Count: 100})
// 		require.NoError(t, err)
// 		require.NotNil(t, resp)

// 		t.Log(time.Now(), "ping")

// 		m := set(resp.GetOperations())
// 		_, ok := m[operation.GetExternalId()]

// 		return ok
// 	}, 20*time.Second, 500*time.Millisecond)
// }

// func set(slice []*desc.KafkaOperation) map[string]*desc.KafkaOperation {
// 	m := make(map[string]*desc.KafkaOperation)

// 	for _, el := range slice {
// 		m[el.GetExternalId()] = el
// 	}

// 	return m
// }

// func TestProduce(t *testing.T) {
// 	ctx := context.Background()
// 	wallet := grpccli.NewWalletClient()

// 	operation := &desc.KafkaOperation{
// 		OperationType: desc.OperationType_TYPE_CREDIT,
// 		Amount:        1,
// 		AccountId:     uuid.NewString(),
// 		ExternalId:    uuid.NewString(),
// 	}

// 	resp, err := wallet.Produce(ctx, &desc.ProduceRequest{Operation: operation})
// 	require.NoError(t, err)
// 	require.NotNil(t, resp)
// }

// func TestKafka(t *testing.T) {
// 	ctx := context.Background()

// 	operation := &desc.KafkaOperation{
// 		OperationType: desc.OperationType_TYPE_CREDIT,
// 		Amount:        1,
// 		AccountId:     uuid.NewString(),
// 		ExternalId:    uuid.NewString(),
// 	}

// 	b, err := proto.Marshal(operation)
// 	require.NoError(t, err)

// 	go kafkaReader(ctx, t)
// 	go kafkaWriter(ctx, t,
// 		kafka.Message{Value: b},
// 		//kafka.Message{Value: []byte("value")},
// 		//kafka.Message{Key: []byte("key"), Value: []byte("value")},
// 		//kafka.Message{Key: []byte("key"), Value: []byte("value")},
// 		//kafka.Message{
// 		//	Key:     []byte("key header"),
// 		//	Value:   []byte("value"),
// 		//	Headers: []kafka.Header{{Key: "header ker", Value: []byte("header value")}},
// 		//},
// 	)

// 	time.Sleep(20 * time.Second)
// }

// func kafkaWriter(ctx context.Context, t *testing.T, msg ...kafka.Message) {
// 	t.Helper()

// 	writer := &kafka.Writer{
// 		Addr:  kafka.TCP("127.0.0.1:29092"),
// 		Topic: "wallet_external_operations",
// 	}

// 	err := writer.WriteMessages(ctx, msg...)
// 	require.NoError(t, err)
// }

// func kafkaReader(ctx context.Context, t *testing.T) {
// 	t.Helper()

// 	reader := kafka.NewReader(kafka.ReaderConfig{
// 		Brokers:   []string{"127.0.0.1:29092"},
// 		Topic:     "wallet_external_operations",
// 		Partition: 0,
// 		GroupID:   "1",
// 		MaxBytes:  10e6, // 10MB
// 	})

// 	for {
// 		m, err := reader.ReadMessage(ctx)
// 		if err != nil {
// 			t.Log(err)
// 			continue
// 		}

// 		//err = reader.CommitMessages(ctx, m)
// 		//require.NoError(t, err)

// 		t.Logf("read msg: %#v", m)

// 		operation := new(desc.KafkaOperation)
// 		err = proto.Unmarshal(m.Value, operation)
// 		require.NoError(t, err)

// 		t.Logf("read msg: %#v", operation)
// 	}
// }
