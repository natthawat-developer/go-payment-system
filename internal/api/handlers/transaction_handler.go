package handlers

import (
	"go-payment-system/internal/api/services"
	 "go-payment-system/internal/api/models"
	"github.com/gofiber/fiber/v2"
)

// TransactionHandler handles transaction routes
type TransactionHandler struct {
	service *services.TransactionService
}

// NewTransactionHandler creates a new instance of TransactionHandler
func NewTransactionHandler(service *services.TransactionService) *TransactionHandler {
	return &TransactionHandler{service: service}
}

func (h *TransactionHandler) CreateTransfer(c *fiber.Ctx) error {
	var request models.TransferRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	response, err := h.service.CreateTransfer(&request)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(response)
}
