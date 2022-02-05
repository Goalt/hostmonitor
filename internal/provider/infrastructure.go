package provider

import (
	"context"
	"net"
	"time"

	"github.com/Goalt/FileSharer/internal/config"
	gorm_repository "github.com/Goalt/FileSharer/internal/infrastructure/gorm"
	"github.com/Goalt/FileSharer/internal/infrastructure/http"
	"github.com/Goalt/FileSharer/internal/infrastructure/logger"
	infrastructure_repository "github.com/Goalt/FileSharer/internal/infrastructure/repository"
	"github.com/Goalt/FileSharer/internal/interface/controller"
	usecase_repository "github.com/Goalt/FileSharer/internal/usecase/repository"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLog "gorm.io/gorm/logger"
)

type ServicesCleanup func()

func provideServer(cfg config.Server, controller controller.HTTPController) http.Server {
	return http.NewHTTPServer(cfg.Port, controller)
}

func ProvideLogger(cnf config.Logger) usecase_repository.Logger {
	return logger.NewLogger(cnf)
}

func provideServicesCleanup(cleanup func()) ServicesCleanup {
	return cleanup
}

var infrastructureSet = wire.NewSet(
	provideServer,
	ProvideLogger,
)