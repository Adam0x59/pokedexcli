package pokeapi

import (
	"net/http"
	"time"
)

// Struct to define http client
type Client struct {
	httpClient http.Client
}

// Function to open and return a new http client with a specified timeout
func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
