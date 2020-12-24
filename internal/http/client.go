package http

import "net/http"

// Clienter is an interface for net/http Client
type Clienter interface {
	Get(url string) (resp *http.Response, err error)
}

// DefaultClient serves as the default http client for the http layer
var DefaultClient = &http.Client{}
