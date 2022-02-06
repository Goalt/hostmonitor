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
		FreeRam: uint64(statistics.FreeRam),
	}

	return &resp, nil
}
