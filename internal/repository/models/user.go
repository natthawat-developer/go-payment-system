package repository

import (
	"time"
	"github.com/google/uuid"
)

// User represents a user in the system
type User struct {
	UserID      uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	FullName    string    `gorm:"type:varchar(255);not null"`
	Email       string    `gorm:"type:varchar(255);unique;not null"`
	PhoneNumber string    `gorm:"type:varchar(20);unique;not null"`
	PasswordHash string   `gorm:"type:varchar(255);not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
}



