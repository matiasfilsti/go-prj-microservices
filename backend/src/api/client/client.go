package client

import (
	"net/http"
	"time"
)

type ClientHttp struct {
	client *http.Client
}

func NewClientHttp(Client *http.Client) *ClientHttp {
	return &ClientHttp{
		client: Client,
	}
}

func CreateHttpClient() *http.Client {
	return &http.Client{
		Timeout: 3 * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:       10,
			IdleConnTimeout:    30 * time.Second,
			DisableCompression: true,
		},
	}
}
