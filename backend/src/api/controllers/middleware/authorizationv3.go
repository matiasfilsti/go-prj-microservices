package middleware

import (
	"backend/src/api/client"
	"backend/src/api/controllers/ctrserrors"
	"backend/src/api/domain/models"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthorizedUserV3(s *client.ClientRedis) gin.HandlerFunc {

	return func(c *gin.Context) {
		user, err := getSessionAuth(c)
		if err != nil {
			ctrserrors.RespondHttpError(c, http.StatusInternalServerError, err, "error reading data token")
			c.Abort()
			return
		}

		redisUser, err := s.Get(c, user.Name)
		if err != nil {
			ctrserrors.RespondHttpError(c, http.StatusInternalServerError, err, "error: redis user not found")
			c.Abort()
			return
		}

		if redisUser.SessionToken != user.SessionToken || redisUser.CsrfToken != user.CSRFToken {
			ctrserrors.RespondHttpError(c, http.StatusUnauthorized, errors.New("not authorized"), "error: unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}

func getSessionAuth(c *gin.Context) (*models.UserV, error) {
	user := &models.UserV{}
	authUser, err := c.Cookie("auth-user")
	if err != nil {
		return user, err
	}

	sessionToken, err := c.Cookie("session_token")
	if err != nil {
		return user, err
	}

	csrfToken := c.GetHeader("csrf_token")
	if csrfToken == "" {
		return user, errors.New("no csrf token")
	}
	user.UpdateLoggedUserValues(fmt.Sprintf("user-%s", authUser), sessionToken, csrfToken)
	return user, nil

}
