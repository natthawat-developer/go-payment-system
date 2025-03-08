package domain

import (
	"time"
)

type Transaction struct {
	ID         string    `json:"id" gorm:"primaryKey"`
	SenderID   string    `json:"sender_id" gorm:"index"`
	ReceiverID string    `json:"receiver_id" gorm:"index"`
	Amount     float64   `json:"amount"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
}
