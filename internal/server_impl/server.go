package server_impl

import "EV3-API/internal/openapi"

// ApiService is a service that implements the logic for the DefaultApiServicer
// This service should implement the business logic for every endpoint for the DefaultApi API.
// Include any external packages or services that will be required by this service.
type ApiService struct {
}

// NewApiService creates a default api service
func NewApiService() openapi.DefaultApiServicer {
	return &ApiService{}
}
