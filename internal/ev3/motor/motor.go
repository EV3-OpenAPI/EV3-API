package motor

import (
	"fmt"
	"github.com/ev3go/ev3dev"
	"log"
	"os"
	"regexp"
)

var (
	TachoMotors = make(map[string]*ev3dev.TachoMotor)
	ServoMotors = make(map[string]*ev3dev.ServoMotor)
	DCMotors    = make(map[string]*ev3dev.DCMotor)
)

func Init() (err error) {
	log.Printf("INFO - Initializing motors")

	re := regexp.MustCompile("\\/(ev3-ports:out([A-Z])):(lego-ev3-([sml])-motor)\\/")

	files, err := os.ReadDir(ev3dev.TachoMotorPath)
	if err != nil {
		return
	}

	for _, file := range files {
		info, err := os.Readlink(ev3dev.TachoMotorPath + "/" + file.Name())
		if err != nil {
			fmt.Println(err)
		}

		matches := re.FindStringSubmatch(info)
		TachoMotors[matches[2]], _ = ev3dev.TachoMotorFor(matches[1], matches[3])
		log.Printf("INFO - Loaded tacho motor - port %s, driver %s", matches[2], matches[4])
	}

	return
}

func StopAll() {
	for _, motor := range TachoMotors {
		motor.Command("reset")
	}
	for _, motor := range ServoMotors {
		motor.Command("float")
	}
	for _, motor := range DCMotors {
		motor.Command("stop")
	}
}
