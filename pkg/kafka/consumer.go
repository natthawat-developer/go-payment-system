package kafka

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// KafkaConsumer holds the Kafka consumer instance
type KafkaConsumer struct {
	Consumer *kafka.Consumer
}

// ConsumerConfig represents the Kafka consumer configuration
type ConsumerConfig struct {
	Brokers string
	GroupID string
}

// NewConsumer creates and returns a new KafkaConsumer instance
func NewConsumer(config *ConsumerConfig) (*KafkaConsumer, error) {
	kafkaConfig := &kafka.ConfigMap{
		"bootstrap.servers": config.Brokers,
		"group.id":          config.GroupID,
		"auto.offset.reset": "earliest",
	}

	consumer, err := kafka.NewConsumer(kafkaConfig)
	if err != nil {
		log.Printf("Failed to create Kafka consumer: %v", err)
		return nil, err
	}

	return &KafkaConsumer{Consumer: consumer}, nil
}

// Close shuts down the Kafka consumer
func (kc *KafkaConsumer) Close() {
	if kc.Consumer != nil {
		kc.Consumer.Close()
	}
}
