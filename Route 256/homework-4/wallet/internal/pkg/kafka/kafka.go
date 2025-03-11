package kafka

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
	"gitlab.ozon.dev/route256/wallet/internal/config"
	"golang.org/x/sync/errgroup"
)

type Client struct {
	writer *kafka.Writer
	reader *kafka.Reader

	consumeTick      time.Duration
	ConsumedMessages *MessageStore
}

type KafkaClient interface {
	ProduceMessages(ctx context.Context, messages ...kafka.Message) error
	ConsumeMessages(ctx context.Context)
	GetConsumedMessages() *MessageStore
}

func (c *Client) GetConsumedMessages() *MessageStore {
	return c.ConsumedMessages
}

func NewClient(config config.Kafka) *Client {
	return &Client{
		ConsumedMessages: NewMessageStore(1000),
		consumeTick:      config.ConsumeTicker,
		writer: &kafka.Writer{
			Addr:                   kafka.TCP(config.Hosts...),
			Topic:                  config.Topic,
			AllowAutoTopicCreation: true,
		},
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers:   config.Hosts,
			Topic:     config.Topic,
			Partition: 0,
			MaxBytes:  10e6, // 10MB
		}),
	}
}

func (c *Client) ProduceMessages(ctx context.Context, messages ...kafka.Message) error {
	return c.writer.WriteMessages(ctx, messages...)
}

func (c *Client) ConsumeMessages(ctx context.Context) {
	ticker := time.NewTicker(c.consumeTick)
	for {
		select {
		case <-ticker.C:
			msg, err := c.reader.ReadMessage(ctx)
			if err != nil {
				log.Error().Err(err).Msgf("error reading message")
				continue
			}

			c.ConsumedMessages.Store(msg)
			log.Info().Msgf("message at offset %d: %s = %s", msg.Offset, string(msg.Key), string(msg.Value))
		case <-ctx.Done():
			return
		}
	}
}

func (c *Client) Stop() error {
	eg := errgroup.Group{}

	eg.Go(c.writer.Close)
	eg.Go(c.reader.Close)

	return eg.Wait()
}
