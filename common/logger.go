package common

import (
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"os"
	"strings"
	"time"
)

var (
	CommonGRPCFlag = []cli.Flag{
		&cli.StringFlag{
			Name:    "runtime-version",
			EnvVars: []string{"RUNTIME_VERSION"},
			Value:   "v1.0.0",
		},
		&cli.IntFlag{
			Name:    "grpc-port",
			Value:   50051,
			EnvVars: []string{"GRPC_PORT"},
			Usage:   "The port for exposing the gRPC endpoints for accessing",
		},
		&cli.IntFlag{
			Name:    "http-port",
			Value:   80,
			EnvVars: []string{"HTTP_PORT"},
			Usage:   "The port for exposing the api endpoints for accessing",
		},
		&cli.IntFlag{
			Name:    "pprof-port",
			Value:   6060,
			EnvVars: []string{"PPROF_PORT"},
			Usage:   "The port for exposing pprof endpoints",
		},
		&cli.BoolFlag{
			Name:    "disable-tracing",
			EnvVars: []string{"DISABLE_TRACING"},
			Usage:   "disable-tracing",
		},
		&cli.BoolFlag{
			Name:    "disable-profiler",
			EnvVars: []string{"DISABLE_PROFILER"},
			Usage:   "disable-profiler",
		},
		&cli.BoolFlag{
			Name:    "disable-stats",
			EnvVars: []string{"DISABLE_STATS"},
			Usage:   "disable-stats",
		},
		&cli.BoolFlag{
			Name:    "allow-kill",
			EnvVars: []string{"ALLOW_KILL"},
			Usage:   "allow remote request to kill server",
		},
	}
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
