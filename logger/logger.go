package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

// Logger represents a simple logger with different levels
type Logger struct {
	debug  bool
	prefix string
	logger *log.Logger
}

// New creates a new Logger instance
func New(debug bool) *Logger {
	fmt.Println("Initializing logger...")
	return &Logger{
		debug:  debug,
		prefix: "[PLAYGROUND] ",
		logger: log.New(os.Stdout, "", log.LstdFlags),
	}
}

// Info logs an info message
func (l *Logger) Info(message string) {
	l.logger.Printf("%sINFO: %s", l.prefix, message)
}

// Debug logs a debug message (only if debug mode is enabled)
func (l *Logger) Debug(message string) {
	if l.debug {
		l.logger.Printf("%sDEBUG: %s", l.prefix, message)
	}
}

// Error logs an error message
func (l *Logger) Error(message string) {
	l.logger.Printf("%sERROR: %s", l.prefix, message)
}

// Warn logs a warning message
func (l *Logger) Warn(message string) {
	l.logger.Printf("%sWARN: %s", l.prefix, message)
}

// LogWithTimestamp logs a message with a custom timestamp format
func (l *Logger) LogWithTimestamp(level, message string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("[%s] %s%s: %s\n", timestamp, l.prefix, level, message)
}
