package helper

import (
	"log"
	"time"
)

var (
	TimeLayout = "2006-01-02T15:04:05Z"
)

func CalculateTimeDifference(t time.Time) time.Duration {
	diff := t.Sub(time.Now())
	log.Println("time difference ", diff)
	return diff
}
