```sql
CREATE TABLE users (
    user_id      UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    full_name    VARCHAR(255) NOT NULL,
    email        VARCHAR(255) UNIQUE NOT NULL,
    phone_number VARCHAR(20) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE accounts (
    account_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id    UUID NOT NULL,
    balance    DECIMAL(18,2) NOT NULL DEFAULT 0.00 CHECK (balance >= 0),
    currency   VARCHAR(10) NOT NULL DEFAULT 'THB',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);

CREATE TABLE withdrawals (
    withdrawal_id  UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    account_id     UUID NOT NULL,
    amount         DECIMAL(18,2) NOT NULL CHECK (amount > 0),
    currency       VARCHAR(10) NOT NULL,
    status         VARCHAR(20) NOT NULL CHECK (status IN ('PENDING', 'COMPLETED', 'FAILED')),
    created_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    completed_at   TIMESTAMP NULL,
    FOREIGN KEY (account_id) REFERENCES accounts(account_id)
);

CREATE TABLE transactions (
    transaction_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    from_account_id UUID NOT NULL,
    to_account_id UUID NOT NULL,
    amount DECIMAL(18,2) NOT NULL CHECK (amount > 0),
    currency VARCHAR(10) NOT NULL,
    status VARCHAR(20) NOT NULL CHECK (status IN ('PENDING', 'COMPLETED', 'FAILED')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    completed_at TIMESTAMP NULL,
    FOREIGN KEY (from_account_id) REFERENCES accounts(account_id),
    FOREIGN KEY (to_account_id) REFERENCES accounts(account_id),
    CHECK (from_account_id <> to_account_id) -- ป้องกันการโอนเงินให้ตัวเอง
);

CREATE TABLE transaction_logs (
    log_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    transaction_id UUID NOT NULL,
    log_message TEXT NOT NULL,
    log_type VARCHAR(10) NOT NULL CHECK (log_type IN ('INFO', 'ERROR', 'WARNING')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (transaction_id) REFERENCES transactions(transaction_id)
);
```





```
go-payment-system
├─ README.md
├─ cmd
│  ├─ consumer
│  │  └─ main.go
│  └─ service
│     └─ main.go
├─ config
│  ├─ config.go
│  └─ config.yaml
├─ deploy
│  ├─ Dockerfile.consumer
│  ├─ Dockerfile.service
│  ├─ docker-compose.yml
│  └─ k8s.yaml
├─ go.mod
├─ go.sum
├─ internal
│  ├─ api
│  │  ├─ handlers
│  │  │  └─ transaction_handler.go
│  │  ├─ models
│  │  │  └─ transaction.go
│  │  ├─ routes
│  │  │  └─ router.go
│  │  └─ services
│  │     └─ transaction_service.go
│  ├─ consumers
│  │  ├─ handlers
│  │  │  └─ transaction_handler.go
│  │  ├─ models
│  │  │  └─ transaction.go
│  │  └─ services
│  │     └─ transaction_service.go
│  └─ repository
│     ├─ models
│     │  ├─ account.go
│     │  ├─ transaction.go
│     │  ├─ user.go
│     │  └─ withdrawal.go
│     └─ transaction_repository.go
└─ pkg
   ├─ database
   │  └─ postgres.go
   ├─ kafka
   │  ├─ consumer.go
   │  └─ producer.go
   ├─ logger
   │  └─ logger.go
   ├─ middleware
   └─ utils

```