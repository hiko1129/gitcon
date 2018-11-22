package infrastructure

import "net/http"

// HTTP struct
type HTTP struct{}

// NewHTTP func
func NewHTTP() *HTTP {
	return &HTTP{}
}

// Get func
func (h *HTTP) Get(url string) (*http.Response, error) {
	return http.Get(url)
}
