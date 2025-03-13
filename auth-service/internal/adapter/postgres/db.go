package postgres

import (
	"context"
	"database/sql"
	"errors"
	"github.com/AnwarSaginbai/auth-service/internal/domain"
	"github.com/AnwarSaginbai/auth-service/internal/service"
	_ "github.com/lib/pq"
)

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage(db *sql.DB) service.Storage {
	return &PostgresStorage{db: db}
}

func (s *PostgresStorage) CreateClient(ctx context.Context, firstname, lastname, email, password string) (int64, error) {
	query := `INSERT INTO clients (first_name, last_name, email, password) VALUES ($1, $2, $3, $4) RETURNING id`
	var id int64
	err := s.db.QueryRowContext(ctx, query, firstname, lastname, email, password).Scan(&id)
	return id, err
}

func (s *PostgresStorage) CreateDriver(ctx context.Context, firstname, lastname, email, carmodel, password string) (int64, error) {
	query := `INSERT INTO drivers (first_name, last_name, email, password, car_model) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	var id int64
	err := s.db.QueryRowContext(ctx, query, firstname, lastname, email, password, carmodel).Scan(&id)
	return id, err
}

func (s *PostgresStorage) GetClientByID(ctx context.Context, id int64) (*domain.UserClient, error) {
	query := `SELECT id, first_name, last_name, email, password FROM clients WHERE id = $1`
	client := &domain.UserClient{}
	err := s.db.QueryRowContext(ctx, query, id).Scan(&client.ID, &client.FirstName, &client.LastName, &client.Email, &client.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("client not found")
		}
		return nil, err
	}
	return client, nil
}

// Получение водителя по ID
func (s *PostgresStorage) GetDriverByID(ctx context.Context, id int64) (*domain.UserDriver, error) {
	query := `SELECT id, first_name, last_name, email, password, car_model FROM drivers WHERE id = $1`
	driver := &domain.UserDriver{}
	err := s.db.QueryRowContext(ctx, query, id).Scan(&driver.ID, &driver.FirstName, &driver.LastName, &driver.Email, &driver.Password, &driver.CarModel)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("driver not found")
		}
		return nil, err
	}
	return driver, nil
}

// Получение клиента по Email
func (s *PostgresStorage) GetClientByEmail(ctx context.Context, email string) (*domain.UserClient, error) {
	query := `SELECT id, first_name, last_name, email, password FROM clients WHERE email = $1`
	client := &domain.UserClient{}
	err := s.db.QueryRowContext(ctx, query, email).Scan(&client.ID, &client.FirstName, &client.LastName, &client.Email, &client.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("client not found")
		}
		return nil, err
	}
	return client, nil
}

// Получение водителя по Email
func (s *PostgresStorage) GetDriverByEmail(ctx context.Context, email string) (*domain.UserDriver, error) {
	query := `SELECT id, first_name, last_name, email, password, car_model FROM drivers WHERE email = $1`
	driver := &domain.UserDriver{}
	err := s.db.QueryRowContext(ctx, query, email).Scan(&driver.ID, &driver.FirstName, &driver.LastName, &driver.Email, &driver.Password, &driver.CarModel)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("driver not found")
		}
		return nil, err
	}
	return driver, nil
}
