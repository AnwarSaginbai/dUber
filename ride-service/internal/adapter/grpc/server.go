package grpc

import (
	"fmt"
	"github.com/AnwarSaginbai/ride-service/internal/config"
	"github.com/AnwarSaginbai/ride-service/internal/service"
	"github.com/AnwarSaginbai/ride-service/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log/slog"
	"net"
)

type Adapter struct {
	api service.API
	cfg *config.Config
	log *slog.Logger
	pb.UnimplementedRideServiceServer
}

func NewAdapter(api service.API, log *slog.Logger, cfg *config.Config) *Adapter {
	return &Adapter{api: api, log: log, cfg: cfg}
}

func (a *Adapter) Run() error {

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(AuthInterceptor))
	pb.RegisterRideServiceServer(grpcServer, a)
	reflection.Register(grpcServer)

	listen, err := net.Listen("tcp", fmt.Sprintf(
		"%s:%d",
		a.cfg.Server.Host,
		a.cfg.Server.Port,
	),
	)

	if err != nil {
		a.log.Info("failed to listen", "error", err)
	}

	return grpcServer.Serve(listen)
}
