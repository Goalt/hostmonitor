package updater

import (
	"time"

	"github.com/Goalt/hostmonitor/internal/config"
	"github.com/Goalt/hostmonitor/internal/usecase/interactor"
	"github.com/Goalt/hostmonitor/internal/usecase/repository"
)

type Updater struct {
	log   usecase_repository.Logger
	hostI interactor.HostI
	statI interactor.StatisticsI

	stopCh chan interface{}
	cnf    config.Updater
}

func NewUpdater(cnf config.Updater, log usecase_repository.Logger, hostI interactor.HostI, statI interactor.StatisticsI) *Updater {
	return &Updater{
		cnf:   cnf,
		log:   log,
		hostI: hostI,
		statI: statI,
	}
}

func (u *Updater) Run() {
	tick := time.NewTicker(u.cnf.Interval)
L:
	for {
		select {
		case <-u.stopCh:
			break L
		case <-tick.C:
			newStat := u.hostI.GetLastStatisticsFromHost()
			u.statI.Update(newStat)
		}
	}
}

func (u *Updater) Stop() {
	u.stopCh <- nil
}
