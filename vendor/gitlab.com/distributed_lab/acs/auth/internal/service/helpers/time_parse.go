package helpers

import (
	"time"
)

func ParseToUnix(timeStr string) int64 {
	duration, _ := time.ParseDuration(timeStr)
	return time.Now().Add(duration).Unix()
}
