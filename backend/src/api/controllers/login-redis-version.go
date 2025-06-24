package controllers

import (
	"backend/src/api/client"
	"backend/src/api/controllers/ctrserrors"
	"backend/src/api/domain/models"
	"encoding/json"
	"errors"
	"fmt"
	"io"
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
	// user, password, ok := c.Request.BasicAuth()
	// fmt.Println(user, password, ok)
	// if !ok {
	// 	ctrserrors.RespondHttpError(c, http.StatusInternalServerError, errors.New("basic Auth error"), "error trying to read user, password")
	// 	return
	// }

	validResponse, err := authenticateUser(s, user, password)
	if err != nil {
		switch {
		case errors.As(err, &rerr):
			ctrserrors.RespondHttpError(c, http.StatusInternalServerError, err, "reading response authentincation error")
			return
		case errors.As(err, &jerr):
			ctrserrors.RespondHttpError(c, http.StatusInternalServerError, err, "reading importing response authentincation error")
			return
		case errors.As(err, &herr):
			ctrserrors.RespondHttpError(c, http.StatusInternalServerError, err, "reading authentication request")
			return
		}
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
			ctrserrors.RespondHttpError(c, http.StatusInternalServerError, err, "saving user into redis error")
			return
		}
		ctrserrors.RespondHttpError(c, http.StatusInternalServerError, err, "redis failing - internal server error")
		return
	}
	setCookies(c, validResponse.SessionToken, validResponse.CSRFToken, user)

	c.JSON(200, "user successfully logged")
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
