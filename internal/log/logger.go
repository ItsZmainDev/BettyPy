package log

import (
    "fmt"
    "time"
)

type LogLevel string

const (
    INFO    LogLevel = "INFO"
    WARNING LogLevel = "WARNING"
    ERROR   LogLevel = "ERROR"
	SUCCESS LogLevel = "SUCCESS"
)

type Logger struct{}

func NewLogger() *Logger {
    return &Logger{}
}

func (l *Logger) Log(level LogLevel, message string) {
    timestamp := time.Now().Format("2006-01-02 15:04:05")
    color := ""

    switch level {
    case INFO:
		// Blue
		color = "\033[34m"
    case WARNING:
		// Yellow
		color = "\033[33m"
    case ERROR:
		// Red
		color = "\033[31m"
	case SUCCESS:
		// Green
		color = "\033[32m"
    }

	// Reset color
    reset := "\033[0m"

    fmt.Printf("%s[%s] [%s] %s%s\n", color, timestamp, level, message, reset)
}

func (l *Logger) Info(message string) {
    l.Log(INFO, message)
}

func (l *Logger) Warning(message string) {
    l.Log(WARNING, message)
}

func (l *Logger) Error(message string) {
    l.Log(ERROR, message)
}

func (l *Logger) Success(message string) {
    l.Log(SUCCESS, message)
}