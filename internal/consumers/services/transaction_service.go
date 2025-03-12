package services

import (
	"go-payment-system/internal/consumers/models"
	"go-payment-system/internal/repository"
	repository_models "go-payment-system/internal/repository/models"
	"go-payment-system/pkg/logger"
)

// TransactionService handles transaction-related operations
type TransactionService struct {
	repo repository.TransactionRepository
}

// NewTransactionService creates a new instance of TransactionService
func NewTransactionService(repo repository.TransactionRepository) *TransactionService {
	return &TransactionService{
		repo: repo,
	}
}

// ProcessTransaction processes a transaction event received from Kafka
func (s *TransactionService) ProcessTransaction(transaction *models.TransactionEvent) error {
	// Update transaction status in database
	if err := s.repo.UpdateTransaction(&repository_models.Transaction{
		FromAccountID: transaction.FromAccountID,
		ToAccountID:   transaction.ToAccountID,
		Amount:        transaction.Amount,
		Status:        transaction.Status,
	}); err != nil {
		logger.Log.Errorf("Failed to update transaction: %v", err)
		return err
	}

	logger.Log.Infof("Processed transaction: %s", transaction.TransactionID)
	return nil
}
