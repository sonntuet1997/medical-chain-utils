package common

import (
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"os"
	"time"
)

func InitLogger(appCtx *cli.Context) *logrus.Logger {
	logger := logrus.New()
	host, _ := os.Hostname()
	isVerbose := appCtx.Bool("verbose")
	logger.Level = logrus.InfoLevel
	if isVerbose {
		logger.Level = logrus.DebugLevel
	}
	logger.Formatter = &logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "timestamp",
			logrus.FieldKeyLevel: "severity",
			logrus.FieldKeyMsg:   "message",
			"version":            appCtx.String("runtime-version"),
			"host":               host,
		},
		TimestampFormat: time.RFC3339Nano,
	}
	logger.Out = os.Stdout
	return logger
}
func InitLoggerWithoutCLIContext() *logrus.Logger {
	logger := logrus.New()
	host, _ := os.Hostname()
	logger.Level = logrus.DebugLevel
	logger.Formatter = &logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "timestamp",
			logrus.FieldKeyLevel: "severity",
			logrus.FieldKeyMsg:   "message",
			"version":            "0.0.1-debugger",
			"host":               host,
		},
		TimestampFormat: time.RFC3339Nano,
	}
	logger.Out = os.Stdout
	return logger
}
