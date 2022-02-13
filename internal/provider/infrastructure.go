package provider

import (
	"github.com/Goalt/hostmonitor/internal/config"
	"github.com/Goalt/hostmonitor/internal/infrastructure/grpc/proxy"
	"github.com/Goalt/hostmonitor/internal/infrastructure/grpc/server"
	"github.com/Goalt/hostmonitor/internal/infrastructure/logger"
	"github.com/Goalt/hostmonitor/internal/infrastructure/updater"
	"github.com/Goalt/hostmonitor/internal/infrastructure/repository"
	"github.com/Goalt/hostmonitor/internal/usecase/interactor"
	usecase_repository "github.com/Goalt/hostmonitor/internal/usecase/repository"
	"github.com/google/wire"
)

func ProvideLogger(cnf config.Logger) usecase_repository.Logger {
	return logger.NewLogger(cnf)
}

func provideServer(cnf config.GRPCServer, hndl *server.Handler) *server.GRPCServer {
	return server.NewGRPCServer(cnf, hndl)
}

func provideHandler(interactor interactor.StatisticsI) *server.Handler {
	return server.NewHandler(interactor)
}

func provideProxy(cnf config.ProxyServer) *proxy.Proxy {
	return proxy.NewProxy(cnf)
}

func provideUpdater(cnf config.Updater, log usecase_repository.Logger, hostI interactor.HostI, statI interactor.StatisticsI) *updater.Updater {
	return updater.NewUpdater(cnf, log, hostI, statI)
}

func provideSSHFactory() usecase_repository.SSHClientFactory {
	return repository.NewSSHFactory()
}

var InfrastructureSet = wire.NewSet(
	ProvideLogger,
	provideServer,
	provideProxy,
	provideHandler,
	provideUpdater,
	provideSSHFactory,
)
