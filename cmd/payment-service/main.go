package main

import (
	"github.com/gofiber/fiber/v2"
	"go-payment-system/internal/adapters/database"
	"go-payment-system/internal/adapters/http"
	"go-payment-system/internal/adapters/redis"
	"go-payment-system/internal/infrastructure/config"
	"go-payment-system/internal/infrastructure/metrics"
	"log"
)

func main() {
	// Load config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("❌ Error loading config: %v", err)
	}

	db, err := database.ConnectDB(&database.DatabaseConfig{
		Host:     cfg.Database.Host,
		Port:     cfg.Database.Port,
		User:     cfg.Database.User,
		Password: cfg.Database.Password,
		Name:     cfg.Database.Name,
		SSLMode:  cfg.Database.SSLMode})
	if err != nil {
		log.Fatalf("❌ Error connecting to database: %v", err)
	}

	// Connect to Redis
	redisClient := redis.NewRedisClient(&redis.RedisConfig{
		Host:     cfg.Redis.Host,
		Port:     cfg.Redis.Port,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	// Initialize Fiber app
	app := fiber.New()

	// Register Metrics
	metrics.RegisterMetrics()

	// Start Prometheus metrics server
	go metrics.StartMetricsServer(cfg.Metrics.Port)

	// Setup Routes
	http.SetupRoutes(app, db, redisClient)

	// Start server
	log.Println("🚀 Payment Service running on port", cfg.Server.Port)
	log.Fatal(app.Listen(":" + cfg.Server.Port))
}
