-- ขึ้นตาราง payment_logs
CREATE TABLE payment_logs (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    transaction_id INT,
    action VARCHAR(255) NOT NULL,
    amount DECIMAL(15, 2),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (transaction_id) REFERENCES transactions(id) ON DELETE SET NULL
);
