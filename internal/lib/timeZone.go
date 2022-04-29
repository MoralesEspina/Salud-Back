package lib

import (
	"fmt"
	"time"
)

type TimeInfo struct {
	Year       int
	Month      int
	Day        int
	Hour       int
	Minute     int
	Second     int
	Milisecond int64
	DateTime   string
}

func TimeZone(region string) TimeInfo {
	location, err := time.LoadLocation(region)
	if err != nil {
		panic(region + " is unknown")
	}

	now := time.Now().In(location)
	year, month, day := now.UTC().Date()

	return TimeInfo{
		Year:       year,
		Month:      int(month),
		Day:        day,
		Hour:       now.Hour(),
		Minute:     now.Minute(),
		Second:     now.Second(),
		Milisecond: int64(now.Nanosecond()),
		DateTime:   fmt.Sprintf("%d-%d-%d %d:%d:%d", year, month, day, now.Hour(), now.Minute(), now.Second()),
	}

}
