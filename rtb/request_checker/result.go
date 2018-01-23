package request_checker

import (
	"fmt"
	"go_rtb/internal/tool/helper"
	"go_rtb/internal/tool/logger"
	"time"
)

type CheckerResult struct {
	success     bool
	description string
	err         error
	checkTime   time.Time
}

func NewCheckerResult(
	success bool,
	description string,
	err error,
) CheckerResult {
	return CheckerResult{
		success:     success,
		description: description,
		err:         err,
		checkTime:   helper.GetCurrentTime(),
	}
}

func (result *CheckerResult) Success() bool {
	return result.success
}

func (result *CheckerResult) GetDescription() string {
	return result.description
}

func (result *CheckerResult) GetError() error {
	return result.err
}

func (result *CheckerResult) GetCheckTime() time.Time {
	return result.checkTime
}

func (result *CheckerResult) ToAppLog() logger.AppLog {
	statusStr := "Successed"
	if result.Success() == false {
		statusStr = "Failed"
	}
	return logger.NewAppLog(
		logger.DspRequestCheckerLog,
		result.err,
		result.checkTime,
		fmt.Sprintf("%s - %s", statusStr, result.description),
	)
}
