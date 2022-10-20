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

func (s *PowerApiService) PowerGet(_ context.Context) (openapi.ImplResponse, error) {
	p := ev3dev.PowerSupply("lego-ev3-battery")
	var internalErrors []string

	resp := openapi.PowerInfo{
		Voltage:    GetFloat32(p.Voltage, &internalErrors),
		Current:    GetFloat32(p.Current, &internalErrors),
		VoltageMax: GetFloat32(p.VoltageMax, &internalErrors),
		VoltageMin: GetFloat32(p.VoltageMax, &internalErrors),
		Technology: GetString(p.Technology, &internalErrors),
		Type:       GetString(p.Type, &internalErrors),
		UEvent:     GetStringMap(p.Uevent, &internalErrors),
	}

	if len(internalErrors) > 0 {
		log.Printf("ERROR - %v", internalErrors)
		return openapi.Response(http.StatusInternalServerError, nil), errors.New(strings.Join(internalErrors, ", "))
	}

	return openapi.Response(http.StatusOK, resp), nil
}
