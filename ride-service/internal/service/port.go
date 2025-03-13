package service

import (
	"context"
	"github.com/AnwarSaginbai/ride-service/pkg/pb"
)

type Store interface {
	CreateRide(ctx context.Context, customerID int, pickup, dropoff string) (int, error)
	GetPendingRides(ctx context.Context) ([]*pb.Ride, error)
	UpdateStatusOfRide(ctx context.Context, rideID int, newStatus string, driverID int64) error
}
type API interface {
	Ride(ctx context.Context, customerID int, pickup string, dropoff string) (int, error)
	GetPendingRides(ctx context.Context) ([]*pb.Ride, error)
	UpdateStatus(ctx context.Context, rideID int, newStatus string, driverID int64) error
}
