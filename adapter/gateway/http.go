package gateway

import "net/http"

// HTTP interface
type HTTP interface {
	Get(string) (*http.Response, error)
}
