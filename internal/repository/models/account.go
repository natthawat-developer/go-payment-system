package repository

import (
	"github.com/google/uuid"
	"time"
)

// Account represents a user's bank account
type Account struct {
	AccountID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID    uuid.UUID `gorm:"type:uuid;not null"`
	Balance   float64   `gorm:"type:decimal(18,2);not null;check:balance >= 0"`
	Currency  string    `gorm:"type:varchar(10);not null;default:'THB'"`
	CreatedAt time.Time `gorm:"autoCreateTime"`

	User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}
