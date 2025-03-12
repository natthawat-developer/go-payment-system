package consumers

import (
	"encoding/json"
	repository_models "go-payment-system/internal/repository/models"
	"go-payment-system/internal/services"
	"log"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// TransactionConsumer listens for transaction events from Kafka
type TransactionConsumer struct {
	Consumer *kafka.Consumer
	Service  *services.TransactionService
	Topic    string
}

// NewTransactionConsumer creates a new instance of TransactionConsumer
func NewTransactionConsumer(consumer *kafka.Consumer, service *services.TransactionService, topic string) *TransactionConsumer {
	return &TransactionConsumer{
		Consumer: consumer,
		Service:  service,
		Topic:    topic,
	}
}

// StartListening starts consuming transaction messages
func (tc *TransactionConsumer) StartListening() {
	tc.Consumer.SubscribeTopics([]string{tc.Topic}, nil)

	log.Println("Transaction Consumer is listening for messages...")

	for {
		msg, err := tc.Consumer.ReadMessage(-1)
		if err == nil {
			var transaction repository_models.Transaction
			if err := json.Unmarshal(msg.Value, &transaction); err != nil {
				log.Printf("Failed to unmarshal transaction: %v", err)
				continue
			}

			// Call TransactionService to process the transaction
			if err := tc.Service.ProcessTransaction(&transaction); err != nil {
				log.Printf("Failed to process transaction: %v", err)
				continue
			}

		} else {
			log.Printf("Consumer error: %v", err)
		}
	}
}
