package main

import (
	"EV3-API/internal/openapi"
	"EV3-API/internal/server_impl"
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Println("Started")

	DefaultApiService := server_impl.NewApiService()
	DefaultApiController := openapi.NewDefaultApiController(DefaultApiService)

	router := openapi.NewRouter(DefaultApiController)

	port := 8080

	log.Printf("INFO - Starting server on port %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
