//go:build wireinject
// +build wireinject

package provider

import (
	"context"

	"github.com/Goalt/hostmonitor/internal/config"
	"github.com/Goalt/hostmonitor/internal/infrastructure/grpc/server"
	"github.com/Goalt/hostmonitor/internal/infrastructure/grpc/proxy"
	usecase_repository "github.com/Goalt/hostmonitor/internal/usecase/repository"
	"github.com/google/wire"
)

type Application struct {
	ctx context.Context
	log usecase_repository.Logger

	server *server.GRPCServer
	proxy  *proxy.Proxy

	config config.Config
}

func (a *Application) Run() error {
	// Server start
	a.log.Info("starting grpc server")
	go func() {
		if err := a.server.Run(); err != nil {
			a.log.Error(err)
		}
	}()

	// Proxy run
	a.log.Info("starting proxy server")
	go func() {
		if err := a.proxy.Run(a.config.GRPCServer.Port); err != nil {
			a.log.Error(err)
		}
	}()

	<-a.ctx.Done()

	//Server stop
	a.server.Stop()
	a.log.Info("stopped grpc server")

	// Proxy stop
	if err := a.proxy.Stop(); err != nil {
		return nil
	}
	a.log.Info("stopped proxy server")

	return nil
}

func provideApp(server *server.GRPCServer, proxy *proxy.Proxy, cfg config.Config, ctx context.Context, log usecase_repository.Logger) Application {
	return Application{
		server: server,
		proxy: proxy,
		ctx:    ctx,
		config: cfg,
		log:    log,
	}
}

func InitializeApp(cfg config.Config, context context.Context) (Application, func(), error) {
	panic(wire.Build(provideApp, infrastructureSet, typesSet))
}
