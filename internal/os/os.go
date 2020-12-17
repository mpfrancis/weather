package os

import (
	"errors"
	"os"
	"time"

	"github.com/mpfrancis/weather"
	"github.com/sirupsen/logrus"
)

const (
	envBaseURL         = "WEATHER_BASEURL"
	envAPIKey          = "WEATHER_APIKEY"
	envUnits           = "WEATHER_UNITS"
	envAddr            = "SERVER_ADDRESS"
	envCacheExpiration = "CACHE_EXPIRATION"
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
	cfg.CacheExpiration = os.Getenv(envCacheExpiration)

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

	var err error
	cfg.CacheExpirationDur, err = time.ParseDuration(cfg.CacheExpiration)
	if err != nil {
		if cfg.CacheExpiration != "" {
			logrus.Warn("Unable to parse cache expiration, defaulting to two minutes")
		}
		cfg.CacheExpirationDur = 2 * time.Minute
	}

	return &cfg, nil
}
