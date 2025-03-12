package config

import (
	"log"

	"github.com/spf13/viper"
)

// Configurations holds all the configurations from config.yaml
type Configurations struct {
	Server   ServerConfig
	Database DatabaseConfig
	Kafka    KafkaConfig
	Logger   LoggerConfig
}

// ServerConfig holds server configuration
type ServerConfig struct {
	Port string
}

// DatabaseConfig holds database configuration
type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// KafkaConfig holds Kafka-related configuration
type KafkaConfig struct {
	Brokers []string
	Topic   string
	GroupID string
}

// LoggerConfig holds logging configuration
type LoggerConfig struct {
	Level string
	File  string
}

var Config *Configurations

// LoadConfig loads configuration from config.yaml
func LoadConfig() error {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("./config")

	if err := v.ReadInConfig(); err != nil {
		log.Printf("Error reading config file: %v", err)
		return err
	}

	Config = &Configurations{}
	if err := v.Unmarshal(Config); err != nil {
		log.Printf("Unable to decode config into struct: %v", err)
		return err
	}

	return nil
}
