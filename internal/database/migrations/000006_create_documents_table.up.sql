-- Document Table
CREATE TABLE documents (
    id SERIAL PRIMARY KEY,
    permit_id INT NOT NULL,
    document_url VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (permit_id) REFERENCES permits(id) 
);