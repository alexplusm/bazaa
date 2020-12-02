package utils

import (
	"strconv"
	"time"
)

// TODO: timeutils package

func FromTimestampToTime(timestamp string) (time.Time, error) {
	ts, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		return time.Now(), err
	}

	return time.Unix(ts, 0), nil
}

func TrimTime(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}
