package logger

import "github.com/sirupsen/logrus"

var NvwaLog *logrus.Logger

func InitNvwaLogger(level string) {
	NvwaLog = logrus.New()
	switch level {
	case "debug":
		NvwaLog.Level = logrus.DebugLevel
	case "info":
		NvwaLog.Level = logrus.InfoLevel
	case "warn":
		NvwaLog.Level = logrus.WarnLevel
	case "error":
		NvwaLog.Level = logrus.ErrorLevel
	case "fatal":
		NvwaLog.Level = logrus.FatalLevel
	case "panic":
		NvwaLog.Level = logrus.PanicLevel
	default:
		NvwaLog.Level = logrus.DebugLevel
	}
}
