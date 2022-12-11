package main

import (
	"EV3-API/internal/ev3/lcd"
	"EV3-API/internal/ev3/motor"
	"EV3-API/internal/ev3/sensor"
	"EV3-API/internal/ev3/sound"
	"EV3-API/internal/ev3/status"
	"EV3-API/internal/gen/openapi"
	"EV3-API/internal/server_impl/motorAPI"
	"EV3-API/internal/server_impl/powerAPI"
	"EV3-API/internal/server_impl/sensorAPI"
	"EV3-API/internal/server_impl/soundAPI"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	log.Println("Started")

	_ = sound.Init()
	defer sound.Close()
	_ = motor.Init()
	_ = sensor.Init()
	_ = lcd.Init()
	status.Start(time.Second * 2)

	MotorApiService := motorAPI.NewMotorApiService()
	MotorApiController := openapi.NewMotorApiController(MotorApiService)

	PowerApiService := powerAPI.NewPowerApiService()
	PowerApiController := openapi.NewPowerApiController(PowerApiService)

	SoundApiService := soundAPI.NewSoundApiService()
	SoundApiController := openapi.NewSoundApiController(SoundApiService)

	SensorApiService := sensorAPI.NewSensorApiService()
	SensorApiController := openapi.NewSensorApiController(SensorApiService)

	router := openapi.NewRouter(MotorApiController, PowerApiController, SoundApiController, SensorApiController)

	port := 8080

	log.Printf("INFO - Starting server on port %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
