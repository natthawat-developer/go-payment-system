package metrics

import (
	"log"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

var (
	// ตัวนับจำนวนธุรกรรมที่สำเร็จ
	SuccessfulTransactions = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "successful_transactions_total",
			Help: "Total number of successful transactions.",
		},
		[]string{"status"}, // เช่น สถานะต่างๆ (success, pending)
	)

	// ตัวนับจำนวนธุรกรรมที่ล้มเหลว
	FailedTransactions = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "failed_transactions_total",
			Help: "Total number of failed transactions.",
		},
		[]string{"status"}, // เช่น สถานะต่างๆ (failed, error)
	)

	// ตัวนับการร้องขอ API
	APIRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "api_requests_total",
			Help: "Total number of API requests.",
		},
		[]string{"method", "endpoint"}, // method (GET, POST) และ endpoint
	)

	// ตัวจับเวลาเวลาในการตอบสนอง
	APIResponseDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "api_response_duration_seconds",
			Help:    "Histogram of API response durations in seconds.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "endpoint"}, // method และ endpoint
	)
)

// RegisterMetrics ทำการลงทะเบียน metrics
func RegisterMetrics() {
	// ลงทะเบียนตัวนับและตัวจับเวลาทั้งหมด
	prometheus.MustRegister(SuccessfulTransactions)
	prometheus.MustRegister(FailedTransactions)
	prometheus.MustRegister(APIRequests)
	prometheus.MustRegister(APIResponseDuration)
}

// StartMetricsServer เริ่มต้น Prometheus HTTP server เพื่อให้สามารถดึงข้อมูล metrics
func StartMetricsServer(port string) {
	http.Handle("/metrics", promhttp.Handler()) // กำหนดให้ Prometheus สามารถดึง metrics ได้ที่ /metrics
	log.Printf("Starting metrics server on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
