-- Create database (PostgreSQL version)
-- NOTE: In Postgres you usually create DB separately (psql or admin tool)
-- CREATE DATABASE drukdeals_db;

-- Connect using:
-- \c drukdeals_db

CREATE TABLE IF NOT EXISTS users (
    user_id SERIAL PRIMARY KEY,
    fullname VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS categories (
    cat_id SERIAL PRIMARY KEY,
    cat_name VARCHAR(50) UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS products (
    prod_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    cat_id INT NOT NULL,
    title VARCHAR(100) NOT NULL,
    price NUMERIC(10,2) NOT NULL,
    description TEXT,
    image_path VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    FOREIGN KEY (cat_id) REFERENCES categories(cat_id)
);

INSERT INTO categories (cat_name) VALUES 
('Electronics'), 
('Clothing'), 
('Books'), 
('Furniture'), 
('Other')
ON CONFLICT (cat_name) DO NOTHING;

-- Instead of SELECT 'message'
SELECT 'Database setup complete!' AS status;