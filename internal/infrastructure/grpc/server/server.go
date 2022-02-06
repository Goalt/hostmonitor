package server

import (
	"context"
	"fmt"
	"net"

	"github.com/Goalt/hostmonitor/internal/config"
	apipb "github.com/Goalt/hostmonitor/internal/infrastructure/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GRPCServer struct {
	cnf config.GRPCServer
	grpcServer *grpc.Server
}

func NewGRPCServer(cnf config.GRPCServer) *GRPCServer {
	return &GRPCServer{
		cnf: cnf,
	}
}

func (s GRPCServer) Run() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", s.cnf.Port))
	if err != nil {
		return err
	}

	s.grpcServer = grpc.NewServer()

	apipb.RegisterHostMonitorServiceServer(s.grpcServer, tmp{})

	return s.grpcServer.Serve(lis)
}

func (s GRPCServer) Stop() {
	s.grpcServer.Stop()
}

type tmp struct {}

func (t tmp) GetStatistics(context.Context, *emptypb.Empty) (*apipb.StatisticsResponse, error) {
	return &apipb.StatisticsResponse{FreeRam: 10}, nil
}