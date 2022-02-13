package provider

import (
	"github.com/Goalt/hostmonitor/internal/usecase/interactor"
	"github.com/google/wire"
)

func provideStatisticsI() interactor.StatisticsI {
	return interactor.NewStatistics()
}

func provideHostI() interactor.HostI {
	return interactor.NewHost()
}

var InteractorSet = wire.NewSet(
	provideStatisticsI,
	provideHostI,
)
