package server_impl

import (
	"fmt"
	"time"
)

func GetInt32(f func() (int, error), internal_errors *[]string) int32 {
	val, err := f()
	if err != nil {
		*internal_errors = append(*internal_errors, fmt.Sprintf("Could not get int: %v", err))
	}

	return int32(val)
}

func GetFloat32(f func() (float64, error), internal_errors *[]string) float32 {
	val, err := f()
	if err != nil {
		*internal_errors = append(*internal_errors, fmt.Sprintf("Could not get float: %v", err))
	}

	return float32(val)
}

func GetDurationAsInt32(f func() (time.Duration, error), internal_errors *[]string) int32 {
	val, err := f()
	if err != nil {
		*internal_errors = append(*internal_errors, fmt.Sprintf("Could not get float: %v", err))
	}

	return int32(val.Milliseconds())
}

func GetString(f func() (string, error), internal_errors *[]string) string {
	val, err := f()
	if err != nil {
		*internal_errors = append(*internal_errors, fmt.Sprintf("Could not get string: %v", err))
	}

	return val
}

func GetStringMap(f func() (map[string]string, error), internal_errors *[]string) map[string]string {
	val, err := f()
	if err != nil {
		*internal_errors = append(*internal_errors, fmt.Sprintf("Could not get string: %v", err))
	}

	return val
}
