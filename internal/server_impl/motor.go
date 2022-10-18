package server_impl

import (
	"EV3-API/internal/openapi"
	"context"
	"github.com/ev3go/ev3dev/motorutil"
	"net/http"
)

func (s *ApiService) MotorStopallPost(ctx context.Context) (openapi.ImplResponse, error) {
	err := motorutil.ResetAll()
	if err != nil {
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	return openapi.Response(http.StatusOK, nil), nil
}
