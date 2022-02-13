package interactor

import (
	"strconv"
	"strings"
	"time"

	"github.com/Goalt/hostmonitor/internal/config"
	"github.com/Goalt/hostmonitor/internal/entity"
	"github.com/Goalt/hostmonitor/internal/usecase/repository"
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

	storage, err := hi.getStorage(sshClient)
	if err != nil {
		return entity.Statistics{}, err
	}

	loadAvg, err := hi.getLoadAvg(sshClient)
	if err != nil {
		return entity.Statistics{}, err
	}

	uptime, err := hi.getUptime(sshClient)
	if err != nil {
		return entity.Statistics{}, err
	}

	return entity.Statistics{
		Ram:       ram,
		Storage:   storage,
		LoadAvg:   loadAvg,
		Uptime:    uptime,
		UpdatedAt: time.Now(),
	}, nil
}

func (hi *host) getRam(cl usecase_repository.SSHClient) (entity.Ram, error) {
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

func (hi *host) getStorage(cl usecase_repository.SSHClient) (entity.Storage, error) {
	storageRaw, err := cl.Execute("df | grep '/dev/sda1 ' | awk -F ' ' '{print $2 \" \" $3 \" \" $4}'")
	if err != nil {
		return entity.Storage{}, err
	}
	storageRaw = strings.Trim(storageRaw, "\n")

	sliced := strings.Split(storageRaw, " ")

	total, err := strconv.Atoi(sliced[0])
	if err != nil {
		return entity.Storage{}, err
	}

	available, err := strconv.Atoi(sliced[1])
	if err != nil {
		return entity.Storage{}, err
	}

	free, err := strconv.Atoi(sliced[2])
	if err != nil {
		return entity.Storage{}, err
	}

	return entity.Storage{
		Total:     total,
		Available: available,
		Free:      free,
	}, nil
}

func (hi *host) getLoadAvg(cl usecase_repository.SSHClient) (entity.LoadAvg, error) {
	loadRaw, err := cl.Execute("uptime | awk -F ' ' '{print $(NF-2) \" \" $(NF-1) \" \" $NF}'")
	if err != nil {
		return entity.LoadAvg{}, err
	}
	loadRaw = strings.Trim(loadRaw, "\n")

	sliced := strings.Split(loadRaw, ", ")

	load1, err := strconv.ParseFloat(sliced[0], 64)
	if err != nil {
		return entity.LoadAvg{}, err
	}

	load5, err := strconv.ParseFloat(sliced[1], 64)
	if err != nil {
		return entity.LoadAvg{}, err
	}

	load15, err := strconv.ParseFloat(sliced[2], 64)
	if err != nil {
		return entity.LoadAvg{}, err
	}

	return entity.LoadAvg{
		Load1:  load1,
		Load5:  load5,
		Load15: load15,
	}, nil
}

func (hi *host) getUptime(cl usecase_repository.SSHClient) (entity.Uptime, error) {
	uptimeRaw, err := cl.Execute("awk '{print $1}' /proc/uptime")
	if err != nil {
		return entity.Uptime{}, err
	}
	uptimeRaw = strings.Trim(uptimeRaw, "\n")

	uptime, err := strconv.ParseFloat(uptimeRaw, 64)
	if err != nil {
		return entity.Uptime{}, err
	}

	return entity.Uptime{
		Dur: int(uptime),
	}, nil
}
