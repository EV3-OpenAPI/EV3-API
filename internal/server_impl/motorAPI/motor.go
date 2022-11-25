package motorAPI

import (
	"EV3-API/internal/ev3"
	ev3motor "EV3-API/internal/ev3/motor"
	"EV3-API/internal/gen/openapi"
	"EV3-API/internal/server_impl"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// ApiService is a service that implements the logic for the DefaultApiServicer
// This service should implement the business logic for every endpoint for the DefaultApi API.
// Include any external packages or services that will be required by this service.
type ApiService struct {
}

// NewMotorApiService creates a default api service
func NewMotorApiService() openapi.MotorApiServicer {
	return &ApiService{}
}

func (s *ApiService) MotorSteerResetPost(ctx context.Context) (openapi.ImplResponse, error) {
	ev3motor.ResetSteeringUnit()

	return openapi.Response(http.StatusNoContent, nil), nil
}

func (s *ApiService) MotorTachoTypePortGet(_ context.Context, mType string, port string) (openapi.ImplResponse, error) {
	internalErrors, resp := getTachoMotorInfo(port)

	if len(internalErrors) > 0 {
		log.Printf("ERROR - %v", internalErrors)
		return openapi.Response(http.StatusInternalServerError, nil), errors.New(strings.Join(internalErrors, ", "))
	}

	return openapi.Response(http.StatusOK, resp), nil
}

func (s *ApiService) MotorTachoPost(_ context.Context, request openapi.MotorRequest) (openapi.ImplResponse, error) {
	for _, m := range request.Motors {
		if motor := ev3motor.TachoMotors[m.Port]; motor.Driver()[9] == m.Size[0] {
			motor.SetSpeedSetpoint(int(request.Speed)).Command(request.Command)
		}
	}

	return openapi.Response(http.StatusOK, nil), nil
}

func (s *ApiService) MotorTachoGet(_ context.Context) (openapi.ImplResponse, error) {
	var tachoMotorInfoList []openapi.TachoMotor

	for key, _ := range ev3motor.TachoMotors {
		internalErrors, resp := getTachoMotorInfo(key)

		if len(internalErrors) > 0 {
			log.Printf("ERROR - %v", internalErrors)
			return openapi.Response(http.StatusInternalServerError, nil), errors.New(strings.Join(internalErrors, ", "))
		}

		tachoMotorInfoList = append(tachoMotorInfoList, resp)
	}

	return openapi.Response(http.StatusOK, tachoMotorInfoList), nil
}

func (s *ApiService) MotorSteerCountsPost(_ context.Context, request openapi.MotorSteerCountsPostRequest) (openapi.ImplResponse, error) {
	err := ev3motor.SteerCounts(request.SteeringUnit.Left.Port, request.SteeringUnit.Right.Port, int(request.Speed), int(request.Turn), int(request.Counts))
	if err != nil {
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	return openapi.Response(http.StatusOK, nil), nil
}

func (s *ApiService) MotorSteerDurationPost(_ context.Context, request openapi.MotorSteerDurationPostRequest) (openapi.ImplResponse, error) {
	err := ev3motor.SteerDuration(request.SteeringUnit.Left.Port, request.SteeringUnit.Right.Port, int(request.Speed), int(request.Turn), ev3.DurationMs(request.DurationMs))
	if err != nil {
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	return openapi.Response(http.StatusOK, nil), nil
}

func (s *ApiService) MotorStopAllPost(_ context.Context) (openapi.ImplResponse, error) {
	ev3motor.StopAll()
	return openapi.Response(http.StatusOK, nil), nil
}

// helper functions

func getTachoMotorInfo(port string) ([]string, openapi.TachoMotor) {
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

	resp := openapi.TachoMotor{
		Commnds:                 m.Commands(),
		StopActions:             m.StopActions(),
		CountPerRot:             int32(m.CountPerRot()),
		DutyCycle:               server_impl.GetInt32(m.DutyCycle, &internalErrors),
		DutyCycleSetpoint:       server_impl.GetInt32(m.DutyCycleSetpoint, &internalErrors),
		Polarity:                string(pol),
		Position:                server_impl.GetInt32(m.Position, &internalErrors),
		HoldPIDKd:               server_impl.GetInt32(m.HoldPIDKd, &internalErrors),
		HoldPIDKi:               server_impl.GetInt32(m.HoldPIDKi, &internalErrors),
		HoldPIDKp:               server_impl.GetInt32(m.HoldPIDKp, &internalErrors),
		MaxSpeed:                int32(m.MaxSpeed()),
		PositionSetpoint:        server_impl.GetInt32(m.PositionSetpoint, &internalErrors),
		CurrentSpeed:            server_impl.GetInt32(m.Speed, &internalErrors),
		CurrentSpeedSetpoint:    server_impl.GetInt32(m.SpeedSetpoint, &internalErrors),
		CurrentRampUpSetpoint:   server_impl.GetDurationAsInt32(m.RampUpSetpoint, &internalErrors),
		CurrentRampDownSetpoint: server_impl.GetDurationAsInt32(m.RampDownSetpoint, &internalErrors),
		SpeedPIDKd:              server_impl.GetInt32(m.SpeedPIDKd, &internalErrors),
		SpeedPIDKi:              server_impl.GetInt32(m.SpeedPIDKi, &internalErrors),
		SpeedPIDKp:              server_impl.GetInt32(m.SpeedPIDKp, &internalErrors),
		State:                   int32(state),
		TimeSetpoint:            server_impl.GetDurationAsInt32(m.TimeSetpoint, &internalErrors),
	}
	return internalErrors, resp
}
