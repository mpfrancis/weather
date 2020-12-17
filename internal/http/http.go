package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mpfrancis/weather"
)

type WeatherHandler struct {
	cfg *weather.Config
}

// NewWeatherHandler returns a new instance of the weather http handler.
func NewWeatherHandler(cfg *weather.Config) *WeatherHandler {
	return &WeatherHandler{cfg: cfg}
}

// ServeHTTP handles a weather request.
// This handler will hit the open weather API and return a more human readable response.
func (h *WeatherHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Parse input parameters
	city := r.FormValue("city")
	if city == "" {
		http.Error(w, "Query parameter 'city' is required", http.StatusUnprocessableEntity)
		return
	}

	country := r.FormValue("country")
	if country == "" {
		http.Error(w, "Query parameter 'country' is required", http.StatusUnprocessableEntity)
		return
	}

	// Call open weather API
	response, err := http.Get(fmt.Sprintf("%s?q=%s,%s&units=%s&appid=%s", h.cfg.BaseURL, city, country, h.cfg.Units, h.cfg.APIKey))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Parse the response
	var owr weather.OpenWeatherResponse
	if err := json.NewDecoder(response.Body).Decode(&owr); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	hr := owr.ToHumanReadable(h.cfg.Units.Symbol())
	if err := json.NewEncoder(w).Encode(hr); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
}
