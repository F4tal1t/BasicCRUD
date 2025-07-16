-- Car Inventory Database Schema
-- This file contains the database schema for the Car Inventory System

-- Create database (run this command separately in psql)
-- CREATE DATABASE mycarsdb;

-- Connect to the database
-- \c mycarsdb;

-- Create the cars table
CREATE TABLE IF NOT EXISTS cars (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    model VARCHAR(255) NOT NULL,
    brand VARCHAR(255) NOT NULL,
    year INTEGER NOT NULL CHECK (year > 1900 AND year <= EXTRACT(YEAR FROM CURRENT_DATE) + 1),
    price DECIMAL(10, 2) NOT NULL CHECK (price >= 0)
);

-- Create indexes for better performance
CREATE INDEX IF NOT EXISTS idx_cars_brand ON cars(brand);
CREATE INDEX IF NOT EXISTS idx_cars_year ON cars(year);
CREATE INDEX IF NOT EXISTS idx_cars_price ON cars(price);

-- Insert sample data (optional)
INSERT INTO cars (name, model, brand, year, price) VALUES
    ('Corolla', 'XLE', 'Toyota', 2023, 25000.00),
    ('Civic', 'LX', 'Honda', 2022, 23000.00),
    ('Model S', 'Plaid', 'Tesla', 2023, 89000.00),
    ('Mustang', 'GT', 'Ford', 2022, 35000.00),
    ('Camry', 'SE', 'Toyota', 2023, 28000.00)
ON CONFLICT DO NOTHING;

-- Grant permissions (adjust as needed)
-- GRANT ALL PRIVILEGES ON DATABASE mycarsdb TO your_username;
-- GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO your_username;
-- GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO your_username;
