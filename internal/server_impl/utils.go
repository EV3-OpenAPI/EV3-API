package server_impl

import (
	"fmt"
	"time"
)

func GetInt32(f func() (int, error), internalErrors *[]string) int32 {
	val, err := f()
	if err != nil {
		*internalErrors = append(*internalErrors, fmt.Sprintf("Could not get int: %v", err))
	}

	return int32(val)
}

func GetFloat32(f func() (float64, error), internalErrors *[]string) float32 {
	val, err := f()
	if err != nil {
		*internalErrors = append(*internalErrors, fmt.Sprintf("Could not get float: %v", err))
	}

	return float32(val)
}

func GetDurationAsInt32(f func() (time.Duration, error), internalErrors *[]string) int32 {
	val, err := f()
	if err != nil {
		*internalErrors = append(*internalErrors, fmt.Sprintf("Could not get float: %v", err))
	}

	return int32(val.Milliseconds())
}

func GetString(f func() (string, error), internalErrors *[]string) string {
	val, err := f()
	if err != nil {
		*internalErrors = append(*internalErrors, fmt.Sprintf("Could not get string: %v", err))
	}

	return val
}

func GetStringMap(f func() (map[string]string, error), internalErrors *[]string) map[string]string {
	val, err := f()
	if err != nil {
		*internalErrors = append(*internalErrors, fmt.Sprintf("Could not get string: %v", err))
	}

	return val
}
