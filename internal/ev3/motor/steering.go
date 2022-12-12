package motor

import (
	"errors"
	"fmt"
	"github.com/ev3go/ev3dev/motorutil"
	"log"
	"time"
)

var (
	steeringUnit       *motorutil.Steering
	defaultWaitTimeout = time.Second * 20
)

func ResetSteeringUnit() {
	steeringUnit = nil
}

func GetSteeringUnit(left, right string) (*motorutil.Steering, error) {
	if steeringUnit != nil {
		return steeringUnit, nil
	}

	leftM, leftOk := TachoMotors[left]
	rightM, rightOk := TachoMotors[right]
	if !leftOk || !rightOk {
		return nil, fmt.Errorf("one or more motors nod found. left port found: %v, right port found %v", leftOk, rightOk)
	}

	steeringUnit = &motorutil.Steering{
		Left:    leftM,
		Right:   rightM,
		Timeout: time.Second * 2,
	}

	return steeringUnit, nil
}

func SteerCounts(left, right string, speed, turn, counts int) error {
	steeringUnit, _ = GetSteeringUnit(left, right)

	steeringUnit.SteerCounts(speed, turn, counts)
	steeringUnit.Timeout = defaultWaitTimeout
	err := steeringUnit.Wait()
	if err != nil {
		log.Printf("ERROR - %v", err)
		return errors.New("something went wrong during last steering action")
	}

	return nil
}

func SteerDuration(left, right string, speed, turn int, duration time.Duration) error {
	steeringUnit, _ = GetSteeringUnit(left, right)

	steeringUnit.SteerDuration(speed, turn, duration)
	steeringUnit.Timeout = defaultWaitTimeout + duration
	err := steeringUnit.Wait()
	if err != nil {
		log.Printf("ERROR - %v", err)
		return errors.New("something went wrong during last steering action")
	}

	return nil
}
