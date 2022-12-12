package status

import (
	"EV3-API/internal/ev3/lcd"
	"EV3-API/internal/ev3/sensor"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

var (
	ifIdx   = -1
	updates = []update{
		{"IP", getIP},
		{"US", getUS},
		{"Gyro", getGyro},
	}
)

type update struct {
	Name string
	Fn   func() string
}

// Start the update loop with the given interval between updates.
// The update loop runs in a separate go routine (almost like a thread).
// The duration between updates can go to 0 if the update takes longer than the interval
func Start(interval time.Duration) {
	findWlanInterface()

	go startLoop(interval)
}

func findWlanInterface() {
	interfaces, err := net.Interfaces()
	if err != nil {
		log.Printf("ERROR - No WLAN interface found. Info will not be avaiable")
	}

	for _, intf := range interfaces {
		if strings.HasPrefix(intf.Name, "wlx") {
			ifIdx = intf.Index
		}
	}
}

func startLoop(interval time.Duration) {
	for true {
		start := time.Now()
		displayStatus()
		duration := time.Now().Sub(start)
		log.Printf("DEBUG - displayStatus: duration %v", duration)
		time.Sleep(interval - duration)
	}
}

func displayStatus() {
	start := time.Now()

	dLines := make([]string, len(updates))
	for i, u := range updates {
		dLines[i] = fmt.Sprintf("%s: %s", u.Name, u.Fn())
	}

	err := lcd.Write(strings.Join(dLines, "\n"))
	if err != nil {
		log.Println(err)
	}

	log.Printf("INFO - time to draw: %v", time.Now().Sub(start))
}

func getIP() string {
	if ifIdx == -1 {
		findWlanInterface()
	}

	if intf, err := net.InterfaceByIndex(ifIdx); err == nil {
		if addrs, err := intf.Addrs(); err == nil {
			return addrs[0].String()
		}
	}

	return "not connected"
}

func getUS() string {
	us, err := sensor.GetSensor("us")
	if err != nil {
		return "no sensor"
	}

	value, err := us.Value(0)
	if err != nil {
		return "err reading value"
	}

	return value
}

func getGyro() string {
	us, err := sensor.GetSensor("gyro")
	if err != nil {
		return "no sensor"
	}

	value, err := us.Value(0)
	if err != nil {
		return "err reading value"
	}

	return value
}