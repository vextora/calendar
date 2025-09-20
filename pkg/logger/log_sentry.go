package logger

import (
	logs "calendarapi/pkg/logutil"
	"fmt"

	sentry "github.com/getsentry/sentry-go"
)

func InfoSentry(format string, a ...any) {
	msg := fmt.Sprintf(format, a...)
	logs.Info("%v", msg)
	sentry.CaptureMessage(msg)
}

func WarnSentry(format string, a ...any) {
	msg := fmt.Sprintf(format, a...)
	logs.Warn("%v", msg)
	sentry.CaptureMessage(msg)
}

func DebugSentry(format string, a ...any) {
	msg := fmt.Sprintf(format, a...)
	logs.Debug("%v", msg)
	sentry.CaptureMessage(msg)
}

func ErrorSentry(err error, format string, a ...any) {
	msg := fmt.Sprintf(format, a...)
	logs.Error("%s %v", msg, err)
	if err != nil {
		sentry.CaptureException(err)
	}
}
