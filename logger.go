package yagolangtbotapi

import (
	"log"
)

const (
	LogLevelInfo = iota
	LogLevelError
)

func writeLog(level int, to *log.Logger, format string, v ...interface{}) {
	if to == nil {
		return
	}
	res := "[UnknownLogLevel]: "
	switch level {
	case LogLevelInfo:
		res = "[INFO]           : "
	case LogLevelError:
		res = "[ERROR]          : "
	}
	to.Printf(res+"golangtbotapi: "+format, v...)
}
