package usecase

import (
	"errors"
	"go-payment-system/internal/domain"
	"go-payment-system/internal/infrastructure/database"
	"go-payment-system/internal/infrastructure/redis"
)

func ProcessPayment(db *database.DB, redisClient *redis.Client, tx domain.Transaction) error {
	// ตรวจสอบยอดเงินใน Redis
	balance, err := redisClient.GetBalance(tx.SenderID)
	if err != nil || balance < tx.Amount {
		return errors.New("insufficient balance")
	}

	// หักเงินออกจากบัญชี (Atomic Transaction)
	err = redisClient.DeductBalance(tx.SenderID, tx.Amount)
	if err != nil {
		return err
	}

	// บันทึกธุรกรรมลง PostgreSQL
	err = db.SaveTransaction(tx)
	if err != nil {
		return err
	}

	return nil
}
