package models

import "github.com/google/uuid"

// TransactionRequest represents the request payload for creating a transaction
type TransferRequest struct {
	FromAccountID uuid.UUID `json:"from_user_id" validate:"required,uuid"`
	ToAccountID   uuid.UUID `json:"to_user_id" validate:"required,uuid"`
	Amount        float64   `json:"amount" validate:"required,gt=0"`
	Currency      string    `json:"currency" validate:"required"`
}

type TransferResponse struct {
	TransferID   uuid.UUID `json:"transfer_id"`
	FromAccountID uuid.UUID `json:"from_account_id"`
	ToAccountID   uuid.UUID `json:"to_account_id"`
	Amount        float64   `json:"amount"`
	Currency      string    `json:"currency"`
	Status        string    `json:"status"`
}