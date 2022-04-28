package yagolangtbotapi

import (
	"log"
)

const (
	LogLevelInfo = iota
	LogLevelWarning
	LogLevelError
	LogLevelFatal
)

func writeLog(level int, to *log.Logger, format string, v ...interface{}) {
	if to == nil {
		return
	}
	res := "[UnknownLogLevel]: "
	switch level {
	case LogLevelInfo:
		res = "[INFO]           : "
	case LogLevelWarning:
		res = "[WARNING]        : "
	case LogLevelError:
		res = "[ERROR]          : "
	case LogLevelFatal:
		res = "[FATAL]          : "
	}
	to.Printf(res+"golangtbotapi: "+format, v...)
}
