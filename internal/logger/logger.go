package logger

import (
	"github.com/sirupsen/logrus"
)

func NewLogger(env string) *logrus.Logger {
	log := logrus.Logger{
		Formatter: &logrus.TextFormatter{
			DisableTimestamp: false,
			TimestampFormat:  "2006-01-02 15:04:05",
			FullTimestamp:    false,
		},
	}

	switch env {
	case "local":
		log.SetLevel(logrus.DebugLevel)
	case "prod":
		log.SetLevel(logrus.InfoLevel)
	}

	return &log
}
