package main

import (
	"go-payment-system/internal/adapters/kafka"
	"go-payment-system/internal/infrastructure/config"
	"go-payment-system/internal/adapters/database"
	"log"
)

func main() {
	// Load config
	cfg := config.LoadConfig()

	// Connect to PostgreSQL
	db := database.ConnectDB(cfg.DB)

	// Start Kafka consumer
	worker := kafka.NewConsumer(cfg.Kafka, db)
	log.Println("Worker Service is running...")
	worker.Start()
}
