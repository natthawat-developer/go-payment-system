package http

import (
	"go-payment-system/internal/adapters/database"
	"go-payment-system/internal/adapters/redis"
	"go-payment-system/internal/domain"
	"go-payment-system/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, db *database.DB, redisClient *redis.Client) {
	app.Post("/payment", func(c *fiber.Ctx) error {
		var tx domain.Transaction
		if err := c.BodyParser(&tx); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
		}

		// Process payment
		err := usecase.ProcessPayment(db, redisClient, tx)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"message": "Payment accepted"})
	})
}
