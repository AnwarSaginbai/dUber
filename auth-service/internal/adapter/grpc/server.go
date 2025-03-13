package grpc

import (
	"fmt"
	"github.com/AnwarSaginbai/auth-service/internal/config"
	"github.com/AnwarSaginbai/auth-service/internal/service"
	"github.com/AnwarSaginbai/auth-service/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type Adapter struct {
	api service.API
	cfg *config.Config
	pb.UnimplementedAuthServiceServer
}

func NewAdapter(cfg *config.Config, api service.API) *Adapter {
	return &Adapter{cfg: cfg, api: api}
}

func (a *Adapter) Run() error {
	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, a)
	reflection.Register(grpcServer)

	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", a.cfg.Server.Host, a.cfg.Server.Port))
	if err != nil {
		return err
	}
	grpcServer.Serve(listen)
	return nil
}
