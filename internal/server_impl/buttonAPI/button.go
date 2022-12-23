package buttonAPI

import (
	"EV3-API/internal/ev3/button"
	"EV3-API/internal/gen/openapi"
	"context"
	"net/http"
	"time"
)

type ApiService struct {
}

// NewButtonApiService creates a default api service
func NewButtonApiService() openapi.ButtonApiServicer {
	return &ApiService{}
}

func (a ApiService) ButtonPressedGet(_ context.Context) (openapi.ImplResponse, error) {
	var resp []string

	evt := button.GetLastButtonEvent(false)
	if evt != nil && time.Now().Sub(evt.TimeStamp) < time.Second*3 {
		resp = append(resp, "button")
	}

	return openapi.Response(http.StatusOK, resp), nil
}
