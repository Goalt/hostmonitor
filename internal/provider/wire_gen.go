// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package provider

import (
	"context"
	"github.com/Goalt/hostmonitor/internal/config"
	"github.com/Goalt/hostmonitor/internal/infrastructure/http"
	"github.com/Goalt/hostmonitor/internal/usecase/repository"
)

// Injectors from wire.go:

func InitializeApp(cfg config.Config, context2 context.Context) (Application, func(), error) {
	server := provideServer()
	logger := provideCnfLogger(cfg)
	usecase_repositoryLogger := ProvideLogger(logger)
	application := provideApp(server, cfg, context2, usecase_repositoryLogger)
	return application, func() {
	}, nil
}

// wire.go:

type Application struct {
	ctx context.Context
	log usecase_repository.Logger

	server http.Server
	config config.Config
}

func (a *Application) Run() error {

	go func() {
		err := a.server.Run()
		if err != nil {
			a.log.Error(err)
		}
	}()

	<-a.ctx.Done()

	err := a.server.Stop()
	if err != nil {
		a.log.Error(err)
	}

	return nil
}

func provideApp(server http.Server, cfg config.Config, ctx context.Context, log usecase_repository.Logger) Application {
	return Application{
		server: server,
		ctx:    ctx,
		config: cfg,
		log:    log,
	}
}
