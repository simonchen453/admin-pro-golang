package logger

import (
	"log"
	"os"
)

// LogLevel represents the severity level of a log entry
type LogLevel int

const (
	// DEBUG is for detailed information, typically of interest only when diagnosing problems
	DEBUG LogLevel = iota
	// INFO is for general informational messages
	INFO
	// WARN is for warning messages
	WARN
	// ERROR is for error messages
	ERROR
)

// Logger is a simple structured logger
// In production, consider using more advanced loggers like zap or zerolog
type Logger struct {
	prefix string
	level  LogLevel
}

// New creates a new logger with the given prefix
func New(prefix string) *Logger {
	return &Logger{
		prefix: prefix,
		level:  INFO,
	}
}

// SetLevel sets the minimum log level
func (l *Logger) SetLevel(level LogLevel) {
	l.level = level
}

// Debug logs a debug message
func (l *Logger) Debug(format string, args ...interface{}) {
	if l.level <= DEBUG {
		log.Printf("[DEBUG] "+l.prefix+" "+format, args...)
	}
}

// Info logs an info message
func (l *Logger) Info(format string, args ...interface{}) {
	if l.level <= INFO {
		log.Printf("[INFO] "+l.prefix+" "+format, args...)
	}
}

// Warn logs a warning message
func (l *Logger) Warn(format string, args ...interface{}) {
	if l.level <= WARN {
		log.Printf("[WARN] "+l.prefix+" "+format, args...)
	}
}

// Error logs an error message
func (l *Logger) Error(format string, args ...interface{}) {
	if l.level <= ERROR {
		log.Printf("[ERROR] "+l.prefix+" "+format, args...)
	}
}

// With returns a new logger with the given prefix
func (l *Logger) With(prefix string) *Logger {
	return &Logger{
		prefix: l.prefix + " " + prefix,
		level:  l.level,
	}
}

// Global logger instance
var std = New("")

// SetLevel sets the level for the global logger
func SetLevel(level LogLevel) {
	std.SetLevel(level)
}

// Debug logs a debug message using the global logger
func Debug(format string, args ...interface{}) {
	std.Debug(format, args...)
}

// Info logs an info message using the global logger
func Info(format string, args ...interface{}) {
	std.Info(format, args...)
}

// Warn logs a warning message using the global logger
func Warn(format string, args ...interface{}) {
	std.Warn(format, args...)
}

// Error logs an error message using the global logger
func Error(format string, args ...interface{}) {
	std.Error(format, args...)
}

// With returns a new logger with the given prefix using the global logger
func With(prefix string) *Logger {
	return std.With(prefix)
}

// Init initializes the global logger
func Init() {
	// Set output to stdout
	log.SetOutput(os.Stdout)
	// Set flags to include date and time
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
