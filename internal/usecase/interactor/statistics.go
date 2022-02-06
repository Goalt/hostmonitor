package interactor

import (
	"sync"

	"github.com/Goalt/hostmonitor/internal/entity"
)

type StatisticsI interface {
	GetLast() entity.Statistics
	Update(entity.Statistics)
}

type statisticsInteractor struct {
	data entity.Statistics
	mu   sync.Mutex
}

func (si *statisticsInteractor) GetLast() entity.Statistics {
	si.mu.Lock()
	defer si.mu.Unlock()

	return si.data
}

func (si *statisticsInteractor) Update(new entity.Statistics) {
	si.mu.Lock()
	defer si.mu.Unlock()

	si.data = new
}

func NewStatistics() StatisticsI {
	return &statisticsInteractor{}
}
