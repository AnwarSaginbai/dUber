package service

import (
	"context"
	"github.com/AnwarSaginbai/auth-service/internal/domain"
)

type Storage interface {
	CreateClient(ctx context.Context, firstname, lastname, email, password string) (int64, error)
	CreateDriver(ctx context.Context, firstname, lastname, email, carmodel, password string) (int64, error)

	GetClientByID(ctx context.Context, id int64) (*domain.UserClient, error)
	GetDriverByID(ctx context.Context, id int64) (*domain.UserDriver, error)

	GetClientByEmail(ctx context.Context, email string) (*domain.UserClient, error)
	GetDriverByEmail(ctx context.Context, email string) (*domain.UserDriver, error)
}
