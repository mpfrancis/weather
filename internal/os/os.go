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
)

var (
	MissingBaseURL = errors.New("WEATHER_BASEURL environment variable is required")
	MissingAPIKey  = errors.New("WEATHER_APIKEY environment variable is required")
	InvalidUnits   = errors.New("Invalid units, use: standard, metric, imperial. Default: metric")
)

// GetConfig gets configuration environment variables and returns them in a config object.
// Environment variables WEATHER_BASEURL and WEATHER_APIKEY are required to be set.
func GetConfig() (*weather.Config, error) {
	var cfg weather.Config

	cfg.BaseURL = os.Getenv(envBaseURL)
	cfg.APIKey = os.Getenv(envAPIKey)
	cfg.Units = weather.Unit(os.Getenv(envUnits))

	if cfg.BaseURL == "" {
		return nil, MissingBaseURL
	}

	if cfg.APIKey == "" {
		return nil, MissingAPIKey
	}

	switch cfg.Units {
	case "standard", "metric", "imperial":
	case "":
		cfg.Units = "metric"
	default:
		return nil, InvalidUnits
	}

	return &cfg, nil
}
