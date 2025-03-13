package grpc

import (
	"context"
	"github.com/AnwarSaginbai/auth-service/pkg/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *Adapter) RegisterClient(ctx context.Context, req *pb.RegisterClientRequest) (*pb.RegisterResponse, error) {
	id, err := a.api.RegisterNewClient(ctx, req.FirstName, req.LastName, req.Email, req.Password)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	response := &pb.RegisterResponse{
		Id:      id,
		Message: "Успешно зарегистрировался",
	}
	return response, nil
}
func (a *Adapter) RegisterDriver(ctx context.Context, req *pb.RegisterDriverRequest) (*pb.RegisterResponse, error) {
	id, err := a.api.RegisterNewDriver(ctx, req.FirstName, req.LastName, req.Email, req.Password, req.CarModel)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	response := &pb.RegisterResponse{
		Id:      id,
		Message: "Успешно зарегистрировался",
	}
	return response, nil
}
func (a *Adapter) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	token, err := a.api.Login(ctx, req.Email, req.Password)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	response := &pb.LoginResponse{
		Token: token,
	}
	return response, nil
}
func (a *Adapter) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {

	response := &pb.GetUserResponse{
		Id:        0,
		FirstName: "",
		LastName:  "",
		Email:     "",
		Role:      "",
		CarModel:  "",
	}
	return response, nil
}
