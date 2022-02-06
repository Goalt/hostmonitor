package interactor

import "github.com/Goalt/hostmonitor/internal/entity"

type StatisticsI interface {
	GetLast() entity.Statistics
	Update(entity.Statistics)
}

type statisticsInteractor struct {
}

func (si *statisticsInteractor) GetLast() entity.Statistics {
	return entity.Statistics{FreeRam: 11}
}

func (si *statisticsInteractor) Update(entity.Statistics) {

}

func NewStatistics() StatisticsI {
	return &statisticsInteractor{}
}
