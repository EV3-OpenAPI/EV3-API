package server_impl

import (
	ev3motor "EV3-API/internal/ev3/motor"
	"EV3-API/internal/gen/openapi"
	"context"
	"errors"
	"fmt"
	"github.com/ev3go/ev3dev"
	"log"
	"net/http"
	"strings"
)

// MotorApiService is a service that implements the logic for the DefaultApiServicer
// This service should implement the business logic for every endpoint for the DefaultApi API.
// Include any external packages or services that will be required by this service.
type MotorApiService struct {
}

// NewMotorApiService creates a default api service
func NewMotorApiService() openapi.MotorApiServicer {
	return &MotorApiService{}
}

func GetTachoMotor(mType string, port string) (*ev3dev.TachoMotor, error) {
	return ev3dev.TachoMotorFor(fmt.Sprintf("ev3-ports:out%s", port), fmt.Sprintf("lego-ev3-%s-ev3motor", mType))
}

func (s *MotorApiService) MotorTachoTypeCommandPost(_ context.Context, _ string, _ openapi.MotorTachoTypeCommandPostRequest) (openapi.ImplResponse, error) {
	//TODO implement me
	return openapi.Response(http.StatusNotImplemented, nil), nil
}

func (s *MotorApiService) MotorTachoTypePortGet(_ context.Context, mType string, port string) (openapi.ImplResponse, error) {
	var internalErrors []string

	m, exists := ev3motor.TachoMotors[port]
	if !exists {
		internalErrors = append(internalErrors, fmt.Sprintf("Could not get specified ev3motor port: %s", port))
	}

	pol, err := m.Polarity()
	if err != nil {
		internalErrors = append(internalErrors, fmt.Sprintf("Could not get polarity: %v", err))
	}

	state, err := m.State()
	if err != nil {
		internalErrors = append(internalErrors, fmt.Sprintf("Could not get state: %v", err))
	}

	resp := openapi.TachoMotorInfo{
		Commnds:                 m.Commands(),
		StopActions:             m.StopActions(),
		CountPerRot:             int32(m.CountPerRot()),
		DutyCycle:               GetInt32(m.DutyCycle, &internalErrors),
		DutyCycleSetpoint:       GetInt32(m.DutyCycleSetpoint, &internalErrors),
		Polarity:                string(pol),
		Position:                GetInt32(m.Position, &internalErrors),
		HoldPIDKd:               GetInt32(m.HoldPIDKd, &internalErrors),
		HoldPIDKi:               GetInt32(m.HoldPIDKi, &internalErrors),
		HoldPIDKp:               GetInt32(m.HoldPIDKp, &internalErrors),
		MaxSpeed:                int32(m.MaxSpeed()),
		PositionSetpoint:        GetInt32(m.PositionSetpoint, &internalErrors),
		CurrentSpeed:            GetInt32(m.Speed, &internalErrors),
		CurrentSpeedSetpoint:    GetInt32(m.SpeedSetpoint, &internalErrors),
		CurrentRampUpSetpoint:   GetDurationAsInt32(m.RampUpSetpoint, &internalErrors),
		CurrentRampDownSetpoint: GetDurationAsInt32(m.RampDownSetpoint, &internalErrors),
		SpeedPIDKd:              GetInt32(m.SpeedPIDKd, &internalErrors),
		SpeedPIDKi:              GetInt32(m.SpeedPIDKi, &internalErrors),
		SpeedPIDKp:              GetInt32(m.SpeedPIDKp, &internalErrors),
		State:                   int32(state),
		TimeSetpoint:            GetDurationAsInt32(m.TimeSetpoint, &internalErrors),
	}

	if len(internalErrors) > 0 {
		log.Printf("ERROR - %v", internalErrors)
		return openapi.Response(http.StatusInternalServerError, nil), errors.New(strings.Join(internalErrors, ", "))
	}

	return openapi.Response(http.StatusOK, resp), nil
}

func (s *MotorApiService) MotorTachoTypeMaxSpeedPost(_ context.Context, mType string, request openapi.MotorTachoTypeMaxSpeedPostRequest) (openapi.ImplResponse, error) {
	var internalErrors []string

	for _, port := range request.Ports {
		m, err := GetTachoMotor(mType, port)
		if err != nil {
			internalErrors = append(internalErrors, fmt.Sprintf("Could not get specified ev3motor: %v", err))
		}

		m.SetSpeedSetpoint(m.MaxSpeed()).Command("run-forever")
	}

	if len(internalErrors) > 0 {
		log.Printf("ERROR - %v", internalErrors)
		return openapi.Response(http.StatusInternalServerError, nil), errors.New(strings.Join(internalErrors, ", "))
	}

	return openapi.Response(http.StatusOK, nil), nil
}

func (s *MotorApiService) MotorTachoTypeSpeedSetpointPost(_ context.Context, _ string, _ openapi.MotorTachoTypeSpeedSetpointPostRequest) (openapi.ImplResponse, error) {
	//TODO implement me
	return openapi.Response(http.StatusNotImplemented, nil), nil
}

func (s *MotorApiService) MotorTachoPost(_ context.Context, request openapi.MotorRequest) (openapi.ImplResponse, error) {
	for _, m := range request.Motors {
		if motor := ev3motor.TachoMotors[m.Port]; motor.Driver()[9] == m.Size[0] {
			motor.SetSpeedSetpoint(int(request.Speed)).Command(request.Command)
		}
	}

	return openapi.Response(http.StatusOK, nil), nil
}

func (s *MotorApiService) MotorStopAllPost(_ context.Context) (openapi.ImplResponse, error) {
	ev3motor.StopAll()
	return openapi.Response(http.StatusOK, nil), nil
}
