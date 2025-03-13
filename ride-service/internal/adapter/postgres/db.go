package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/AnwarSaginbai/ride-service/pkg/pb"
	_ "github.com/lib/pq"
)

type Postgres struct {
	db *sql.DB
}

func NewPostgres(db *sql.DB) *Postgres {
	return &Postgres{db: db}
}

func (p *Postgres) CreateRide(ctx context.Context, customerID int, pickup, dropoff string) (int, error) {
	query := `INSERT INTO rides (client_id, status, pickup_location, dropoff_location) 
              VALUES ($1, 'pending', $2, $3) RETURNING id`
	var rideID int

	err := p.db.QueryRowContext(ctx, query, customerID, pickup, dropoff).Scan(&rideID)
	if err != nil {
		return 0, fmt.Errorf("failed to insert ride: %w", err)
	}
	return rideID, nil
}

func (p *Postgres) GetPendingRides(ctx context.Context) ([]*pb.Ride, error) {
	query := `SELECT id, client_id, pickup_location, dropoff_location, status FROM rides WHERE status = 'pending'`
	rows, err := p.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query pending rides: %w", err)
	}
	defer rows.Close()

	rides := make([]*pb.Ride, 0)

	for rows.Next() {
		var ride pb.Ride
		err = rows.Scan(&ride.RideId, &ride.UserId, &ride.PickupLocation, &ride.DropoffLocation, &ride.Status)
		if err != nil {
			return nil, fmt.Errorf("failed to scan ride: %w", err)
		}
		rides = append(rides, &ride)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}

	return rides, nil
}

func (p *Postgres) UpdateStatusOfRide(ctx context.Context, rideID int, newStatus string, driverID int64) error {
	query := `UPDATE rides SET status = $1, driver_id = $2 WHERE id = $3`
	_, err := p.db.ExecContext(ctx, query, newStatus, driverID, rideID)
	if err != nil {
		return fmt.Errorf("failed to update ride status: %w", err)
	}
	return nil
}
