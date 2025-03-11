package wallet

import (
	"context"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	desc "gitlab.ozon.dev/route256/wallet/pkg/api/wallet"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
)

func (i *Implementation) Produce(ctx context.Context, req *desc.ProduceRequest) (*desc.ProduceResponse, error) {
	err := validateProduceRequest(req)
	if err != nil {
		return nil, err
	}

	reqB, err := protojson.Marshal(req.GetOperation())
	if err != nil {
		return nil, err
	}

	err = i.kafka.ProduceMessages(ctx, kafka.Message{Value: reqB})
	if err != nil {
		return nil, err
	}

	return &desc.ProduceResponse{}, nil
}

func validateProduceRequest(req *desc.ProduceRequest) error {
	if req.Operation == nil {
		return status.Errorf(codes.InvalidArgument, "operation must bu required")
	}

	operation := req.GetOperation()

	if operation.GetOperationType() == desc.OperationType_TYPE_UNDEFINED {
		return status.Errorf(codes.InvalidArgument, "operationType: %s should not be empty", operation.GetOperationType())
	}

	if operation.GetAmount() < 0 {
		return status.Errorf(codes.InvalidArgument, "amount: %d should be more than 0", operation.GetAmount())
	}

	_, err := uuid.Parse(operation.GetAccountId())
	if err != nil {
		return status.Errorf(codes.InvalidArgument, "accountID: %s should have UUID format", operation.GetAccountId())
	}

	_, err = uuid.Parse(operation.GetExternalId())
	if err != nil {
		return status.Errorf(codes.InvalidArgument, "externalID: %s should have UUID format", operation.GetExternalId())
	}

	return nil
}
