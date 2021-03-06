// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package provider

import (
	"context"
	"github.com/Goalt/hostmonitor/internal/config"
	"github.com/Goalt/hostmonitor/internal/infrastructure/grpc/proxy"
	"github.com/Goalt/hostmonitor/internal/infrastructure/grpc/server"
	"github.com/Goalt/hostmonitor/internal/infrastructure/updater"
	"github.com/Goalt/hostmonitor/internal/usecase/repository"
)

// Injectors from wire.go:

func InitializeApp(cfg config.Config, context2 context.Context) (Application, func(), error) {
	grpcServer := provideCnfGRPCServer(cfg)
	statisticsI := provideStatisticsI()
	handler := provideHandler(statisticsI)
	serverGRPCServer := provideServer(grpcServer, handler)
	proxyServer := provideCnfProxyServer(cfg)
	proxy := provideProxy(proxyServer)
	updater := provideCnfUpdater(cfg)
	logger := provideCnfLogger(cfg)
	usecase_repositoryLogger := ProvideLogger(logger)
	host := provideCnfHost(cfg)
	sshClientFactory := provideSSHFactory()
	hostI := provideHostI(host, sshClientFactory)
	updaterUpdater := provideUpdater(updater, usecase_repositoryLogger, hostI, statisticsI)
	application := provideApp(serverGRPCServer, proxy, updaterUpdater, cfg, context2, usecase_repositoryLogger)
	return application, func() {
	}, nil
}

// wire.go:

type Application struct {
	ctx context.Context
	log usecase_repository.Logger

	server  *server.GRPCServer
	proxy   *proxy.Proxy
	updater *updater.Updater

	config config.Config
}

func (a *Application) Run() error {

	a.log.Info("starting grpc server")
	go func() {
		if err := a.server.Run(); err != nil {
			a.log.Error(err)
		}
	}()

	a.log.Info("starting proxy server")
	go func() {
		if err := a.proxy.Run(a.config.GRPCServer.Port); err != nil {
			a.log.Error(err)
		}
	}()

	go func() {
		a.updater.Run()
	}()

	<-a.ctx.Done()

	a.server.Stop()
	a.log.Info("stopped grpc server")

	if err := a.proxy.Stop(); err != nil {
		return nil
	}
	a.log.Info("stopped proxy server")

	a.updater.Stop()
	a.log.Info("stopped updater")

	return nil
}

func provideApp(server2 *server.GRPCServer, proxy2 *proxy.Proxy, updater2 *updater.Updater,
	cfg config.Config,
	ctx context.Context,
	log usecase_repository.Logger,
) Application {
	return Application{
		server:  server2,
		proxy:   proxy2,
		updater: updater2,
		ctx:     ctx,
		config:  cfg,
		log:     log,
	}
}
