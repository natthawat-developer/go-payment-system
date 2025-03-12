package kafka

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// KafkaProducer holds the Kafka producer instance
type KafkaProducer struct {
	Producer *kafka.Producer
}

// ProducerConfig represents the Kafka producer configuration
type ProducerConfig struct {
	Brokers string
}

// NewProducer creates and returns a new KafkaProducer instance
func NewProducer(config *ProducerConfig) (*KafkaProducer, error) {
	kafkaConfig := &kafka.ConfigMap{
		"bootstrap.servers": config.Brokers,
	}

	producer, err := kafka.NewProducer(kafkaConfig)
	if err != nil {
		log.Printf("Failed to create Kafka producer: %v", err)
		return nil, err
	}

	return &KafkaProducer{Producer: producer}, nil
}

// Produce sends a message to a specified Kafka topic
func (kp *KafkaProducer) Produce(topic string, message []byte) error {
	err := kp.Producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          message,
	}, nil)

	if err != nil {
		log.Printf("Failed to produce message: %v", err)
		return err
	}

	return nil
}

// Close shuts down the Kafka producer
func (kp *KafkaProducer) Close() {
	if kp.Producer != nil {
		kp.Producer.Close()
	}
}
