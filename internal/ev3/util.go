package ev3

import "time"

func DurationMs(ms int32) time.Duration {
	return time.Duration(ms * 1000 * 1000)
}
