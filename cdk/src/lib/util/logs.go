package util

import (
	"github.com/sirupsen/logrus"
)

func SetLogLevel(logger *logrus.Logger, level string) {
	//	_ = os.Setenv("LOG_LEVEL", "debug")
	switch level {
	case "error":
		logger.SetLevel(logrus.ErrorLevel)
	case "info":
		logger.SetLevel(logrus.InfoLevel)
	case "debug":
		logger.SetLevel(logrus.DebugLevel)
	default:
		logger.SetLevel(logrus.DebugLevel)
	}
}
