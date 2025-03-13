package service

import (
	"context"
	"fmt"
	"github.com/AnwarSaginbai/ride-service/pkg/pb"
)

type Service struct {
	store Store
}

func NewService(store Store) *Service {
	return &Service{
		store: store,
	}
}

func (s *Service) Ride(ctx context.Context, customerID int, pickup string, dropoff string) (int, error) {
	id, err := s.store.CreateRide(ctx, customerID, pickup, dropoff)
	if err != nil {
		return 0, fmt.Errorf("failed to get id in service layer: %v", err)
	}
	return id, nil
}

func (s *Service) GetPendingRides(ctx context.Context) ([]*pb.Ride, error) {
	rides, err := s.store.GetPendingRides(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get pending rides: %v", err)
	}
	return rides, nil
}

func (s *Service) UpdateStatus(ctx context.Context, rideID int, newStatus string, driverID int64) error {
	return s.store.UpdateStatusOfRide(ctx, rideID, newStatus, driverID)
}
