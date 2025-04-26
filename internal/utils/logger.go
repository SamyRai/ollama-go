package utils

import (
	"log"
	"os"
)

// Logger is a simple logger instance for the client.
var Logger = log.New(os.Stdout, "[OLLAMA] ", log.LstdFlags|log.Lshortfile)

// Info logs an informational message.
func Info(v ...interface{}) {
	Logger.Println("[INFO]", v)
}

// Error logs an error message.
func Error(v ...interface{}) {
	Logger.Println("[ERROR]", v)
}

// Debug logs a debug message.
func Debug(v ...interface{}) {
	Logger.Println("[DEBUG]", v)
}
