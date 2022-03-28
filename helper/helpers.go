package helper

import (
	"log"
	"strings"
	"time"
)

var (
	TimeLayout = "2006-01-02T15:04:05Z"
	DateLayout = "2006-01-02"
)

func CalculateTimeDifference(t time.Time) time.Duration {
	diff := t.Sub(time.Now())
	log.Println("time difference ", diff)
	return diff
}

func GetTimeTomorrow() time.Time {
	t := time.Now().UTC()
	f := t.Add(time.Hour*24).Format(TimeLayout)
	elms := strings.Split(f, "T")
	f = elms[0] + "T00:00:00Z"
	t, err := time.Parse(TimeLayout, f)
	if err != nil {
		log.Println(err)
	}
	return t
}

func PointToString(text string) *string {
	return &text
}
