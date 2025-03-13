package domain

import "time"

type Ride struct {
	ID         int
	CustomerID int
	DriverID   *int
	Status     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
