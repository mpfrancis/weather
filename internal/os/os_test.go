package os

import (
	"errors"
	"os"
	"testing"
	"time"

	"github.com/mpfrancis/weather"
	"github.com/stretchr/testify/assert"
)

type Case struct {
	name           string
	baseURL        string
	apiKey         string
	units          string
	addr           string
	cacheExpiry    string
	expectedError  error
	expectedConfig *weather.Config
}

func TestGetConfig(t *testing.T) {
	cases := []Case{
		{"Success", "url", "key", "imperial", ":11000", "5m", nil, &weather.Config{BaseURL: "url", APIKey: "key", Units: "imperial", ServerAddress: ":11000", CacheExpiration: "5m", CacheExpirationDur: 5 * time.Minute}},
		{"Defaults", "url", "key", "", "", "", nil, &weather.Config{BaseURL: "url", APIKey: "key", Units: "metric", ServerAddress: ":10000", CacheExpirationDur: 2 * time.Minute}},
		{"Missing URL", "", "key", "", "", "", errMissingBaseURL, nil},
		{"Missing API Key", "url", "", "", "", "", errMissingAPIKey, nil},
		{"Invalid Units", "url", "key", "abc", "", "", errInvalidUnits, nil},
	}

	for i := range cases {
		if err := os.Setenv(envBaseURL, cases[i].baseURL); err != nil {
			t.Fatal(err)
		}
		if err := os.Setenv(envAPIKey, cases[i].apiKey); err != nil {
			t.Fatal(err)
		}
		if err := os.Setenv(envUnits, cases[i].units); err != nil {
			t.Fatal(err)
		}
		if err := os.Setenv(envAddr, cases[i].addr); err != nil {
			t.Fatal(err)
		}
		if err := os.Setenv(envCacheExpiration, cases[i].cacheExpiry); err != nil {
			t.Fatal(err)
		}

		cfg, err := GetConfig()
		if !errors.Is(err, cases[i].expectedError) {
			t.Fatalf("Test %s: Expected error '%s' but got '%s'", cases[i].name, cases[i].expectedError, err)
		}

		assert.Equal(t, cfg, cases[i].expectedConfig)
	}
}
