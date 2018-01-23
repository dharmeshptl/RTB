package helper

import (
	"strconv"
	"time"
)

const dateFormat string = "2006010215"

func GetLogId(entityIds ...uint) string {
	t := time.Now().UTC()
	result := t.Format(dateFormat)
	for _, id := range entityIds {
		result += "_" + strconv.FormatUint(uint64(id), 10)
	}

	return result
}

func GetCurrentStatHour() uint {
	t := time.Now().UTC()
	result, _ := strconv.ParseUint(t.Format(dateFormat), 10, 64)
	return uint(result)
}

func GetTotalSecondsOfCurrentHour() uint64 {
	t := time.Now().UTC()
	return uint64(t.Second() + (t.Minute() * 60))
}

func GetCurrentTime() time.Time {
	return time.Now().UTC()
}
