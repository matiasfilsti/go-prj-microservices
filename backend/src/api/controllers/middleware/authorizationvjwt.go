package middleware

import (
	"backend/src/api/client"
	"backend/src/api/controllers/ctrserrors"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthorizedUserJWT(s *client.ClientHttp) gin.HandlerFunc {

	return func(c *gin.Context) {
		var validResponse ValidUser
		resp, err := s.AthznRequestJwt(c)
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
