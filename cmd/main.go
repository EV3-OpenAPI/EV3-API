package main

import (
	"EV3-API/internal/ev3"
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
	"EV3-API/internal/utils"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	log.Println("INFO - Starting")

	getHostname := flag.Bool("get-hostname", false, "only return hostname for this device")
	noMonitor := flag.Bool("no-monitor", false, "do not create a display overlay")
	update := flag.Bool("update", false, "check if new versions are available")
	port := flag.Int("port", 8080, "port to listen on")
	flag.Parse()

	if *getHostname {
		fmt.Print(ev3.GetHostname())
		return
	}

	if *update {
		utils.CheckForNewVersion()
	}

	initDevices(*noMonitor)
	startServer(*port)
}

func startServer(port int) {
	MotorApiService := motorAPI.NewMotorApiService()
	MotorApiController := openapi.NewMotorApiController(MotorApiService)

	PowerApiService := powerAPI.NewPowerApiService()
	PowerApiController := openapi.NewPowerApiController(PowerApiService)

	SoundApiService := soundAPI.NewSoundApiService()
	SoundApiController := openapi.NewSoundApiController(SoundApiService)

	SensorApiService := sensorAPI.NewSensorApiService()
	SensorApiController := openapi.NewSensorApiController(SensorApiService)

	router := openapi.NewRouter(MotorApiController, PowerApiController, SoundApiController, SensorApiController)

	log.Printf("INFO - Starting server on port %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}

func initDevices(noMonitor bool) {
	_ = sound.Init()
	defer sound.Close()
	_ = motor.Init()
	_ = sensor.Init()
	_ = lcd.Init()

	if !noMonitor {
		status.Start(time.Second * 2)
	}
}
