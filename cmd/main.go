package main

import (
	"github.com/mpfrancis/weather/internal/http"
	"github.com/mpfrancis/weather/internal/os"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := run(); err != nil {
		logrus.Fatal(err)
	}
}

func run() error {
	cfg, err := os.GetConfig()
	if err != nil {
		return err
	}

	return http.NewServer(cfg, http.DefaultClient).ListenAndServe()
}
