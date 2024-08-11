-- Permit Table
CREATE TABLE permits (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    description TEXT NOT NULL,
    status VARCHAR(50) NOT NULL, -- e.g., 'pending', 'approved', 'rejected', 'paid'
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) 
);