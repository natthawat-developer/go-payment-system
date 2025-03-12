package handlers

import (
	"encoding/json"
	"log"

	"go-payment-system/internal/consumers/models"
	"go-payment-system/internal/consumers/services"
	"go-payment-system/pkg/kafka"
)

// TransactionHandler processes Kafka transaction events
type TransactionHandler struct {
	Service  *services.TransactionService
	Consumer *kafka.KafkaConsumer
	Topic    string
}

// NewTransactionHandler creates a new handler
func NewTransactionHandler(service *services.TransactionService, consumer *kafka.KafkaConsumer, topic string) *TransactionHandler {
	return &TransactionHandler{
		Service:  service,
		Consumer: consumer,
		Topic:    topic,
	}
}

// StartListening starts consuming transaction messages
func (h *TransactionHandler) StartListening() {
	h.Consumer.SubscribeTopics([]string{h.Topic}, nil)

	log.Println("Transaction Consumer is listening for messages...")

	for {
		msg, err := h.Consumer.ReadMessage(-1)
		if err == nil {
			var event models.TransactionEvent
			if err := json.Unmarshal(msg.Value, &event); err != nil {
				log.Printf("Failed to unmarshal transaction event: %v", err)
				continue
			}

			// Process the transaction event
			if err := h.ProcessTransaction(event); err != nil {
				log.Printf("Failed to process transaction: %v", err)
				continue
			}

		} else {
			log.Printf("Consumer error: %v", err)
		}
	}
}

// ProcessTransaction processes a received transaction event
func (h *TransactionHandler) ProcessTransaction(event models.TransactionEvent) error {
	log.Printf("Processing transaction: %s", event.TransactionID)
	return h.Service.ProcessTransaction(&event)
}
