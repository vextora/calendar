package logger

import (
	"calendarapi/pkg/config"
	"time"

	sentry "github.com/getsentry/sentry-go"
	"go.uber.org/zap"
)

func InitSentry() {
	dsn := config.GetEnvString(config.SentryDsn)
	env := config.GetEnvString(config.Environment)

	if dsn == "" {
		return
	}

	err := sentry.Init(sentry.ClientOptions{
		Dsn:              dsn,
		Environment:      env,
		TracesSampleRate: 1.0,
	})

	if err != nil {
		Logger.Error("Sentry init failed: ", zap.Error(err))

	}
}

func FlushSentry() {
	sentry.Flush(2 * time.Second)
}
