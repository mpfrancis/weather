package os

import (
	"errors"
	"os"
	"reflect"
	"testing"

	"github.com/mpfrancis/weather"
)

type Case struct {
	name           string
	baseURL        string
	apiKey         string
	expectedError  error
	expectedConfig *weather.Config
}

func TestGetConfig(t *testing.T) {
	cases := []Case{
		{"Success", "url", "key", nil, &weather.Config{BaseURL: "url", APIKey: "key"}},
		{"Missing URL", "", "key", MissingBaseURL, nil},
		{"Missing API Key", "url", "", MissingAPIKey, nil},
	}

	for i := range cases {
		if err := os.Setenv(envBaseURL, cases[i].baseURL); err != nil {
			t.Fatal(err)
		}
		if err := os.Setenv(envAPIKey, cases[i].apiKey); err != nil {
			t.Fatal(err)
		}

		cfg, err := GetConfig()
		if !errors.Is(err, cases[i].expectedError) {
			t.Fatalf("Test %s: Expected error '%s' but got '%s'", cases[i].name, cases[i].expectedError, err)
		}

		if !reflect.DeepEqual(cfg, cases[i].expectedConfig) {
			t.Fatalf("Test %s: Expected config '%+v' but got '%+v'", cases[i].name, cases[i].expectedConfig, cfg)
		}
	}
}