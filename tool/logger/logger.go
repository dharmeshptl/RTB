package logger

import (
	"fmt"
	"log"
	"os"
)

var (
	trace   *log.Logger // Just about anything
	info    *log.Logger // Important information
	warning *log.Logger // Be concerned
	err     *log.Logger // Critical problem
)

const flag int = log.Ldate | log.Ltime | log.Lshortfile

func init() {
	trace = log.New(os.Stdout, "TRACE: ", flag)
	info = log.New(os.Stdout, "INFO: ", flag)
	warning = log.New(os.Stdout, "WARNING: ", flag)
	err = log.New(os.Stdout, "ERROR: ", flag)
}

func Trace(data interface{}) {
	trace.Println(data)
}

func Info(data interface{}) {
	info.Println(data)
}

func Warning(data interface{}) {
	warning.Println(data)
}

func Error(data interface{}) {
	err.Println(data)
}

func ShowLog(logs map[string][]AppLog) {
	for logType, logList := range logs {
		fmt.Println("LOG FROM STATE: " + logType)
		for idx, log := range logList {
			errStr := ""
			if log.GetError() != nil {
				errStr = " - ERROR: " + log.GetError().Error()
			}
			logStr := fmt.Sprintf(
				"%d. %s  %s",
				idx+1,
				log.GetDescription(),
				errStr,
			)
			trace.Println(logStr)
		}
	}
}
