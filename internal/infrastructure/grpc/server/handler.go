package server

import (
	"context"

	apipb "github.com/Goalt/hostmonitor/internal/infrastructure/grpc/proto"
	"github.com/Goalt/hostmonitor/internal/usecase/interactor"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Handler struct {
	statisticsI interactor.StatisticsI
}

func NewHandler(statisticsI interactor.StatisticsI) *Handler {
	return &Handler{
		statisticsI: statisticsI,
	}
}

func (h *Handler) GetStatistics(context.Context, *emptypb.Empty) (*apipb.StatisticsResponse, error) {
	statistics := h.statisticsI.GetLast()

	resp := apipb.StatisticsResponse{
		Ram: &apipb.Ram{
			Total:     uint64(statistics.Ram.Total),
			Available: uint64(statistics.Ram.Available),
		},
		Storage: &apipb.Storage{
			Total:     uint64(statistics.Storage.Total),
			Available: uint64(statistics.Storage.Available),
			Free:      uint64(statistics.Storage.Free),
		},
		Loadavg: &apipb.LoadAvg{
			Load1:  float32(statistics.LoadAvg.Load1),
			Load5:  float32(statistics.LoadAvg.Load5),
			Load15: float32(statistics.LoadAvg.Load15),
		},
		Uptime: &apipb.Uptime{
			Dur: float32(statistics.Uptime.Dur),
		},
		DockerContainers: statistics.DockerContainers.Statuses,
		UpdatedAt:        uint64(statistics.UpdatedAt.Unix()),
	}

	return &resp, nil
}
