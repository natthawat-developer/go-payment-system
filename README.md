
```
go-payment-system
├─ Dockerfile
├─ README.md
├─ cmd
│  ├─ payment-service
│  │  └─ main.go
│  └─ worker-service
│     └─ main.go
├─ deployments
│  ├─ docker-compose.yaml
│  └─ k8s
├─ go.mod
├─ go.sum
├─ internal
│  ├─ adapters
│  │  ├─ database
│  │  │  └─ db.go
│  │  ├─ http
│  │  │  └─ handlers.go
│  │  ├─ kafka
│  │  │  └─ consumer.go
│  │  └─ redis
│  │     └─ cache.go
│  ├─ domain
│  │  └─ models.go
│  ├─ infrastructure
│  │  ├─ config
│  │  │  ├─ config.go
│  │  │  └─ config.yaml
│  │  ├─ logger
│  │  └─ metrics
│  └─ usecase
│     └─ process_payment.go
├─ migrations
└─ test

```