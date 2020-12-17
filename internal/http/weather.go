package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/mpfrancis/weather"
	"github.com/patrickmn/go-cache"
)

// WeatherHandler is the handler for the /weather endpoint.
type WeatherHandler struct {
	cfg           *weather.Config
	responseCache *cache.Cache
}

// NewWeatherHandler returns a new instance of the weather http handler.
func NewWeatherHandler(cfg *weather.Config) *WeatherHandler {
	return &WeatherHandler{cfg: cfg, responseCache: cache.New(2*time.Minute, time.Minute)}
}

// ServeHTTP handles a weather request.
// This handler will hit the open weather API and return a more human readable response.
func (h *WeatherHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Check cache
	if hr, ok := h.responseCache.Get(r.URL.String()); ok {
		if err := json.NewEncoder(w).Encode(hr); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		return
	}

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
	response, err := http.Get(fmt.Sprintf("%s/weather?q=%s,%s&units=%s&appid=%s", h.cfg.BaseURL, city, country, h.cfg.Units, h.cfg.APIKey))
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

	forecast := r.FormValue("forecast")
	if forecast != "" {
		day, err := strconv.Atoi(forecast)
		if err != nil {
			http.Error(w, "Query parameter 'country' is required", http.StatusUnprocessableEntity)
			return
		}

		// Call open weather API
		response, err := http.Get(fmt.Sprintf("%s/onecall?lat=%g&lon=%g&units=%s&appid=%s", h.cfg.BaseURL, owr.Coord.Lat, owr.Coord.Lon, h.cfg.Units, h.cfg.APIKey))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Parse the response
		var ocr weather.OneCallResponse
		if err := json.NewDecoder(response.Body).Decode(&ocr); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		hr.Forecast = &ocr.Daily[day]
	}

	h.responseCache.Set(r.URL.String(), hr, cache.DefaultExpiration)

	if err := json.NewEncoder(w).Encode(hr); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
}
