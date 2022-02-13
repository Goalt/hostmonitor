package provider

import (
	"github.com/Goalt/hostmonitor/internal/config"
	"github.com/Goalt/hostmonitor/internal/usecase/interactor"
	usecase_repository "github.com/Goalt/hostmonitor/internal/usecase/repository"
	"github.com/google/wire"
)

func provideStatisticsI() interactor.StatisticsI {
	return interactor.NewStatistics()
}

func provideHostI(cnf config.Host, sshFactory usecase_repository.SSHClientFactory) interactor.HostI {
	return interactor.NewHost(cnf, sshFactory)
}

var InteractorSet = wire.NewSet(
	provideStatisticsI,
	provideHostI,
)
