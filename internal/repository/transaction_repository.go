package repository

import (
	"errors"
	repository_models "go-payment-system/internal/repository/models"

	"gorm.io/gorm"
)

// TransactionRepository defines the interface for transaction database operations
type TransactionRepository interface {
	CreateTransaction(transaction *repository_models.Transaction) error
	GetTransactionByID(transactionID string) (*repository_models.Transaction, error)
	UpdateTransaction(transaction *repository_models.Transaction) error
	DeleteTransaction(transactionID string) error
}

// transactionRepository is the concrete implementation of TransactionRepository
type transactionRepository struct {
	db *gorm.DB
}

// NewTransactionRepository creates a new instance of TransactionRepository
func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db: db}
}

// CreateTransaction inserts a new transaction into the database
func (r *transactionRepository) CreateTransaction(transaction *repository_models.Transaction) error {
	return r.db.Create(transaction).Error
}

// GetTransactionByID retrieves a transaction by its ID
func (r *transactionRepository) GetTransactionByID(transactionID string) (*repository_models.Transaction, error) {
	var transaction repository_models.Transaction
	err := r.db.Where("transaction_id = ?", transactionID).First(&transaction).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil // Return nil instead of an error for "not found" case
	}
	return &transaction, err
}

// UpdateTransaction updates an existing transaction
func (r *transactionRepository) UpdateTransaction(transaction *repository_models.Transaction) error {
	return r.db.Model(&repository_models.Transaction{}).
		Where("transaction_id = ?", transaction.TransactionID).
		Updates(transaction).Error
}

// DeleteTransaction removes a transaction from the database
func (r *transactionRepository) DeleteTransaction(transactionID string) error {
	return r.db.Where("transaction_id = ?", transactionID).Delete(&repository_models.Transaction{}).Error
}
