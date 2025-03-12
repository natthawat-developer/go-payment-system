package main

import (
	"log"

	"go-payment-system/config"
	"go-payment-system/internal/consumers/handlers"
	"go-payment-system/internal/consumers/services"
	"go-payment-system/internal/repository"
	"go-payment-system/pkg/database"
	"go-payment-system/pkg/kafka"
)

func main() {
	// โหลดค่า Config
	if err := config.LoadConfig(); err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// เชื่อมต่อ Database
	if err := database.ConnectDB(&database.DatabaseConfig{
		Host:     config.Config.Database.Host,
		Port:     config.Config.Database.Port,
		User:     config.Config.Database.User,
		Password: config.Config.Database.Password,
		DBName:   config.Config.Database.DBName,
		SSLMode:  config.Config.Database.SSLMode,
	}); err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	// สร้าง Kafka Consumer
	consumer, err := kafka.NewConsumer(&kafka.ConsumerConfig{
		Brokers: config.Config.Kafka.Brokers[0],
		GroupID: config.Config.Kafka.GroupID,
	})
	if err != nil {
		log.Fatalf("Failed to create Kafka consumer: %v", err)
	}
	defer consumer.Close()

	// สร้าง Repository และ Service
	transactionRepo := repository.NewTransactionRepository(database.DB)
	transactionService := services.NewTransactionService(transactionRepo)

	// สร้าง Handler สำหรับ Kafka Consumer
	transactionHandler := handlers.NewTransactionHandler(transactionService, consumer, config.Config.Kafka.Topic)

	// เริ่มฟัง Kafka
	transactionHandler.StartListening()
}
