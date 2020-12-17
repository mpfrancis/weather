package os

import (
	"errors"
	"os"

	"github.com/mpfrancis/weather"
)

const (
	envBaseURL = "WEATHER_BASEURL"
	envAPIKey  = "WEATHER_APIKEY"
)

var (
	MissingBaseURL = errors.New("WEATHER_BASEURL environment variable is required")
	MissingAPIKey  = errors.New("WEATHER_APIKEY environment variable is required")
)

func GetConfig() (*weather.Config, error) {
	var cfg weather.Config

	cfg.BaseURL = os.Getenv(envBaseURL)
	cfg.APIKey = os.Getenv(envAPIKey)

	if cfg.BaseURL == "" {
		return nil, MissingBaseURL
	}

	if cfg.APIKey == "" {
		return nil, MissingAPIKey
	}

	return &cfg, nil
}
