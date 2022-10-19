package server_impl

import (
	"EV3-API/internal/openapi"
	"context"
	"errors"
	"fmt"
	"github.com/ev3go/ev3dev"
	"github.com/ev3go/ev3dev/motorutil"
	"log"
	"net/http"
	"strings"
	"time"
)

// MotorApiService is a service that implements the logic for the DefaultApiServicer
// This service should implement the business logic for every endpoint for the DefaultApi API.
// Include any external packages or services that will be required by this service.
type MotorApiService struct {
}

// NewPowerApiService creates a default api service
func NewMotorApiService() openapi.MotorApiServicer {
	return &MotorApiService{}
}

func GetTachoMotor(mType string, port string) (*ev3dev.TachoMotor, error) {
	return ev3dev.TachoMotorFor(fmt.Sprintf("ev3-ports:out%s", port), fmt.Sprintf("lego-ev3-%s-motor", mType))
}

func (s *MotorApiService) MotorTachoTypeCommandPost(ctx context.Context, mType string, request openapi.MotorTachoTypeCommandPostRequest) (openapi.ImplResponse, error) {
	//TODO implement me
	return openapi.Response(http.StatusNotImplemented, nil), nil
}

func (s *MotorApiService) MotorTachoTypePortGet(ctx context.Context, mType string, port string) (openapi.ImplResponse, error) {
	var internal_errors []string

	m, err := GetTachoMotor(mType, port)
	if err != nil {
		internal_errors = append(internal_errors, fmt.Sprintf("Could not get specified motor: %v", err))
	}

	pol, err := m.Polarity()
	if err != nil {
		internal_errors = append(internal_errors, fmt.Sprintf("Could not get polarity: %v", err))
	}

	state, err := m.State()
	if err != nil {
		internal_errors = append(internal_errors, fmt.Sprintf("Could not get state: %v", err))
	}

	resp := openapi.TachoMotorInfo{
		Commnds:                 m.Commands(),
		StopActions:             m.StopActions(),
		CountPerRot:             int32(m.CountPerRot()),
		DutyCycle:               GetInt32(m.DutyCycle, &internal_errors),
		DutyCycleSetpoint:       GetInt32(m.DutyCycleSetpoint, &internal_errors),
		Polarity:                string(pol),
		Position:                GetInt32(m.Position, &internal_errors),
		HoldPIDKd:               GetInt32(m.HoldPIDKd, &internal_errors),
		HoldPIDKi:               GetInt32(m.HoldPIDKi, &internal_errors),
		HoldPIDKp:               GetInt32(m.HoldPIDKp, &internal_errors),
		MaxSpeed:                int32(m.MaxSpeed()),
		PositionSetpoint:        GetInt32(m.PositionSetpoint, &internal_errors),
		CurrentSpeed:            GetInt32(m.Speed, &internal_errors),
		CurrentSpeedSetpoint:    GetInt32(m.SpeedSetpoint, &internal_errors),
		CurrentRampUpSetpoint:   GetDurationAsInt32(m.RampUpSetpoint, &internal_errors),
		CurrentRampDownSetpoint: GetDurationAsInt32(m.RampDownSetpoint, &internal_errors),
		SpeedPIDKd:              GetInt32(m.SpeedPIDKd, &internal_errors),
		SpeedPIDKi:              GetInt32(m.SpeedPIDKi, &internal_errors),
		SpeedPIDKp:              GetInt32(m.SpeedPIDKp, &internal_errors),
		State:                   int32(state),
		TimeSetpoint:            GetDurationAsInt32(m.TimeSetpoint, &internal_errors),
	}

	if len(internal_errors) > 0 {
		log.Printf("ERROR - %v", internal_errors)
		return openapi.Response(http.StatusInternalServerError, nil), errors.New(strings.Join(internal_errors, ", "))
	}

	return openapi.Response(http.StatusOK, resp), nil
}

func (s *MotorApiService) MotorTachoTypeMaxSpeedPost(ctx context.Context, mType string, request openapi.MotorTachoTypeMaxSpeedPostRequest) (openapi.ImplResponse, error) {
	var internal_errors []string

	for _, port := range request.Ports {
		m, err := GetTachoMotor(mType, port)
		if err != nil {
			internal_errors = append(internal_errors, fmt.Sprintf("Could not get specified motor: %v", err))
		}

		m.SetSpeedSetpoint(m.MaxSpeed()).Command("run-forever")
	}

	if len(internal_errors) > 0 {
		log.Printf("ERROR - %v", internal_errors)
		return openapi.Response(http.StatusInternalServerError, nil), errors.New(strings.Join(internal_errors, ", "))
	}

	return openapi.Response(http.StatusOK, nil), nil
}

func (s *MotorApiService) MotorTachoTypeSpeedSetpointPost(ctx context.Context, mType string, request openapi.MotorTachoTypeSpeedSetpointPostRequest) (openapi.ImplResponse, error) {
	//TODO implement me
	return openapi.Response(http.StatusNotImplemented, nil), nil
}

func (s *MotorApiService) MotorStopallPost(ctx context.Context) (openapi.ImplResponse, error) {
	start := time.Now()
	err := motorutil.ResetAll()
	if err != nil {
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	log.Printf("INFO - execution time: %s", time.Since(start))
	return openapi.Response(http.StatusOK, nil), nil
}
