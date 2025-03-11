package kafka

import (
	"sync"

	"github.com/segmentio/kafka-go"
)

type MessageStore struct {
	mu               *sync.Mutex
	maxLimit         int
	consumedMessages []kafka.Message
}

func NewMessageStore(maxLimit int) *MessageStore {
	return &MessageStore{
		mu:               &sync.Mutex{},
		maxLimit:         maxLimit,
		consumedMessages: make([]kafka.Message, 0),
	}
}

func (m *MessageStore) Store(message kafka.Message) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.consumedMessages = append(m.consumedMessages, message)
}

func (m *MessageStore) Get(limit int) []kafka.Message {
	m.mu.Lock()
	defer m.mu.Unlock()

	if len(m.consumedMessages) > limit {
		return m.consumedMessages[:limit]
	}
	return m.consumedMessages
}
