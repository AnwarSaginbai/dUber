package service

import (
	"context"
	"errors"
	"github.com/AnwarSaginbai/auth-service/internal/config"
	"github.com/AnwarSaginbai/auth-service/internal/token"
	"golang.org/x/crypto/bcrypt"
)

type API interface {
	RegisterNewClient(ctx context.Context, firstname, lastname, email, password string) (int64, error)
	RegisterNewDriver(ctx context.Context, firstname, lastname, email, password, carModel string) (int64, error)
	Login(ctx context.Context, email, password string) (string, error)
}

type Service struct {
	store Storage
	cfg   *config.Config
}

func NewAPI(store Storage) *Service {
	return &Service{store: store}
}

func (s *Service) RegisterNewClient(ctx context.Context, firstname, lastname, email, password string) (int64, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	id, err := s.store.CreateClient(ctx, firstname, lastname, email, string(hash))
	if err != nil {
		return 0, err
	}
	return id, nil
}
func (s *Service) RegisterNewDriver(ctx context.Context, firstname, lastname, email, password, carModel string) (int64, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}
	id, err := s.store.CreateDriver(ctx, firstname, lastname, email, carModel, string(hash))
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *Service) Login(ctx context.Context, email, password string) (string, error) {
	client, err := s.store.GetClientByEmail(ctx, email)
	if err == nil {
		if err := bcrypt.CompareHashAndPassword(client.Password, []byte(password)); err == nil {
			return token.GenerateToken(client.ID, client.Email, "client")
		}
	}

	driver, err := s.store.GetDriverByEmail(ctx, email)
	if err == nil {
		if err = bcrypt.CompareHashAndPassword(driver.Password, []byte(password)); err == nil {
			return token.GenerateToken(driver.ID, driver.Email, "driver")
		}
	}

	return "", errors.New("invalid email or password")
}
