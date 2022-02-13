package server

import (
	"context"
	"testing"

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

		statisticsInteractorI := mock_interactor.NewMockStatisticsI(ctrl)
		statisticsInteractorI.EXPECT().GetLast().Return(entity.Statistics{})

		h := NewHandler(statisticsInteractorI)

		expected := apipb.StatisticsResponse{}

		got, err := h.GetStatistics(context.Background(), &emptypb.Empty{})

		require.NoError(t, err)
		assert.Equal(t, &expected, got)
	})
}
