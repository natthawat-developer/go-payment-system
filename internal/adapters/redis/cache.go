package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"strconv"
)

var ctx = context.Background()

type CacheHandler struct {
	Conn *redis.Client
}

// NewCacheHandler สร้าง handler สำหรับการทำงานกับ Redis
func NewCacheHandler(client *redis.Client) *CacheHandler {
	return &CacheHandler{
		Conn: client,
	}
}

// GetBalance ดึงยอดเงินของผู้ใช้จาก Redis
func (r *CacheHandler) GetBalance(userID string) (float64, error) {
	balanceStr, err := r.Conn.Get(ctx, "balance:"+userID).Result()
	if err == redis.Nil {
		// หากไม่มีข้อมูลใน Redis จะถือว่าเป็น 0
		return 0, nil
	}
	if err != nil {
		return 0, err
	}

	// แปลงค่าจาก string เป็น float64
	balance, err := strconv.ParseFloat(balanceStr, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to convert balance to float64: %v", err)
	}
	return balance, nil
}

// DeductBalance หักยอดเงินจากบัญชีผู้ใช้
func (r *CacheHandler) DeductBalance(userID string, amount float64) error {
	// ดึงยอดเงินปัจจุบัน
	balance, err := r.GetBalance(userID)
	if err != nil {
		return err
	}

	// ตรวจสอบว่าเงินเพียงพอหรือไม่
	if balance < amount {
		return fmt.Errorf("insufficient balance")
	}

	// คำนวณยอดเงินใหม่
	newBalance := balance - amount

	// ตั้งค่าความสมดุลใหม่ใน Redis
	err = r.Conn.Set(ctx, "balance:"+userID, newBalance, 0).Err()
	if err != nil {
		return fmt.Errorf("failed to update balance: %v", err)
	}

	log.Printf("Deducted %.2f from user %s. New balance: %.2f", amount, userID, newBalance)
	return nil
}
