package sensorAPI

import (
	"EV3-API/internal/ev3"
	"EV3-API/internal/ev3/sensor"
	"EV3-API/internal/gen/openapi"
	"EV3-API/internal/server_impl"
	"context"
	"errors"
	"fmt"
	"golang.org/x/exp/slices"
	"net/http"
	"strings"
)

// ApiService is a service that implements the logic for the DefaultApiServicer
// This service should implement the business logic for every endpoint for the DefaultApi API.
// Include any external packages or services that will be required by this service.
type ApiService struct {
}

// NewSensorApiService creates a default api service
func NewSensorApiService() openapi.SensorApiServicer {
	return &ApiService{}
}

func (a ApiService) SensorGet(_ context.Context) (openapi.ImplResponse, error) {
	var sensorList []openapi.Sensor

	for sensorDriver, _ := range sensor.Sensors {
		s, err := getSensorInfo(sensorDriver)
		if len(err) > 0 {
			return openapi.Response(http.StatusInternalServerError, nil), errors.New(strings.Join(err, ", "))
		}

		sensorList = append(sensorList, s)
	}

	return openapi.Response(http.StatusOK, sensorList), nil
}

func (a ApiService) SensorTypeGet(_ context.Context, sensorDriver string) (openapi.ImplResponse, error) {
	resp, err := getSensorInfo(sensorDriver)
	if len(err) > 0 {
		return openapi.Response(http.StatusInternalServerError, nil), errors.New(strings.Join(err, ", "))
	}

	return openapi.Response(http.StatusOK, resp), nil
}

func (a ApiService) SensorTypePut(_ context.Context, sensorDriver string, oapiSensor openapi.Sensor) (openapi.ImplResponse, error) {
	s, err := sensor.GetSensor(sensorDriver)
	if err != nil {
		return openapi.Response(http.StatusNotFound, nil), err
	}

	if oapiSensor.Mode != "" {
		if !slices.Contains(s.Modes(), oapiSensor.Mode) {
			return openapi.Response(http.StatusBadRequest, nil), nil
		}
		s.SetMode(oapiSensor.Mode)
	}

	if oapiSensor.PollRateMs != 0 {
		s.SetPollRate(ev3.DurationMs(oapiSensor.PollRateMs))
	}

	return openapi.Response(http.StatusNoContent, nil), nil
}

func (a ApiService) SensorTypeNumValueGet(_ context.Context, sensorDriver string) (openapi.ImplResponse, error) {
	resp, err := sensor.ReadNumValue(sensorDriver)
	if err != nil { // TODO: check error type if actually not found or something else
		return openapi.Response(http.StatusNotFound, nil), err
	}

	return openapi.Response(http.StatusOK, resp), nil
}

func (a ApiService) SensorTypeTextValuesGet(_ context.Context, sensorDriver string) (openapi.ImplResponse, error) {
	resp, err := sensor.ReadTextValues(sensorDriver)
	if err != nil { // TODO: check error type if actually not found or something else
		return openapi.Response(http.StatusNotFound, nil), err
	}

	return openapi.Response(http.StatusOK, resp), nil
}

// helper functions
func getSensorInfo(sensorDriver string) (s openapi.Sensor, internalErrors []string) {
	ev3sensor, err := sensor.GetSensor(sensorDriver)
	if err != nil {
		internalErrors = append(internalErrors, fmt.Sprintf("Could not get specified ev3 sensor driver: %s", sensorDriver))
		return
	}

	s.Type = ev3sensor.Type()
	s.DriverName = ev3sensor.Driver()
	s.Port = ev3sensor.Path()
	s.Modes = ev3sensor.Modes()
	s.Commands = ev3sensor.Commands()
	s.Mode = server_impl.GetString(ev3sensor.Mode, &internalErrors)
	s.Decimals = int32(ev3sensor.Decimals())
	s.PollRateMs = server_impl.GetDurationAsInt32(ev3sensor.PollRate, &internalErrors)
	s.Units = ev3sensor.Units()

	return
}
