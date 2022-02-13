package interactor

import (
	"strconv"
	"strings"

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
		cnf: cnf,
		sshFactory: sshFactory,
	}
}

func (hi *host) GetLastStatisticsFromHost() (entity.Statistics, error) {
	sshClient, err := hi.sshFactory.NewSSHClient(hi.cnf.User, hi.cnf.Address, hi.cnf.Password)
	if err != nil {
		return entity.Statistics{}, err
	}

	result, err := sshClient.Execute("cat /proc/meminfo | grep 'Available' | awk -F ' ' '{print $2}'")
	if err != nil {
		return entity.Statistics{}, err
	}

	result = strings.Trim(result, "\n")

	ram, err := strconv.Atoi(result)
	if err != nil {
		return entity.Statistics{}, err
	}

	return entity.Statistics{FreeRam: ram}, nil
}