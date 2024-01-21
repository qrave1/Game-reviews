package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

func NewLogger() *logrus.Logger {
	logrus.New()
	log := logrus.Logger{
		Out: os.Stderr,
		Formatter: &logrus.JSONFormatter{
			DisableTimestamp: false,
			TimestampFormat:  "2006-01-02 15:04:05",
		},
		Level: logrus.InfoLevel,
	}

	return &log
}
