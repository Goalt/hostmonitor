package interactor

import (
	"github.com/Goalt/hostmonitor/internal/entity"
)

type HostI interface {
	GetLastStatisticsFromHost() entity.Statistics
}

type host struct {}

func NewHost() HostI {
	return &host{}
}

func (hi *host) GetLastStatisticsFromHost() entity.Statistics {
	return entity.Statistics{}
}