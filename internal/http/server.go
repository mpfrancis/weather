package http

import (
	"fmt"
	"net/http"

	"github.com/mpfrancis/weather"
	"github.com/sirupsen/logrus"
)

// Server is the weather API's server object.
type Server struct {
	*http.Server
}

// NewServer creates a new instance of the server object for serving up the API.
func NewServer(cfg *weather.Config) *Server {
	mux := http.NewServeMux()
	mux.Handle("/weather", recovery(NewWeatherHandler(cfg)))
	return &Server{&http.Server{Addr: cfg.ServerAddress, Handler: mux}}
}

func recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				logrus.Error(err)

				http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
				return
			}

		}()

		next.ServeHTTP(w, r)
	})
}
