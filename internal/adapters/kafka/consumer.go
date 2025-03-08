package kafka

import (
	"encoding/json"
	"go-payment-system/internal/domain"
	"go-payment-system/internal/usecase"
	"log"

	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	Reader *kafka.Reader
}

func NewConsumer(brokers []string, topic string) *Consumer {
	return &Consumer{
		Reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers: brokers,
			Topic:   topic,
			GroupID: "payment-workers",
		}),
	}
}

func (c *Consumer) Start() {
	for {
		msg, err := c.Reader.ReadMessage(ctx)
		if err != nil {
			log.Println("Error reading message:", err)
			continue
		}

		var tx domain.Transaction
		json.Unmarshal(msg.Value, &tx)

		// Process payment transaction
		err = usecase.ProcessPayment(tx)
		if err != nil {
			log.Println("Transaction failed:", err)
		} else {
			log.Println("Transaction completed:", tx.ID)
		}
	}
}
