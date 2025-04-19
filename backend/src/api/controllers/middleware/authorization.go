package middleware

import (
	"backend/src/api/client"
	"backend/src/api/controllers/ctrserrors"
	"backend/src/api/domain/models"
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthorizedUser(s *client.ClientHttp) gin.HandlerFunc {

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

		resp, err := s.AthznRequestV1(body)
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
