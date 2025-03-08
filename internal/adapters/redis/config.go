package redis

import (
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
}

// NewRedisClient สร้าง Redis client จากการตั้งค่า config
func NewRedisClient(cfg *RedisConfig) *redis.Client {
	addr := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	// ตรวจสอบการเชื่อมต่อกับ Redis
	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("❌ Error connecting to Redis: %v", err)
	}
	log.Println("✅ Connected to Redis")
	return client
}
