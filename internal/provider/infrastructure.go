package provider

import (
	"github.com/Goalt/hostmonitor/internal/config"
	"github.com/Goalt/hostmonitor/internal/infrastructure/logger"
	"github.com/Goalt/hostmonitor/internal/infrastructure/http"
	usecase_repository "github.com/Goalt/hostmonitor/internal/usecase/repository"
	"github.com/google/wire"
)

func ProvideLogger(cnf config.Logger) usecase_repository.Logger {
	return logger.NewLogger(cnf)
}

func provideServer() http.Server {
	return http.NewHTTPServer()
}

var infrastructureSet = wire.NewSet(
	ProvideLogger,
	provideServer,
)