package client

import (
	"backend/src/api/config"
	"fmt"
	"io"
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

func (hc *ClientHttp) DoLoginReq(user string) (*http.Response, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s:%v/get/%s", config.AuthHostname, config.AuthPort, user), nil)
	if err != nil {
		fmt.Println("Error creando la solicitud:", err)
		return nil, err
	}
	return hc.client.Do(req)

}

func (hc *ClientHttp) LoginRequest(user string, password string) (*http.Response, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s:%v/allowedusersLogin", config.AuthHostname, config.AuthPort), nil)
	if err != nil {
		fmt.Println("Error creando la solicitud:", err)
		return nil, err
	}
	req.SetBasicAuth(user, password)
	req.Header.Set("Content-Type", "application/json")
	return hc.client.Do(req)
}

func (hc *ClientHttp) AthznRequestV1(body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s:%v/allowedusers", config.AuthHostname, config.AuthPort), body)
	if err != nil {
		fmt.Println("Error creando la solicitud:", err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return hc.client.Do(req)
}

func (hc *ClientHttp) AthznRequestV2(user string, password string) (*http.Response, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s:%v/allowedusersV2", config.AuthHostname, config.AuthPort), nil)
	if err != nil {
		fmt.Println("Error creando la solicitud:", err)
		return nil, err
	}
	req.SetBasicAuth(user, password)
	req.Header.Set("Content-Type", "application/json")
	return hc.client.Do(req)
}
