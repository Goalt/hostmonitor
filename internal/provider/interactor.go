package provider

import (
	"github.com/Goalt/hostmonitor/internal/usecase/interactor"
	"github.com/google/wire"
)

func provideStatisticsI() interactor.StatisticsI {
	return interactor.NewStatistics()
}

var InteractorSet = wire.NewSet(
	provideStatisticsI,
)