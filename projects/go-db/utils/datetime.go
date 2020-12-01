package utils

import (
	"strconv"
	"time"
)

func FromTimestampToTime(timestamp string) (time.Time, error) {
	ts, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		return time.Now(), err
	}

	return time.Unix(ts, 0), nil
}
