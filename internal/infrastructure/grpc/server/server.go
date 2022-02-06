package server

import (
	"fmt"
	"net"

	"github.com/Goalt/hostmonitor/internal/config"
	apipb "github.com/Goalt/hostmonitor/internal/infrastructure/grpc/proto"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	cnf config.GRPCServer
	grpcServer *grpc.Server
	handler *Handler 
}

func NewGRPCServer(cnf config.GRPCServer, hd *Handler) *GRPCServer {
	return &GRPCServer{
		cnf: cnf,
		handler: hd,
	}
}

func (s GRPCServer) Run() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", s.cnf.Port))
	if err != nil {
		return err
	}

	s.grpcServer = grpc.NewServer()

	apipb.RegisterHostMonitorServiceServer(s.grpcServer, s.handler)

	return s.grpcServer.Serve(lis)
}

func (s GRPCServer) Stop() {
	s.grpcServer.Stop()
}