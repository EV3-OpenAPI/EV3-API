package server_impl

import (
	"EV3-API/internal/openapi"
	"context"
	"errors"
	"github.com/ev3go/ev3dev"
	"log"
	"net/http"
	"strings"
)

// PowerApiService is a service that implements the logic for the DefaultApiServicer
// This service should implement the business logic for every endpoint for the DefaultApi API.
// Include any external packages or services that will be required by this service.
type PowerApiService struct {
}

// NewPowerApiService creates a default api service
func NewPowerApiService() openapi.PowerApiServicer {
	return &PowerApiService{}
}

func (s *PowerApiService) PowerGet(ctx context.Context) (openapi.ImplResponse, error) {
	p := ev3dev.PowerSupply("lego-ev3-battery")
	var internal_errors []string

	resp := openapi.PowerInfo{
		Voltage:    GetFloat32(p.Voltage, &internal_errors),
		Current:    GetFloat32(p.Current, &internal_errors),
		VoltageMax: GetFloat32(p.VoltageMax, &internal_errors),
		VoltageMin: GetFloat32(p.VoltageMax, &internal_errors),
		Technology: GetString(p.Technology, &internal_errors),
		Type:       GetString(p.Type, &internal_errors),
		UEvent:     GetStringMap(p.Uevent, &internal_errors),
	}

	if len(internal_errors) > 0 {
		log.Printf("ERROR - %v", internal_errors)
		return openapi.Response(http.StatusInternalServerError, nil), errors.New(strings.Join(internal_errors, ", "))
	}

	return openapi.Response(http.StatusOK, resp), nil
}
