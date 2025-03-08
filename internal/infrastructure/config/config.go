// internal/infrastructure/config/config.go
package config

import (
	"log"

	"github.com/spf13/viper"
)

// Config struct เก็บค่าคอนฟิกทั่วไปของระบบ
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Redis    RedisConfig
	Kafka    KafkaConfig
}

type ServerConfig struct {
	Port string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

type RedisConfig struct {
	Host string
	Port string
	Password string
	DB       int
}

type KafkaConfig struct {
	Broker string
	Topic  string
}

// LoadConfig โหลดค่าคอนฟิกจาก YAML + ENV Variables
func LoadConfig() (*Config, error) {
	viper.SetConfigName("config") // ชื่อไฟล์ (ไม่ต้องใส่นามสกุล)
	viper.SetConfigType("yaml")   // ประเภทไฟล์
	viper.AddConfigPath(".")      // ค้นหาในโฟลเดอร์ปัจจุบัน
	viper.AddConfigPath("/etc/payment-system/") // ค้นหาในโฟลเดอร์โปรดักชัน
	viper.AutomaticEnv() // รองรับ Environment Variables

	// อ่านไฟล์ config.yaml
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	cfg := &Config{}
	// โหลดค่าคอนฟิกทั้งหมด
	if err := viper.Unmarshal(cfg); err != nil {
		return nil, err
	}

	log.Println("✅ Config loaded successfully from", viper.ConfigFileUsed())
	return cfg, nil
}
