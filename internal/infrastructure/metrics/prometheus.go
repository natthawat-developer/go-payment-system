package prometheus

import (
	"go-payment-system/internal/infrastructure/metrics"
	"time"
)

// RecordTransactionSuccess สำหรับบันทึกการทำธุรกรรมที่สำเร็จ
func RecordTransactionSuccess(status string) {
	metrics.SuccessfulTransactions.WithLabelValues(status).Inc()
}

// RecordTransactionFailure สำหรับบันทึกการทำธุรกรรมที่ล้มเหลว
func RecordTransactionFailure(status string) {
	metrics.FailedTransactions.WithLabelValues(status).Inc()
}

// RecordAPIRequest สำหรับบันทึกการร้องขอ API
func RecordAPIRequest(method, endpoint string) {
	metrics.APIRequests.WithLabelValues(method, endpoint).Inc()
}

// RecordAPIResponseDuration สำหรับบันทึกระยะเวลาในการตอบสนอง API
func RecordAPIResponseDuration(method, endpoint string, duration time.Duration) {
	metrics.APIResponseDuration.WithLabelValues(method, endpoint).Observe(duration.Seconds())
}
