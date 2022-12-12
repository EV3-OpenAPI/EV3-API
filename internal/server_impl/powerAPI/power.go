package powerAPI

import (
	"EV3-API/internal/gen/openapi"
	"EV3-API/internal/server_impl"
	"context"
	"errors"
	"github.com/ev3go/ev3dev"
	"log"
	"net/http"
	"strings"
)

// ApiService is a service that implements the logic for the DefaultApiServicer
// This service should implement the business logic for every endpoint for the DefaultApi API.
// Include any external packages or services that will be required by this service.
type ApiService struct {
}

// NewPowerApiService creates a default api service
func NewPowerApiService() openapi.PowerApiServicer {
	return &ApiService{}
}

func (s *ApiService) PowerGet(_ context.Context) (openapi.ImplResponse, error) {
	p := ev3dev.PowerSupply("lego-ev3-battery")
	var internalErrors []string

	resp := openapi.PowerInfo{
		Voltage:    server_impl.GetFloat32(p.Voltage, &internalErrors),
		Current:    server_impl.GetFloat32(p.Current, &internalErrors),
		VoltageMax: server_impl.GetFloat32(p.VoltageMax, &internalErrors),
		VoltageMin: server_impl.GetFloat32(p.VoltageMax, &internalErrors),
		Technology: server_impl.GetString(p.Technology, &internalErrors),
		Type:       server_impl.GetString(p.Type, &internalErrors),
		UEvent:     server_impl.GetStringMap(p.Uevent, &internalErrors),
	}

	if len(internalErrors) > 0 {
		log.Printf("ERROR - %v", internalErrors)
		return openapi.Response(http.StatusInternalServerError, nil), errors.New(strings.Join(internalErrors, ", "))
	}

	return openapi.Response(http.StatusOK, resp), nil
}
