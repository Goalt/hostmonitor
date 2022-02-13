package provider

import (
	"github.com/Goalt/hostmonitor/internal/config"
	"github.com/google/wire"
)

func provideCnfLogger(config config.Config) config.Logger {
	return config.Logger
}

func provideCnfGRPCServer(cnf config.Config) config.GRPCServer {
	return cnf.GRPCServer
}

func provideCnfProxyServer(cnf config.Config) config.ProxyServer {
	return cnf.ProxyServer
}

func provideCnfUpdater(cnf config.Config) config.Updater {
	return cnf.Updater
}

var TypesSet = wire.NewSet(
	provideCnfLogger,
	provideCnfGRPCServer,
	provideCnfProxyServer,
	provideCnfUpdater,
)