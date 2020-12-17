package os

import (
	"errors"
	"os"

	"github.com/mpfrancis/weather"
)

const (
	envBaseURL = "WEATHER_BASEURL"
	envAPIKey  = "WEATHER_APIKEY"
	envUnits   = "WEATHER_UNITS"
	envAddr    = "SERVER_ADDRESS"
)

var (
	errMissingBaseURL = errors.New("WEATHER_BASEURL environment variable is required")
	errMissingAPIKey  = errors.New("WEATHER_APIKEY environment variable is required")
	errInvalidUnits   = errors.New("Invalid units, use: standard, metric, imperial. Default: metric")
)

// GetConfig gets configuration environment variables and returns them in a config object.
// Environment variables WEATHER_BASEURL and WEATHER_APIKEY are required to be set.
func GetConfig() (*weather.Config, error) {
	var cfg weather.Config

	cfg.BaseURL = os.Getenv(envBaseURL)
	cfg.APIKey = os.Getenv(envAPIKey)
	cfg.Units = weather.Unit(os.Getenv(envUnits))
	cfg.ServerAddress = os.Getenv(envAddr)

	if cfg.BaseURL == "" {
		return nil, errMissingBaseURL
	}

	if cfg.APIKey == "" {
		return nil, errMissingAPIKey
	}

	switch cfg.Units {
	case "standard", "metric", "imperial":
	case "":
		cfg.Units = "metric"
	default:
		return nil, errInvalidUnits
	}

	if cfg.ServerAddress == "" {
		cfg.ServerAddress = ":10000"
	}

	return &cfg, nil
}
