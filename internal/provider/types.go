package provider

import (
	"github.com/Goalt/hostmonitor/internal/config"
	"github.com/google/wire"
)

func provideCnfLogger(config config.Config) config.Logger {
	return config.Logger
}

var typesSet = wire.NewSet(
	provideCnfLogger,
)