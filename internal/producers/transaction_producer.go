package producers

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// TransactionProducer handles producing transaction messages
type TransactionProducer struct {
	Producer *kafka.Producer
	Topic    string
}

// NewTransactionProducer creates a new TransactionProducer
func NewTransactionProducer(producer *kafka.Producer, topic string) *TransactionProducer {
	return &TransactionProducer{
		Producer: producer,
		Topic:    topic,
	}
}

// PublishTransactionMessage sends a transaction message to Kafka
func (tp *TransactionProducer) PublishTransactionMessage(message []byte) error {
	err := tp.Producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &tp.Topic, Partition: kafka.PartitionAny},
		Value:          message,
	}, nil)

	if err != nil {
		return err
	}

	fmt.Println("Transaction sent to Kafka")
	return nil
}
