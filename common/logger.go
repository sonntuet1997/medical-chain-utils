package common

import (
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"os"
	"strings"
	"time"
)

var (
	LoggerFlag = []cli.Flag{
		&cli.StringFlag{
			Name:    "log-level",
			EnvVars: []string{"LOG_LEVEL"},
			Usage:   "Log level: (panic|fatal|error|warn|warning|info|debug|trace)",
			Value:   "info",
		},
		&cli.StringFlag{
			Name:    "log-format",
			EnvVars: []string{"LOG_FORMAT"},
			Usage:   "Log format: (plain|json)",
			Value:   "json",
		},
	}
)

func InitLogger(appCtx *cli.Context) *logrus.Logger {
	logger := logrus.New()
	host, _ := os.Hostname()
	logLevel := appCtx.String("log-level")
	level, err := logrus.ParseLevel(logLevel)
	logger.Out = os.Stdout
	if err != nil {
		logger.Fatalf("Unknown log-level type: %s", logLevel)
		return logger
	}
	logger.Level = level
	logFormat := appCtx.String("log-format")
	switch strings.ToLower(logFormat) {
	case "json":
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
	case "plain":
		logger.Formatter = &logrus.TextFormatter{}
	default:
		logger.Fatalf("Unknown log-format type: %s", logFormat)
		return logger
	}
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
