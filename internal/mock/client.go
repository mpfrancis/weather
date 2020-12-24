package mock

import "net/http"

// Client is the mock client
type Client struct {
	GetFn      func(url string) (resp *http.Response, err error)
	GetInvoked bool
}

// Get is a mock function for the Get function on net/http.Client
func (c *Client) Get(url string) (*http.Response, error) {
	c.GetInvoked = true
	return c.GetFn(url)
}
