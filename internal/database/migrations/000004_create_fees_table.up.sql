-- Fee Table
CREATE TABLE fees (
    id SERIAL PRIMARY KEY,
    permit_id INT NOT NULL,
    amount DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (permit_id) REFERENCES permits(id)
);