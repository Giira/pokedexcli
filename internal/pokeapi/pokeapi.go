package pokeapi

import (
	"net/http"
	"time"
)

const (
	apiUrlBase = "https://pokeapi.co/api/v2"
)

type Client struct {
	httpclient http.Client
}

func newClient() Client {
	return Client{
		httpclient: http.Client{
			Timeout: 5 * time.Second,
		},
	}
}
