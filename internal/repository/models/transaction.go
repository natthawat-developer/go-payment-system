package repository

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	TransactionID uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	FromAccountID uuid.UUID `gorm:"type:uuid;not null"`
	ToAccountID   uuid.UUID `gorm:"type:uuid;not null"`
	Amount        float64   `gorm:"not null;check:amount > 0"`
	Currency      string    `gorm:"type:varchar(10);not null"`
	Status        string    `gorm:"type:varchar(20);not null;default:'PENDING'"`
	CreatedAt     time.Time `gorm:"default:current_timestamp"`
	CompletedAt   *time.Time
}
