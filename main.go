package main

import (
	"fmt"
	"log"

	"github.com/shubangii/playground/config"
	"github.com/shubangii/playground/logger"
	"github.com/shubangii/playground/utils"
	"go.opentelemetry.io/otel/exporters/prometheus"
)

func main() {
	fmt.Println("=== Go Module Dependencies Demo ===")

	// Load configuration
	cfg := config.Load()
	fmt.Printf("App Name: %s\n", cfg.AppName)
	fmt.Printf("Version: %s\n", cfg.Version)
	fmt.Printf("Debug Mode: %t\n", cfg.Debug)

	// Initialize logger
	logger := logger.New(cfg.Debug)
	logger.Info("Application started successfully")
	logger.Debug("This is a debug message")

	// Use utilities
	data := []string{"apple", "banana", "cherry", "date"}
	logger.Info("Original data: " + utils.Join(data, ", "))

	reversed := utils.Reverse(data)
	logger.Info("Reversed data: " + utils.Join(reversed, ", "))

	unique := utils.Unique([]string{"apple", "banana", "apple", "cherry", "banana"})
	logger.Info("Unique items: " + utils.Join(unique, ", "))

	if cfg.Debug {
		logger.Debug("Application completed in debug mode")
	}

	log.Println("Main application finished")
	_, err := prometheus.New(prometheus.WithoutUnits())
	if err != nil {
		logger.Error("Failed to create Prometheus registry: " + err.Error())
	}
}
