-- Payment Table
CREATE TABLE payments (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    permit_id INT NOT NULL,
    amount DECIMAL(10, 2) NOT NULL,
    status VARCHAR(50) NOT NULL, -- e.g., 'success', 'failed', 'pending'
    payment_method VARCHAR(50) NOT NULL, -- e.g., 'credit_card', 'bank_transfer'
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (permit_id) REFERENCES permits(id)
);