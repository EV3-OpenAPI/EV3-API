package main

import (
	"EV3-API/internal/ev3/motor"
	"EV3-API/internal/ev3/sound"
	"EV3-API/internal/gen/openapi"
	"EV3-API/internal/server_impl"
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Println("Started")

	_ = sound.Init()
	defer sound.Close()

	_ = motor.Init()

	MotorApiService := server_impl.NewMotorApiService()
	MotorApiController := openapi.NewMotorApiController(MotorApiService)

	PowerApiService := server_impl.NewPowerApiService()
	PowerApiController := openapi.NewPowerApiController(PowerApiService)

	SoundApiService := server_impl.NewSoundApiService()
	SoundApiController := openapi.NewSoundApiController(SoundApiService)

	router := openapi.NewRouter(MotorApiController, PowerApiController, SoundApiController)

	port := 8080

	log.Printf("INFO - Starting server on port %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
