package controllers

import (
	"backend/src/api/controllers/ctrserrors"
	"backend/src/api/domain/models"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) Login(c *gin.Context) {

	var validResponse models.ValidUser
	user, password, ok := c.Request.BasicAuth()
	fmt.Println(user, password, ok)
	if !ok {
		ctrserrors.RespondHttpError(c, http.StatusInternalServerError, errors.New("basic Auth error"), "error trying to read user, password")
		return
	}

	validResponse, err := authenticateUser(s, user, password)
	if err != nil {
		ctrserrors.RespondHttpError(c, http.StatusInternalServerError, err, "error trying to authenticate")
		return
	}

	if !validResponse.Valid {
		ctrserrors.RespondHttpError(c, http.StatusUnauthorized, errors.New("not allowed"), "User Not authorized")
		return
	}

	if err := SaveSessionUser(s, c, user, validResponse.SessionToken, validResponse.CSRFToken); err != nil {
		ctrserrors.RespondHttpError(c, http.StatusInternalServerError, err, "redis issue")
	}
	setCookies(c, validResponse.SessionToken, validResponse.CSRFToken, user)

	c.JSON(200, "user successfully logged")
}

func setCookies(c *gin.Context, session_token string, csrf_token string, user string) {
	c.SetCookie("session_token", session_token, 200, "/", "", false, true)
	c.SetCookie("csrf_token", csrf_token, 200, "/", "", false, false)
	c.SetCookie("auth-user", user, 200, "/", "", false, false)
}

func authenticateUser(s *Server, user, password string) (models.ValidUser, error) {
	var validResponse models.ValidUser

	resp, err := s.Hclient.LoginRequest(user, password)
	fmt.Println(resp, err)
	if err != nil {
		return validResponse, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	fmt.Println(string(respBody), err)
	if err != nil {
		return validResponse, err
	}

	err = json.Unmarshal(respBody, &validResponse)
	fmt.Println(err)
	if err != nil {
		return validResponse, err
	}

	return validResponse, nil
}

func SaveSessionUser(s *Server, c *gin.Context, user string, session_token string, crsf_token string) error {
	if err := s.RdsClient.Set(c, fmt.Sprintf("user-%s", user), session_token, crsf_token); err != nil {
		return err
	}
	return nil
}

// func (s *Server) Login(c *gin.Context) {
// 	var loginUser models.LoginUser
// 	user, password, ok := c.Request.BasicAuth()
// 	fmt.Println(user, password, ok)
// 	if !ok {
// 		ctrserrors.RespondHttpError(c, http.StatusInternalServerError, errors.New("basic Auth error"), "error trying to read user, password")
// 		return
// 	}

// 	resp, err := s.Hclient.DoLoginReq(user)
// 	fmt.Println(resp, err)
// 	if err != nil {
// 		ctrserrors.RespondHttpError(c, http.StatusInternalServerError, err, "Error trying to authenticate")
// 		return
// 	}
// 	defer resp.Body.Close()

// 	respBody, _ := io.ReadAll(resp.Body)
// 	fmt.Println(respBody, resp)
// 	err = json.Unmarshal(respBody, &loginUser)
// 	if err != nil {
// 		ctrserrors.RespondHttpError(c, http.StatusInternalServerError, err, "Error reading response data")
// 		return
// 	}
// 	userValidation()
// }

// func userValidation(password string, dbUser models.LoginUser) {
// 	hashedPassword, err := hashPassword(password)
// 	if hashedPassword != dbUser.Password {
// 		return false
// 	}
// 	return true
// }

// func getUserAuth(user string) (*http.Response, error) {
// 	httpclient := client.CreateHttpClient()
// 	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s:%v/get/%s", config.AuthHostname, config.AuthPort, user), nil)
// 	if err != nil {
// 		fmt.Println("Error creando la solicitud:", err)
// 		return nil, err
// 	}
// 	return httpclient.Do(req)
// }

// c.SetCookie(
// 	"session_token",
// 	validResponse.SessionToken,
// 	200,
// 	"/",
// 	"",
// 	false,
// 	true,
// )
// c.SetCookie(
// 	"csrf_token",
// 	validResponse.SessionToken,
// 	200,
// 	"/",
// 	"",
// 	false,
// 	false,
// )

// resp, err := s.Hclient.LoginRequest(user, password)
// if err != nil {
// 	ctrserrors.RespondHttpError(c, http.StatusInternalServerError, err, "Error trying to authenticate")
// 	return
// }
// defer resp.Body.Close()

// respBody, _ := io.ReadAll(resp.Body)
// fmt.Println(string(respBody), resp)
// err = json.Unmarshal(respBody, &validResponse)
// if err != nil {
// 	ctrserrors.RespondHttpError(c, http.StatusInternalServerError, err, "Error reading response data")
// 	return
// }
