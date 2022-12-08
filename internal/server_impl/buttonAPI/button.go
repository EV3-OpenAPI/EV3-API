package buttonAPI

import (
	"EV3-API/internal/gen/openapi"
	"context"
	"github.com/ev3go/ev3dev"
)

type ApiService struct {
}

// NewButtonApiService creates a default api service
func NewButtonApiService() openapi.ButtonApiServicer {
	return &ApiService{}
}

func (a ApiService) ButtonPressedGet(ctx context.Context) (openapi.ImplResponse, error) {
	ev3dev.ButtonEvent{}
	//TODO implement me
	panic("implement me")
}
