package client

import (
	"backend/src/api/config"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
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

func (hc *ClientHttp) LoginRequestWithRedis(user string, password string) (*http.Response, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s:%v/authorizeduserredis", config.AuthHostname, config.AuthPort), nil)
	if err != nil {
		return nil, NewHttpReqCreationError("error creating request for allowed users on login")
	}
	req.SetBasicAuth(user, password)
	req.Header.Set("Content-Type", "application/json")
	return hc.client.Do(req)
}

func (hc *ClientHttp) LoginRequestWithJWT(user string, password string) (*http.Response, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s:%v/loginuserjwt", config.AuthHostname, config.AuthPort), nil)
	if err != nil {
		return nil, NewHttpReqCreationError("error creating request for allowed users on login")
	}
	req.SetBasicAuth(user, password)
	req.Header.Set("Content-Type", "application/json")
	return hc.client.Do(req)
}

func (hc *ClientHttp) AuthRequestWithJson(body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s:%v/authorizeduserjson", config.AuthHostname, config.AuthPort), body)
	if err != nil {
		return nil, NewHttpReqCreationError("error creating request for authorizantion json request")
	}
	req.Header.Set("Content-Type", "application/json")
	return hc.client.Do(req)
}

func (hc *ClientHttp) AuthRequestWithBasic(user string, password string) (*http.Response, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s:%v/authorizeduserbasic", config.AuthHostname, config.AuthPort), nil)
	if err != nil {
		return nil, NewHttpReqCreationError("error creating request for authorizantion basic request")
	}
	req.SetBasicAuth(user, password)
	req.Header.Set("Content-Type", "application/json")
	return hc.client.Do(req)
}

func (hc *ClientHttp) AuthRequestWithJwt(c *gin.Context) (*http.Response, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s:%v/authorizeduserjwt", config.AuthHostname, config.AuthPort), nil)
	if err != nil {
		return nil, NewHttpReqCreationError("error creating request for authorizantion jwt request")
	}
	req.Header.Set("Auth-Header", c.GetHeader("Auth-Header"))
	return hc.client.Do(req)
}
