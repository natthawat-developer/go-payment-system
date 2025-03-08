package middleware

import (
	"go-payment-system/internal/infrastructure/metrics"
	"go-payment-system/internal/infrastructure/prometheus"
	"time"

	"github.com/gofiber/fiber/v2"
)

// MetricsMiddleware ใช้ในการติดตามข้อมูล API request และ response
func MetricsMiddleware(c *fiber.Ctx) error {
	start := time.Now()

	// บันทึกข้อมูล request (method และ endpoint)
	prometheus.RecordAPIRequest(c.Method(), c.Path())

	// ทำงานตามปกติ
	err := c.Next()

	// คำนวณระยะเวลาในการตอบสนอง
	duration := time.Since(start)

	// บันทึกข้อมูล response duration
	prometheus.RecordAPIResponseDuration(c.Method(), c.Path(), duration)

	return err
}
