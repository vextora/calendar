package logs

import (
	"fmt"
	"strings"
	"time"
)

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
)

func logWithColor(level, color, format string, a ...any) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	var msg string

	if strings.Contains(format, "%") {
		msg = fmt.Sprintf(format, a...)
	} else {
		//args := append([]any{format}, a...)
		// msg = fmt.Sprint(args...)
		msg = format
		if len(a) > 0 {
			msg += " " + fmt.Sprint(a...)
		}
	}

	fmt.Printf("%s[%s] %s | %s%s\n", color, level, timestamp, msg, colorReset)
}

func Info(format string, a ...any)  { logWithColor("INFO", colorGreen, format, a...) }
func Warn(format string, a ...any)  { logWithColor("WARNING", colorYellow, format, a...) }
func Error(format string, a ...any) { logWithColor("ERROR", colorRed, format, a...) }
func Debug(format string, a ...any) { logWithColor("DEBUG", colorBlue, format, a...) }
