package provider

import (
	"github.com/Goalt/hostmonitor/internal/config"
	"github.com/Goalt/hostmonitor/internal/infrastructure/grpc/proxy"
	"github.com/Goalt/hostmonitor/internal/infrastructure/grpc/server"
	"github.com/Goalt/hostmonitor/internal/infrastructure/logger"
	usecase_repository "github.com/Goalt/hostmonitor/internal/usecase/repository"
	"github.com/google/wire"
)

func ProvideLogger(cnf config.Logger) usecase_repository.Logger {
	return logger.NewLogger(cnf)
}

func provideServer(cnf config.GRPCServer) *server.GRPCServer {
	return server.NewGRPCServer(cnf)
}

func provideProxy(cnf config.ProxyServer) *proxy.Proxy {
	return proxy.NewProxy(cnf)
}

var infrastructureSet = wire.NewSet(
	ProvideLogger,
	provideServer,
	provideProxy,
)