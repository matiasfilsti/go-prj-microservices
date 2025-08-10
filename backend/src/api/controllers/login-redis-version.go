package controllers

import (
	"backend/src/api/client"
	"backend/src/api/controllers/ctrserrors"
	"backend/src/api/domain/models"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) LoginWithRedis(c *gin.Context) {
	var rerr *ctrserrors.ReadBodyResponseError
	var jerr *ctrserrors.JsonImportError
	var herr *client.HttpReqCreationError
	var serr *client.RedisUserSaveError

	fmt.Print("LoginWithRedis called\n")
	user, password := c.Request.FormValue("user"), c.Request.FormValue("password")
	if user == "" || password == "" {
		ctrserrors.RespondHttpError(c, http.StatusBadRequest, errors.New("user or password not provided"), "error trying to read user, password")
		return
	}

	validResponse, err := authenticateUser(s, user, password)
	if err != nil {
		switch {
		case errors.As(err, &rerr):
			log.Println("error reading response authentincation error:", err)
			ctrserrors.RespondHttpError(c, http.StatusInternalServerError, err, "reading response authentincation error")
			return
		case errors.As(err, &jerr):
			log.Println("error reading importing response authentincation error:", err)
			ctrserrors.RespondHttpError(c, http.StatusInternalServerError, err, "reading importing response authentincation error")
			return
		case errors.As(err, &herr):
			log.Println("error reading authentication request:", err)
			ctrserrors.RespondHttpError(c, http.StatusInternalServerError, err, "reading authentication request")
			return
		}
		log.Println("error trying to authenticate:", err)
		ctrserrors.RespondHttpError(c, http.StatusInternalServerError, err, "error trying to authenticate")
		return
	}

	if !validResponse.Valid {
		ctrserrors.RespondHttpError(c, http.StatusUnauthorized, errors.New("not allowed"), "User Not authorized")
		return
	}

	err = SaveSessionUser(s, c, user, validResponse.SessionToken, validResponse.CSRFToken)
	if err != nil {
		switch {
		case errors.As(err, &serr):
			log.Println("error saving user into redis:", err)
			ctrserrors.RespondHttpError(c, http.StatusInternalServerError, err, "saving user into redis error")
			return
		}
		log.Println("error saving user into redis:", err)
		ctrserrors.RespondHttpError(c, http.StatusInternalServerError, err, "redis failing - internal server error")
		return
	}
	setCookies(c, validResponse.SessionToken, validResponse.CSRFToken, user)

	c.JSON(200, "user successfully logged")
	// c.JSON(200, gin.H{
	// 	"message": "user successfully logged",
	// 	"tokens": gin.H{
	// 		"session_token": validResponse.SessionToken,
	// 		"csrf_token":    validResponse.CSRFToken,
	// 		"auth_user":     user,
	// 	},
	// })
}

func setCookies(c *gin.Context, session_token string, csrf_token string, user string) {
	c.SetCookie("session_token", session_token, 200, "/", "", false, true)
	c.SetCookie("csrf_token", csrf_token, 200, "/", "", false, false)
	c.SetCookie("auth-user", user, 200, "/", "", false, false)
}

func authenticateUser(s *Server, user, password string) (*models.ValidUser, error) {
	var validResponse *models.ValidUser

	resp, err := s.Hclient.LoginRequestWithRedis(user, password)
	if err != nil {
		return validResponse, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return validResponse, ctrserrors.NewReadBodyResponseError("error reading response body")
	}

	err = json.Unmarshal(respBody, &validResponse)
	if err != nil {
		return validResponse, ctrserrors.NewJsonImportError("error importing json into model")
	}

	return validResponse, nil
}

func SaveSessionUser(s *Server, c *gin.Context, user string, session_token string, crsf_token string) error {
	if err := s.RdsClient.Set(c, fmt.Sprintf("user-%s", user), session_token, crsf_token); err != nil {
		return err
	}
	return nil
}

func (s *Server) AuthorizeRedisLogin(c *gin.Context) {
	log.Println("AuthorizeRedisLogin called")
	user, err := getSessionConfig(c)
	if err != nil {
		log.Println("error reading session config:", err)
		ctrserrors.RespondHttpError(c, http.StatusInternalServerError, err, "error reading data token")
		return
	}

	redisUser, err := s.RdsClient.Get(c, user.Name)
	if err != nil {
		log.Println("error getting user from redis:", err)
		ctrserrors.RespondHttpError(c, http.StatusInternalServerError, err, "error: redis user not found")
		return
	}

	if redisUser.SessionToken != user.SessionToken || redisUser.CsrfToken != user.CSRFToken {
		ctrserrors.RespondHttpError(c, http.StatusUnauthorized, errors.New("not authorized"), "error: unauthorized")
		return
	}
	c.JSON(200, "user authorized successfully")
}

func getSessionConfig(c *gin.Context) (*models.UserValidation, error) {
	user := &models.UserValidation{}
	authUser, err := c.Cookie("auth-user")
	if err != nil {
		return user, err
	}

	sessionToken, err := c.Cookie("session_token")
	if err != nil {
		return user, err
	}

	csrfToken := c.GetHeader("X-CSRF-Token")
	if csrfToken == "" {
		return user, errors.New("no csrf token")
	}
	user.UpdateLoggedUserValues(fmt.Sprintf("user-%s", authUser), sessionToken, csrfToken)
	return user, nil

}

func (s *Server) LogoutHandlerRedis(c *gin.Context) {
	log.Println("LogoutHandlerRedis called")
	authUser, err := c.Cookie("auth-user")
	if err == nil && authUser != "" {
		err = s.RdsClient.Del(c, fmt.Sprintf("user-%s", authUser))
		if err != nil {
			log.Println("error deleting user from redis:", err)
			ctrserrors.RespondHttpError(c, http.StatusInternalServerError, err, "error: redis user not found")
			return
		}
	}
	log.Println("deleting cookies")
	c.SetCookie("auth-user", "", -1, "/", "", false, true)
	c.SetCookie("session_token", "", -1, "/", "", false, true)
	c.SetCookie("csrf_token", "", -1, "/", "", false, true)
	log.Println("redirecting to login")
	c.JSON(200, "logout successfully")
}
