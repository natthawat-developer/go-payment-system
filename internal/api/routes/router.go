package routes

import (
	"go-payment-system/internal/api/handlers"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes registers routes with dependency injection
func SetupRoutes(app *fiber.App, transactionHandler *handlers.TransactionHandler) {
	api := app.Group("/api")

	// Transaction Routes
	api.Post("/transfers", transactionHandler.CreateTransfer)
	//api.Post("/withdrawals", handlers.CreateWithdrawal)
}
