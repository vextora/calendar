package logger

import (
	"log"
	"oncomapi/pkg/config"
	"oncomapi/pkg/shared"
	"time"

	sentry "github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

type Config struct {
	Environment string
	SentryDSN   string
}

func InitZap() {
	environment := config.GetEnvString(config.Environment)
	zapOutputPath := config.GetEnvString(config.ZapOutputPath)
	zapErrorOutputPath := config.GetEnvString(config.ZapErrorOutputPath)

	if zapOutputPath == "" {
		zapOutputPath = "stdout"
	}
	if zapErrorOutputPath == "" {
		zapErrorOutputPath = "stderr"
	}

	var zapCfg zap.Config
	if environment == "development" {
		zapCfg = zap.NewDevelopmentConfig()
	} else {
		zapCfg = zap.NewProductionConfig()
	}

	zapCfg.Encoding = "json"
	zapCfg.OutputPaths = []string{zapOutputPath}
	zapCfg.ErrorOutputPaths = []string{zapErrorOutputPath}
	zapCfg.EncoderConfig = zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.MillisDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	var err error
	Logger, err = zapCfg.Build()
	if err != nil {
		log.Fatalf("failed to init zap logger: %v", err)
	}

	InitSentry()
}

// Flush before shutdown
func Sync() {
	_ = Logger.Sync()
	sentry.Flush(2 * time.Second)
}

// Function to take logger with trace ID
func GetLoggerWithTraceID(c *gin.Context) *zap.Logger {
	logger := Logger

	traceID, exists := c.Get(shared.TraceIDKey)
	if exists {
		if tid, ok := traceID.(string); ok {
			return logger.With(zap.String("trace_id", tid))
		}
	}

	return logger
}
