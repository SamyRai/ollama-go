package utils

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"time"
)

// LogLevel represents the severity level of log messages
type LogLevel int

const (
	// LogLevelNone disables all logging
	LogLevelNone LogLevel = iota
	// LogLevelError logs only errors
	LogLevelError
	// LogLevelWarn logs warnings and errors
	LogLevelWarn
	// LogLevelInfo logs informational messages, warnings, and errors
	LogLevelInfo
	// LogLevelDebug logs debug messages and all above
	LogLevelDebug
)

var levelNames = map[LogLevel]string{
	LogLevelNone:  "NONE",
	LogLevelError: "ERROR",
	LogLevelWarn:  "WARN",
	LogLevelInfo:  "INFO",
	LogLevelDebug: "DEBUG",
}

// Logger provides structured logging functionality
type Logger struct {
	mu      sync.Mutex
	level   LogLevel
	logger  *log.Logger
	service string
}

// Global logger instance with default settings
var globalLogger *Logger
var once sync.Once

// GetLogger returns the global logger instance
func GetLogger() *Logger {
	once.Do(func() {
		// Default to INFO level
		level := LogLevelInfo

		// Check environment variable for log level
		if envLevel := os.Getenv("OLLAMA_LOG_LEVEL"); envLevel != "" {
			switch envLevel {
			case "NONE":
				level = LogLevelNone
			case "ERROR":
				level = LogLevelError
			case "WARN":
				level = LogLevelWarn
			case "INFO":
				level = LogLevelInfo
			case "DEBUG":
				level = LogLevelDebug
			}
		}

		globalLogger = NewLogger("OLLAMA", level)
	})
	return globalLogger
}

// NewLogger creates a new logger instance with the specified service name and log level
func NewLogger(service string, level LogLevel) *Logger {
	return &Logger{
		service: service,
		level:   level,
		logger:  log.New(os.Stderr, "", 0), // No prefix or flags, we handle those ourselves
	}
}

// SetOutput sets the output destination for the logger
func (l *Logger) SetOutput(w io.Writer) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.logger.SetOutput(w)
}

// SetLevel sets the logging level
func (l *Logger) SetLevel(level LogLevel) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.level = level
}

// GetLevel returns the current logging level
func (l *Logger) GetLevel() LogLevel {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.level
}

// formatMessage creates a standardized log message format
func (l *Logger) formatMessage(level LogLevel, message string) string {
	timestamp := time.Now().Format("2006-01-02T15:04:05.000Z07:00")
	return fmt.Sprintf("%s [%s] [%s] %s", timestamp, l.service, levelNames[level], message)
}

// log writes a message to the log at the specified level
func (l *Logger) log(level LogLevel, format string, v ...interface{}) {
	if l.level >= level {
		l.mu.Lock()
		defer l.mu.Unlock()
		message := fmt.Sprintf(format, v...)
		l.logger.Println(l.formatMessage(level, message))
	}
}

// Debug logs a debug message
func (l *Logger) Debug(format string, v ...interface{}) {
	l.log(LogLevelDebug, format, v...)
}

// Info logs an informational message
func (l *Logger) Info(format string, v ...interface{}) {
	l.log(LogLevelInfo, format, v...)
}

// Warn logs a warning message
func (l *Logger) Warn(format string, v ...interface{}) {
	l.log(LogLevelWarn, format, v...)
}

// Error logs an error message
func (l *Logger) Error(format string, v ...interface{}) {
	l.log(LogLevelError, format, v...)
}

// DebugEnabled returns true if debug logging is enabled
func (l *Logger) DebugEnabled() bool {
	return l.level >= LogLevelDebug
}

// InfoEnabled returns true if info logging is enabled
func (l *Logger) InfoEnabled() bool {
	return l.level >= LogLevelInfo
}

// WarnEnabled returns true if warning logging is enabled
func (l *Logger) WarnEnabled() bool {
	return l.level >= LogLevelWarn
}

// ErrorEnabled returns true if error logging is enabled
func (l *Logger) ErrorEnabled() bool {
	return l.level >= LogLevelError
}

// Convenience functions for using the global logger

// Debug logs a debug message using the global logger
func Debug(format string, v ...interface{}) {
	GetLogger().Debug(format, v...)
}

// Info logs an informational message using the global logger
func Info(format string, v ...interface{}) {
	GetLogger().Info(format, v...)
}

// Warn logs a warning message using the global logger
func Warn(format string, v ...interface{}) {
	GetLogger().Warn(format, v...)
}

// Error logs an error message using the global logger
func Error(format string, v ...interface{}) {
	GetLogger().Error(format, v...)
}

// SetGlobalLevel sets the logging level for the global logger
func SetGlobalLevel(level LogLevel) {
	GetLogger().SetLevel(level)
}

// SetGlobalOutput sets the output destination for the global logger
func SetGlobalOutput(w io.Writer) {
	GetLogger().SetOutput(w)
}
