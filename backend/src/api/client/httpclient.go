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

// func (hc *ClientHttp) DoLoginReq(user string) (*http.Response, error) {
// 	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s:%v/get/%s", config.AuthHostname, config.AuthPort, user), nil)
// 	if err != nil {
// 		return nil, NewHttpReqCreationError("error creating request for login")
// 	}
// 	return hc.client.Do(req)

// }

func (hc *ClientHttp) LoginRequest(user string, password string) (*http.Response, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s:%v/allowedusersLogin", config.AuthHostname, config.AuthPort), nil)
	if err != nil {
		return nil, NewHttpReqCreationError("error creating request for allowed users on login")
	}
	req.SetBasicAuth(user, password)
	req.Header.Set("Content-Type", "application/json")
	return hc.client.Do(req)
}

func (hc *ClientHttp) LoginJwtRequest(user string, password string) (*http.Response, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s:%v/allowuserwithjwt", config.AuthHostname, config.AuthPort), nil)
	if err != nil {
		return nil, NewHttpReqCreationError("error creating request for allowed users on login")
	}
	req.SetBasicAuth(user, password)
	req.Header.Set("Content-Type", "application/json")
	return hc.client.Do(req)
}

func (hc *ClientHttp) AthznRequestV1(body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s:%v/allowedusers", config.AuthHostname, config.AuthPort), body)
	if err != nil {
		return nil, NewHttpReqCreationError("error creating request for authorizantion V1 request")
	}
	req.Header.Set("Content-Type", "application/json")
	return hc.client.Do(req)
}

func (hc *ClientHttp) AthznRequestV2(user string, password string) (*http.Response, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s:%v/allowedusersV2", config.AuthHostname, config.AuthPort), nil)
	if err != nil {
		return nil, NewHttpReqCreationError("error creating request for authorizantion V2 request")
	}
	req.SetBasicAuth(user, password)
	req.Header.Set("Content-Type", "application/json")
	return hc.client.Do(req)
}

func (hc *ClientHttp) AthznRequestJwt(c *gin.Context) (*http.Response, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s:%v/validateuserjwt", config.AuthHostname, config.AuthPort), nil)
	if err != nil {
		return nil, NewHttpReqCreationError("error creating request for authorizantion jwt request")
	}
	req.Header.Set("Auth-Header", c.GetHeader("Auth-Header"))
	return hc.client.Do(req)
}
