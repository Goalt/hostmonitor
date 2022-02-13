package server

import (
	"context"
	"testing"
	"time"

	"github.com/Goalt/hostmonitor/internal/entity"
	apipb "github.com/Goalt/hostmonitor/internal/infrastructure/grpc/proto"
	"github.com/Goalt/hostmonitor/internal/mocks/interactor"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestHandler_GetStatistics(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		input := entity.Statistics{
			UpdatedAt: time.Unix(957643200, 0),
		}

		statisticsInteractorI := mock_interactor.NewMockStatisticsI(ctrl)
		statisticsInteractorI.EXPECT().GetLast().Return(input)

		h := NewHandler(statisticsInteractorI)

		expected := apipb.StatisticsResponse{
			Loadavg: &apipb.LoadAvg{},
			Ram: &apipb.Ram{},
			Storage: &apipb.Storage{},
			Uptime: &apipb.Uptime{},
			UpdatedAt: 957643200,
		}

		got, err := h.GetStatistics(context.Background(), &emptypb.Empty{})

		require.NoError(t, err)
		assert.Equal(t, &expected, got)
	})
}
