package services

import (
	"encoding/json"
	"errors"

	handlers_models "go-payment-system/internal/handlers/models"
	repository_models "go-payment-system/internal/repository/models"

	"go-payment-system/internal/repository"
	"go-payment-system/pkg/kafka"
	"go-payment-system/pkg/logger"

)

// TransactionService handles transaction-related operations
type TransactionService struct {
	repo     repository.TransactionRepository
	producer *kafka.KafkaProducer
}

// NewTransactionService creates a new instance of TransactionService
func NewTransactionService(repo repository.TransactionRepository, producer *kafka.KafkaProducer) *TransactionService {
	return &TransactionService{
		repo:     repo,
		producer: producer,
	}
}

// CreateTransaction handles transaction creation logic
func (s *TransactionService) CreateTransfer(request *handlers_models.TransferRequest) (*handlers_models.TransferResponse, error) {

	transaction := repository_models.Transaction{
		FromAccountID: request.FromAccountID,
		ToAccountID:   request.ToAccountID,
		Amount:        request.Amount,
		Currency:      request.Currency,
		Status:        "PENDING",
	}

	// Save transaction to database
	if err := s.repo.CreateTransaction(&transaction); err != nil {
		return nil, errors.New("transaction failed to save")
	}

	// Send Kafka message
	message, err := json.Marshal(transaction)
	if err != nil {
		return nil, errors.New("failed to serialize transaction data")
	}

	err = s.producer.Produce("transaction-events", message)
	if err != nil {
		return nil, errors.New("failed to send transaction to Kafka")
	}

	logger.Log.Info("Transaction sent to Kafka: ", transaction.TransactionID)
	
	response := &handlers_models.TransferResponse{
		TransferID:   transaction.TransactionID,
		FromAccountID: transaction.FromAccountID,
		ToAccountID:   transaction.ToAccountID,
		Amount:        transaction.Amount,
		Currency:      transaction.Currency,
		Status:        transaction.Status,
	}

	return response, nil
}

// ProcessTransaction processes a transaction event received from Kafka
func (s *TransactionService) ProcessTransaction(transaction *repository_models.Transaction) error {
	// Update transaction status in database
	if err := s.repo.UpdateTransaction(transaction); err != nil {
		logger.Log.Errorf("Failed to update transaction: %v", err)
		return err
	}

	logger.Log.Infof("Processed transaction: %s", transaction.TransactionID)
	return nil
}