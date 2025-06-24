package utils

import (
	"strings"

	"github.com/shubangii/playground/logger"
)

// Join concatenates string slice with a separator
func Join(slice []string, sep string) string {
	return strings.Join(slice, sep)
}

// Reverse returns a new slice with elements in reverse order
func Reverse(slice []string) []string {
	reversed := make([]string, len(slice))
	for i, v := range slice {
		reversed[len(slice)-1-i] = v
	}
	return reversed
}

// Unique returns a new slice with duplicate elements removed
func Unique(slice []string) []string {
	seen := make(map[string]bool)
	result := []string{}

	for _, item := range slice {
		if !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}

	return result
}

// ProcessWithLogging demonstrates using another sub-module dependency
func ProcessWithLogging(data []string, debug bool) []string {
	logger := logger.New(debug)
	logger.Debug("Starting data processing")
	
	unique := Unique(data)
	logger.Info("Removed duplicates from data")
	
	reversed := Reverse(unique)
	logger.Info("Reversed processed data")
	
	logger.Debug("Data processing completed")
	return reversed
}