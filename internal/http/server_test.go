package http

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"testing"
	"time"

	"github.com/mpfrancis/weather"
	"github.com/mpfrancis/weather/internal/mock"
	"github.com/stretchr/testify/assert"
)

func TestServer(t *testing.T) {
	// Get ephemeral port
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}
	port := listener.Addr().(*net.TCPAddr).Port
	assert.Equal(t, nil, listener.Close())

	// Create server with mock client for panic test
	cfg := weather.Config{ServerAddress: fmt.Sprintf(":%d", port)}
	var mockClient mock.Client
	mockClient.GetFn = func(url string) (resp *http.Response, err error) {
		panic("PANIC TEST")
	}
	s := NewServer(&cfg, &mockClient)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	defer s.Shutdown(ctx)

	go func(s *Server) {
		if err := s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			t.Error(err)
		}
	}(s)

	baseURL := fmt.Sprintf("http://localhost:%d", port)

	// Simple health check
	{
		resp, err := http.Get(baseURL + "/healthcheck")
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, 200, resp.StatusCode)
	}

	// Induce a panic
	{
		resp, err := http.Get(baseURL + "/weather?city=Bogota&country=co")
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, 500, resp.StatusCode)
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, "PANIC TEST\n", string(body))
	}

	// Ensure server still operates after panic
	{
		resp, err := http.Get(baseURL + "/healthcheck")
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, 200, resp.StatusCode)
	}
}
