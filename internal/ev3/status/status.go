package status

import (
	"EV3-API/internal/ev3"
	"EV3-API/internal/ev3/button"
	"EV3-API/internal/ev3/lcd"
	"EV3-API/internal/ev3/sensor"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

var (
	updates = []update{
		{"IP", getIP},
		{"US", getUS},
		{"Gyro", getGyro},
	}

	lastDText string
	interval  time.Duration
)

type update struct {
	Name string
	Fn   func() string
}

// Start the update loop with the given interval between updates.
// The update loop runs in a separate go routine (almost like a thread).
// The duration between updates can go to 0 if the update takes longer than the interval
func Start(interv time.Duration) {
	interval = interv
	go startLoop()
}

func startLoop() {
	lcd.ShowSystemTTY(false)      // Hide system screen
	defer lcd.ShowSystemTTY(true) // Show system screen on exit

	for true {
		start := time.Now()
		displayStatus()

		if evt := button.GetLastButtonEvent(false); evt.TimeStamp < interval {
			break
		}

		duration := time.Now().Sub(start)
		log.Printf("DEBUG - displayStatus: duration %v", duration)
		time.Sleep(interval - duration)

	}
}

func displayStatus() {
	dLines := make([]string, len(updates))
	for i, u := range updates {
		dLines[i] = fmt.Sprintf("%s: %s", u.Name, u.Fn())
	}

	dText := strings.Join(dLines, "\n")
	if dText != lastDText {
		err := lcd.Write(dText)
		if err != nil {
			log.Println(err)
		}
		lastDText = dText
	} else {
		err := lcd.FastWrite(dText)
		if err != nil {
			log.Println(err)
		}
	}
}

func getIP() string {
	ifIdx := ev3.GetWlanInterfaceIndex()

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
