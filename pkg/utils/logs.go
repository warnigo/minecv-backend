package utils

import (
	"log"
	"time"
)

// @Example Usage:
// logs.LogError("Database connection failed")
// logs.LogWarning("API response time is slow")
// logs.LogSuccess("User created successfully")

// Initialize logger to disable default timestamps
func init() {
	log.SetFlags(0)
}

// logMessage logs a formatted message with a specific prefix and color
func logMessage(prefix, color, message string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	log.Printf("%s[%s] %s - %s%s\n", color, prefix, timestamp, message, colorReset)
}

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorYellow = "\033[33m"
	colorGreen  = "\033[32m"
)

// LogError logs an error message in red color
func LogError(message string) {
	logMessage("ERROR", colorRed, message)
}

// LogWarning logs a warning message in yellow color
func LogWarning(message string) {
	logMessage("WARNING", colorYellow, message)
}

// LogSuccess logs a success message in green color
func LogSuccess(message string) {
	logMessage("SUCCESS", colorGreen, message)
}
