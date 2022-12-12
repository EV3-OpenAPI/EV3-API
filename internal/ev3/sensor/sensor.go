package sensor

import (
	"fmt"
	"github.com/ev3go/ev3dev"
	"log"
	"os"
	"strings"
)

var (
	Sensors = make(map[string]*ev3dev.Sensor)
	Drivers = map[string]string{
		"lego-ev3-us":   "us",
		"lego-ev3-gyro": "gyro",
	}
)

// Init initializes all connected ev3dev.Sensor and exposes them under Sensors.
// Other functions in this package depend on Init having been called first.
func Init() error {
	log.Printf("INFO - Initializing sensors")

	files, err := os.ReadDir(ev3dev.SensorPath)
	if err != nil {
		return err
	}

	for _, file := range files {
		driverNameBytes, err := os.ReadFile(fmt.Sprintf("%s/%s/driver_name", ev3dev.SensorPath, file.Name()))
		if err != nil {
			return err
		}

		addressBytes, err := os.ReadFile(fmt.Sprintf("%s/%s/address", ev3dev.SensorPath, file.Name()))
		if err != nil {
			return err
		}

		driverName := strings.TrimSpace(string(driverNameBytes))
		address := strings.TrimSpace(string(addressBytes))

		friendlyDriver := Drivers[driverName]
		Sensors[friendlyDriver], err = ev3dev.SensorFor(address, driverName)
		if err != nil {
			return err
		}

		log.Printf("INFO - Loaded %s sensor - port: %s, driver: %s", friendlyDriver, address[:len(address)-1], driverName)
	}

	return nil
}

func GetSensor(driver string) (*ev3dev.Sensor, error) {
	if sensor, ok := Sensors[driver]; ok {
		return sensor, nil
	}

	return nil, fmt.Errorf("no sensor with the driver %s found", driver)
}

// ReadTextValues returns a list of strings from the sensor with the given driver name.
// Error either if there is no sensor with the given driver or something during the read went wrong.
func ReadTextValues(driver string) ([]string, error) {
	if sensor, ok := Sensors[driver]; ok {
		return sensor.TextValues()
	}

	return nil, fmt.Errorf("no sensor with the driver %s found", driver)
}

// ReadNumValue returns a number from the sensor with the given driver name.
// Error if there is no sensor with the given driver.
func ReadNumValue(driver string) ([]string, error) {
	if sensor, err := GetSensor(driver); err != nil {
		return make([]string, 0), err
	} else {
		values := make([]string, sensor.NumValues())
		for i := 0; i < sensor.NumValues(); i++ {
			val, err := sensor.Value(i)
			if err == nil {
				values[i] = val
			}
		}
		return values, nil
	}
}
