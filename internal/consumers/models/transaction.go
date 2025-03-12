package models

import (
	"github.com/google/uuid"
)

// TransactionEvent represents the structure of a transaction event from Kafka
type TransactionEvent struct {
	TransactionID uuid.UUID `json:"transaction_id"`
	FromAccountID uuid.UUID `json:"from_account_id"`
	ToAccountID   uuid.UUID `json:"to_account_id"`
	Amount        float64   `json:"amount"`
	Currency      string    `json:"currency"`
	Status        string    `json:"status"`
}
