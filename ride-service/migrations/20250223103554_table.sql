-- +goose Up
-- +goose StatementBegin
CREATE TABLE rides (
                       id SERIAL PRIMARY KEY,
                       client_id INT NOT NULL,
                       driver_id INT,
                       status TEXT NOT NULL CHECK (status IN ('pending', 'accepted', 'completed', 'cancelled')),
                       pickup_location TEXT NOT NULL,
                       dropoff_location TEXT NOT NULL,
                       created_at TIMESTAMP DEFAULT NOW(),
                       updated_at TIMESTAMP DEFAULT NOW()
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE rides;
-- +goose StatementEnd
