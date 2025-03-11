package wallet

import (
	"context"

	"github.com/segmentio/kafka-go"
	desc "gitlab.ozon.dev/route256/wallet/pkg/api/wallet"
	"google.golang.org/protobuf/encoding/protojson"
)

func (i *Implementation) Consume(_ context.Context, req *desc.ConsumeRequest) (*desc.ConsumeResponse, error) {
	operations := make([]*desc.KafkaOperation, 0, req.GetCount())
	for _, consumedMsg := range i.kafka.GetConsumedMessages().Get(int(req.GetCount())) {
		operation, err := convertKafkaMessagesToProto(consumedMsg)
		if err != nil {
			return nil, err
		}

		operations = append(operations, operation)
	}

	return &desc.ConsumeResponse{
		Operations: operations,
	}, nil
}

func convertKafkaMessagesToProto(msg kafka.Message) (*desc.KafkaOperation, error) {
	operation := new(desc.KafkaOperation)
	err := protojson.Unmarshal(msg.Value, operation)
	if err != nil {
		return nil, err
	}

	return operation, nil
}
