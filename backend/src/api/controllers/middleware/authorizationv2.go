package middleware

import (
	"backend/src/api/client"
	"backend/src/api/controllers/ctrserrors"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthorizedUserV2(s *client.ClientHttp) gin.HandlerFunc {

	return func(c *gin.Context) {
		var validResponse ValidUser

		user, password, ok := c.Request.BasicAuth()
		fmt.Println(user, password, ok)
		if !ok {
			ctrserrors.RespondHttpError(c, http.StatusInternalServerError, errors.New("basic Auth error"), "error trying to read user, password")
			c.Abort()
			return
		}

		resp, err := s.AthznRequestV2(user, password)
		if err != nil {
			ctrserrors.RespondHttpError(c, http.StatusInternalServerError, err, "Error trying to authenticate")
			c.Abort()
			return
		}
		defer resp.Body.Close()

		respBody, _ := io.ReadAll(resp.Body)
		fmt.Println(respBody, resp)
		err = json.Unmarshal(respBody, &validResponse)
		if err != nil {
			ctrserrors.RespondHttpError(c, http.StatusInternalServerError, err, "Error reading response data")
			c.Abort()
			return
		}
		if !validResponse.Valid {
			ctrserrors.RespondHttpError(c, http.StatusUnauthorized, errors.New("not permitted"), "User Not authorized")
			c.Abort()
			return

		}
		c.Next()
	}
}

// func athznRequestV2(httpclient *http.Client, user string, password string) (*http.Response, error) {

// 	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s:%v/allowedusersV2", config.AuthHostname, config.AuthPort), nil)
// 	if err != nil {
// 		fmt.Println("Error creando la solicitud:", err)
// 		return nil, err
// 	}
// 	req.SetBasicAuth(user, password)
// 	req.Header.Set("Content-Type", "application/json")
// 	return httpclient.Do(req)
// }
