package logger

import "time"

type AppLogType string

const (
	DspRequestCheckerLog AppLogType = "DSP request checker"
	DspUrlCallLog        AppLogType = "DSP URL call log"
	RTBResponseCallLog   AppLogType = "Rtb response parse log"
)

type AppLog struct {
	logType     AppLogType
	err         error
	datetime    time.Time
	description string
}

func NewAppLog(
	logType AppLogType,
	err error,
	datetime time.Time,
	description string,
) AppLog {
	return AppLog{
		logType:     logType,
		err:         err,
		datetime:    datetime,
		description: description,
	}
}

func (log *AppLog) GetLogType() AppLogType {
	return log.logType
}

func (log *AppLog) GetError() error {
	return log.err
}

func (log *AppLog) GetTime() time.Time {
	return log.datetime
}

func (log *AppLog) GetDescription() string {
	return log.description
}
