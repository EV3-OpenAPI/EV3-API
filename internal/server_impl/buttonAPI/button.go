package buttonAPI

import (
	"EV3-API/internal/gen/openapi"
	"context"
	"net/http"
)

type ApiService struct {
}

// NewButtonApiService creates a default api service
func NewButtonApiService() openapi.ButtonApiServicer {
	return &ApiService{}
}

func (a ApiService) ButtonPressedGet(_ context.Context) (openapi.ImplResponse, error) {
	// TODO: implement
	return openapi.Response(http.StatusNotImplemented, nil), nil
}
