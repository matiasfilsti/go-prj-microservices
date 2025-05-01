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

func (s *Server) LoginJwt(c *gin.Context) {
	var rerr *ctrserrors.ReadBodyResponseError
	var jerr *ctrserrors.JsonImportError
	var herr *client.HttpReqCreationError

	user, password, ok := c.Request.BasicAuth()
	fmt.Println(user, password, ok)
	if !ok {
		ctrserrors.RespondHttpError(c, http.StatusInternalServerError, errors.New("basic Auth error"), "error trying to read user, password")
		return
	}

	validResponse, err := authenticateUserJwt(s, user, password)
	fmt.Println("respuesta: ", validResponse)
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

	setHeader(c, validResponse.JwtToken)
	c.JSON(200, "user successfully logged")
}

func setHeader(c *gin.Context, jwt string) {
	c.Header("Auth-Header", fmt.Sprintf("Bearer %s", jwt))
}

func authenticateUserJwt(s *Server, user, password string) (*models.ValidUserJwt, error) {
	var validResponse *models.ValidUserJwt

	resp, err := s.Hclient.LoginJwtRequest(user, password)
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
