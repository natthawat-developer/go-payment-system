package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

type DB struct {
	Conn *gorm.DB
}

// สร้าง DSN สำหรับเชื่อมต่อกับฐานข้อมูล
func BuildDSN(cfg *DatabaseConfig) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name, cfg.SSLMode)
}

// เชื่อมต่อกับ PostgreSQL
func ConnectDB(cfg *DatabaseConfig) (*gorm.DB, error) {
	// สร้าง DSN จากคอนฟิก
	dsn := BuildDSN(cfg)

	// เชื่อมต่อกับฐานข้อมูล PostgreSQL
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("❌ Database connection failed: %v", err)
	}

	log.Println("✅ Connected to PostgreSQL")
	return db, nil
}
