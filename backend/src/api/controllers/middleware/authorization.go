package middleware

import (
	"backend/src/api/client"
	"backend/src/api/config"
	"backend/src/api/controllers/ctrserrors"
	"backend/src/api/domain/models"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthorizedUser() gin.HandlerFunc {
	fmt.Println("inicio de http client")
	httpclient := client.CreateHttpClient()

	return func(c *gin.Context) {
		var user models.User
		var validResponse ValidUser

		if err := c.ShouldBindJSON(&user); err != nil {
			ctrserrors.RespondHttpError(c, http.StatusBadRequest, err, "Invalid Body Data")
			c.Abort()
			return
		}
		// before request
		body, err := encodeStructToJson(user)
		if err != nil {
			ctrserrors.RespondHttpError(c, http.StatusInternalServerError, err, "Error trying encode body data")
			c.Abort()
			return
		}
		resp, err := athznRequest(httpclient, body)
		if err != nil {
			ctrserrors.RespondHttpError(c, http.StatusInternalServerError, err, "Error trying to authenticate")
			c.Abort()
			return
		}
		defer resp.Body.Close()

		respBody, _ := io.ReadAll(resp.Body)
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

func encodeStructToJson(data models.User) (io.Reader, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(jsonData), nil
}

func athznRequest(httpclient *http.Client, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s:%v/allowedusers", config.AuthHostname, config.AuthPort), body)
	if err != nil {
		fmt.Println("Error creando la solicitud:", err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return httpclient.Do(req)
}
