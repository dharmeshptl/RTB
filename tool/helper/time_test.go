package helper_test

import (
	"fmt"
	"go_rtb/internal/tool/helper"
	"testing"
)

func TestGetLogId(t *testing.T) {
	id := helper.GetLogId(12, 13, 14)
	fmt.Println(id)
}

func TestGetCurrentStatHour(t *testing.T) {
	currentStatHour := helper.GetCurrentStatHour()
	fmt.Println(currentStatHour)
}

func TestGetTotalSecondsOfCurrentHour(t *testing.T) {
	seconds := helper.GetTotalSecondsOfCurrentHour()
	fmt.Println(seconds)
}
