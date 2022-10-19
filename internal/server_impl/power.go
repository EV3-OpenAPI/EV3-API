package server_impl

import (
	"EV3-API/internal/openapi"
	"context"
	"errors"
	"fmt"
	"github.com/ev3go/ev3dev"
	"log"
	"net/http"
	"strings"
)

func (s *ApiService) PowerGet(ctx context.Context) (openapi.ImplResponse, error) {
	p := ev3dev.PowerSupply("lego-ev3-battery")
	var internal_errors []string

	v, err := p.Voltage()
	if err != nil {
		internal_errors = append(internal_errors, fmt.Sprintf("could not read voltage: %v", err))
	}

	i, err := p.Current()
	if err != nil {
		internal_errors = append(internal_errors, fmt.Sprintf("could not read current: %v", err))
	}

	vMax, err := p.VoltageMax()
	if err != nil {
		internal_errors = append(internal_errors, fmt.Sprintf("could not read max design voltage: %v", err))
	}

	vMin, err := p.VoltageMin()
	if err != nil {
		internal_errors = append(internal_errors, fmt.Sprintf("could not read min design voltage: %v", err))
	}

	tech, err := p.Technology()
	if err != nil {
		internal_errors = append(internal_errors, fmt.Sprintf("could not read technology: %v", err))
	}

	t, err := p.Type()
	if err != nil {
		internal_errors = append(internal_errors, fmt.Sprintf("could not read type: %v", err))
	}

	uevt, err := p.Uevent()
	if err != nil {
		internal_errors = append(internal_errors, fmt.Sprintf("could not read min design voltage: %v", err))
	}

	if len(internal_errors) > 0 {
		log.Printf("ERROR - %v", internal_errors)
		return openapi.Response(http.StatusInternalServerError, nil), errors.New(strings.Join(internal_errors, ", "))
	}

	resp := openapi.PowerInfo{
		Voltage:    float32(v),
		Current:    float32(i),
		VoltageMax: float32(vMax),
		VoltageMin: float32(vMin),
		Technology: tech,
		Type:       t,
		UEvent:     uevt,
	}

	return openapi.Response(http.StatusOK, resp), nil
}
