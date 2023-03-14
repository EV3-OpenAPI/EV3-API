package main

import (
	"EV3-API/internal/ev3"
	"EV3-API/internal/ev3/button"
	"EV3-API/internal/ev3/lcd"
	"EV3-API/internal/ev3/motor"
	"EV3-API/internal/ev3/sensor"
	"EV3-API/internal/ev3/sound"
	"EV3-API/internal/ev3/status"
	"EV3-API/internal/gen/openapi"
	"EV3-API/internal/server_impl/buttonAPI"
	"EV3-API/internal/server_impl/motorAPI"
	"EV3-API/internal/server_impl/powerAPI"
	"EV3-API/internal/server_impl/sensorAPI"
	"EV3-API/internal/server_impl/soundAPI"
	"EV3-API/internal/utils"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	log.Println("INFO - Starting")

	getHostname := flag.Bool("get-hostname", false, "only return hostname for this device")
	noMonitor := flag.Bool("no-monitor", false, "do not create a display overlay")
	noButton := flag.Bool("no-button", false, "do not create a button event listener loop")
	verify := flag.Bool("verify", false, "exit with status code 0, check if executable")
	update := flag.Bool("update", false, "check if new versions are available")
	port := flag.Int("port", 8080, "port to listen on")
	flag.Parse()

	if *verify {
		log.Printf("INFO - Verify mode, exiting...")
		os.Exit(0)
	}

	if *getHostname {
		fmt.Print(ev3.GetHostname())
		os.Exit(0)
	}

	if *update {
		utils.CheckForNewVersion()
	}

	initDevices(*noMonitor, *noButton)
	defer closeDevices()
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

	ButtonApiService := buttonAPI.NewButtonApiService()
	ButtonApiController := openapi.NewButtonApiController(ButtonApiService)

	router := openapi.NewRouter(MotorApiController, PowerApiController, SoundApiController, SensorApiController, ButtonApiController)

	_ = sound.Speak(fmt.Sprintf("%s at your service", ev3.GetHostname()))

	log.Printf("INFO - Starting server on port %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}

func initDevices(noMonitor, noButton bool) {
	_ = sound.Init()
	_ = motor.Init()
	_ = sensor.Init()
	_ = lcd.Init()

	if !noButton {
		_ = button.Init()
	}

	if !noMonitor {
		status.Start(time.Second * 2)
	}
}

func closeDevices() {
	_ = sound.Close()
}
