package repository

import (
	"time"
	"github.com/google/uuid"
)


// Withdrawal represents a withdrawal transaction
type Withdrawal struct {
	WithdrawalID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	AccountID    uuid.UUID `gorm:"type:uuid;not null"`
	Amount       float64   `gorm:"type:decimal(18,2);not null;check:amount > 0"`
	Currency     string    `gorm:"type:varchar(10);not null"`
	Status       string    `gorm:"type:varchar(20);not null;check:status IN ('PENDING', 'COMPLETED', 'FAILED')"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	CompletedAt  *time.Time `gorm:"default:null"`

	Account Account `gorm:"foreignKey:AccountID;constraint:OnDelete:CASCADE"`
}