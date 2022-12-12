package led

import (
	ev32 "EV3-API/internal/ev3"
	"fmt"
	"github.com/ev3go/ev3"
	"github.com/ev3go/ev3dev"
	"log"
)

const (
	minBrightness = 0
	maxBrightness = 255
)

var leds = [4]*ev3dev.LED{ev3.GreenLeft, ev3.GreenRight, ev3.RedLeft, ev3.RedRight}
var colors = map[string]colorRG{
	"red":         {255, 0},
	"orange":      {255, 128},
	"yellow":      {255, 255},
	"lime":        {128, 255},
	"green":       {0, 255},
	"dark_red":    {153, 0},
	"dar_orange":  {153, 76},
	"dark_yellow": {153, 153},
	"dark_lime":   {76, 153},
	"dark_green":  {0, 153},
}

type colorRG struct {
	Red   int
	Green int
}

func Init() error {
	log.Printf("INFO - Initializing LEDs")
	log.Printf("INFO - Checking LEDs")

	for _, led := range leds {
		if bright, err := led.MaxBrightness(); err != nil {
			log.Printf("ERROR - failed to read %v led, err: %v", led, err)
		} else if bright != maxBrightness {
			log.Printf("ERROR - LED %v brightness not %d but %d", led, maxBrightness, bright)
		}
	}

	return nil
}

func SetColorString(side string, color string, onDuration int32) error {
	if colRG, ok := colors[color]; !ok {
		return fmt.Errorf("color not found")
	} else {
		return SetColorValues(side, colRG.Red, colRG.Green, onDuration)
	}
}

func SetColorValues(side string, red, green int, onDuration int32) error {
	if !betweenMinMax(red) || !betweenMinMax(green) {
		return fmt.Errorf("red %d or green %d value not between %d and %d", red, green, minBrightness, maxBrightness)
	}

	switch side {
	case "left":
		ev3.RedLeft.SetDelayOff(ev32.DurationMs(onDuration))
		ev3.GreenLeft.SetDelayOff(ev32.DurationMs(onDuration))
		ev3.RedLeft.SetBrightness(red)
		ev3.GreenLeft.SetBrightness(green)
	case "right":
		ev3.RedRight.SetDelayOff(ev32.DurationMs(onDuration))
		ev3.GreenRight.SetDelayOff(ev32.DurationMs(onDuration))
		ev3.RedRight.SetBrightness(red)
		ev3.GreenRight.SetBrightness(green)
	default:
		return fmt.Errorf("side %s not found, allowed are 'left' and 'right'", side)
	}

	return nil
}

func betweenMinMax(val int) bool {
	return minBrightness <= val && val <= maxBrightness
}
