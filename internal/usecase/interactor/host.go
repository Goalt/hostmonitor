package interactor

import (
	"strconv"
	"strings"
	"time"

	"github.com/Goalt/hostmonitor/internal/config"
	"github.com/Goalt/hostmonitor/internal/entity"
	usecase_repository "github.com/Goalt/hostmonitor/internal/usecase/repository"
)

type HostI interface {
	GetLastStatisticsFromHost() (entity.Statistics, error)
}

type host struct {
	sshFactory usecase_repository.SSHClientFactory

	cnf config.Host
}

func NewHost(cnf config.Host, sshFactory usecase_repository.SSHClientFactory) HostI {
	return &host{
		cnf:        cnf,
		sshFactory: sshFactory,
	}
}

func (hi *host) GetLastStatisticsFromHost() (entity.Statistics, error) {
	sshClient, err := hi.sshFactory.NewSSHClient(hi.cnf.User, hi.cnf.Address, hi.cnf.Password)
	if err != nil {
		return entity.Statistics{}, err
	}

	ram, err := hi.getRam(sshClient)
	if err != nil {
		return entity.Statistics{}, err
	}

	return entity.Statistics{
		Ram: ram,
		UpdatedAt: time.Now(),
	}, nil
}

func (hi *host) getRam(cl usecase_repository.SSHClient) (entity.Ram, error) {
	// Ram
	availableRaw, err := cl.Execute("cat /proc/meminfo | grep 'Available' | awk -F ' ' '{print $2}'")
	if err != nil {
		return entity.Ram{}, err
	}

	available, err := strconv.Atoi(strings.Trim(availableRaw, "\n"))
	if err != nil {
		return entity.Ram{}, nil
	}

	totalRaw, err := cl.Execute("cat /proc/meminfo | grep 'MemTotal' | awk -F ' ' '{print $2}'")
	if err != nil {
		return entity.Ram{}, err
	}

	total, err := strconv.Atoi(strings.Trim(strings.Trim(totalRaw, "\n"), "\n"))
	if err != nil {
		return entity.Ram{}, nil
	}

	return entity.Ram{
		Total:     total,
		Available: available,
	}, nil
}
