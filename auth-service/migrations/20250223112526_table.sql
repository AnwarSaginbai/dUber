-- +goose Up
-- +goose StatementBegin
CREATE TABLE clients (
                         id SERIAL PRIMARY KEY,
                         first_name VARCHAR(50) NOT NULL,
                         last_name VARCHAR(50) NOT NULL,
                         email VARCHAR(100) UNIQUE NOT NULL,
                         password BYTEA NOT NULL,
                         created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE drivers (
                         id SERIAL PRIMARY KEY,
                         first_name VARCHAR(50) NOT NULL,
                         last_name VARCHAR(50) NOT NULL,
                         email VARCHAR(100) UNIQUE NOT NULL,
                         password BYTEA NOT NULL,
                         car_model VARCHAR(100) NOT NULL,
                         created_at TIMESTAMP DEFAULT NOW()
);

CREATE UNIQUE INDEX idx_clients_email ON clients (email);
CREATE UNIQUE INDEX idx_drivers_email ON drivers (email);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS clients;
DROP TABLE IF EXISTS drivers;
-- +goose StatementEnd
