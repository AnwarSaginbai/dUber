package grpc

import (
	"context"
	"github.com/AnwarSaginbai/ride-service/pkg/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (a *Adapter) CreateRide(ctx context.Context, req *pb.CreateRideRequest) (*pb.CreateRideResponse, error) {
	userID, role, err := getUserInfoFromContext(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid token: %v", err)
	}
	log.Printf("Extracted userID: %d, role: %s", userID, role)

	if role != "client" {
		return nil, status.Errorf(codes.PermissionDenied, "only clients can create rides")
	}

	if req.PickupLocation == "" || req.DropoffLocation == "" {
		return nil, status.Errorf(codes.FailedPrecondition, "locations is required")
	}

	rideID, err := a.api.Ride(ctx, userID, req.PickupLocation, req.DropoffLocation)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot create ride: %v", err)
	}
	response := &pb.CreateRideResponse{
		RideId: int64(rideID),
		Status: "pending",
	}
	return response, nil
}
func (a *Adapter) GetPendingRides(ctx context.Context, req *pb.GetPendingRidesRequest) (*pb.GetPendingRidesResponse, error) {
	rides, err := a.api.GetPendingRides(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot get pending rides: %v", err)
	}
	response := &pb.GetPendingRidesResponse{
		Rides: rides,
	}
	return response, nil
}
func (a *Adapter) UpdateRideStatus(ctx context.Context, req *pb.UpdateRideStatusRequest) (*pb.UpdateRideStatusResponse, error) {
	userID, role, err := getUserInfoFromContext(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid token: %v", err)
	}
	log.Printf("Extracted userID: %d, role: %s", userID, role)

	if role != "driver" {
		return nil, status.Errorf(codes.PermissionDenied, "only driver can accept rides")
	}

	if req.Status != "accepted" && req.Status != "completed" {
		return nil, status.Errorf(codes.InvalidArgument, "invalid status: must be 'accepted' or 'completed'")
	}

	if err = a.api.UpdateStatus(ctx, int(req.RideId), req.Status, int64(userID)); err != nil {
		return nil, status.Errorf(codes.Internal, "cannot update ride status: %v", err)
	}
	response := &pb.UpdateRideStatusResponse{
		Status: req.Status,
	}
	return response, nil
}
