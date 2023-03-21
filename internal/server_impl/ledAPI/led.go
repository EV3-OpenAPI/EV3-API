package ledAPI

import (
	"EV3-API/internal/ev3/led"
	"EV3-API/internal/gen/openapi"
	"context"
	"net/http"
)

type ApiService struct {
}

// NewLedApiService creates a default api service
func NewLedApiService() openapi.LedApiServicer {
	return &ApiService{}
}

func (a ApiService) LedFlashPost(_ context.Context, leds []openapi.Led) (openapi.ImplResponse, error) {
	// var errs []error = []

	for _, l := range leds {
		if l.Color != "" {
			_ = led.FlashString(l.Side, l.Color)
		} else {
			_ = led.Flash(l.Side, int(l.Red), int(l.Green))
		}
	}

	return openapi.Response(http.StatusOK, nil), nil
}

func (a ApiService) LedOffPost(_ context.Context) (openapi.ImplResponse, error) {
	_ = led.SetColorValues("left", 0, 0)
	_ = led.SetColorValues("right", 0, 0)

	return openapi.Response(http.StatusOK, nil), nil
}
