package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"go-payment-system/config"
	"go-payment-system/internal/api/handlers"
	"go-payment-system/internal/repository"
	"go-payment-system/internal/api/routes"
	"go-payment-system/internal/api/services"
	"go-payment-system/pkg/database"
	"go-payment-system/pkg/kafka"
)

func main() {
	// Load configuration
	if err := config.LoadConfig(); err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Connect to database
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

	producerConfig := &kafka.ProducerConfig{
		Brokers: config.Config.Kafka.Brokers[0],
	}

	producer, err := kafka.NewProducer(producerConfig)
	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %v", err)
	}
	defer producer.Close()

	// Initialize Fiber app
	app := fiber.New()

	// Initialize repositories
	transactionRepo := repository.NewTransactionRepository(database.DB)

	// Initialize services
	transactionService := services.NewTransactionService(transactionRepo, producer)

	// Initialize handlers with dependency injection
	transactionHandler := handlers.NewTransactionHandler(transactionService)

	// Register routes with injected handlers
	routes.SetupRoutes(app, transactionHandler)

	// Start server
	port := ":" + config.Config.Server.Port
	log.Printf("Server is running on port %s", port)
	log.Fatal(app.Listen(":" + port))
}
