package http

import (
	"fmt"
	"net/http"

	"github.com/mpfrancis/weather"
	"github.com/sirupsen/logrus"
)

type Server struct {
	*http.Server
}

func NewServer(cfg *weather.Config) *Server {
	mux := http.NewServeMux()
	mux.Handle("/weather", recovery(NewWeatherHandler(cfg)))

	// TODO: Configure
	s := &Server{&http.Server{Addr: ":10000", Handler: mux}}

	return s
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
