package config

import (
	"os"
	"strconv"
)

// Config holds application configuration
type Config struct {
	AppName string
	Version string
	Debug   bool
	Port    int
}

// Load returns a new Config with default values and environment overrides
func Load() *Config {
	cfg := &Config{
		AppName: "Playground App",
		Version: "1.0.0",
		Debug:   false,
		Port:    8080,
	}

	// Override with environment variables if they exist
	if appName := os.Getenv("APP_NAME"); appName != "" {
		cfg.AppName = appName
	}

	if version := os.Getenv("APP_VERSION"); version != "" {
		cfg.Version = version
	}

	if debugStr := os.Getenv("DEBUG"); debugStr != "" {
		if debug, err := strconv.ParseBool(debugStr); err == nil {
			cfg.Debug = debug
		}
	}

	if portStr := os.Getenv("PORT"); portStr != "" {
		if port, err := strconv.Atoi(portStr); err == nil {
			cfg.Port = port
		}
	}

	return cfg
}

// IsProduction returns true if the app is not in debug mode
func (c *Config) IsProduction() bool {
	return !c.Debug
}