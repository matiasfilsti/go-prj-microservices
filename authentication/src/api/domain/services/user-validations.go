package services

import (
	"authentication/src/api/domain/errors"
	"authentication/src/api/domain/models"
	"regexp"
)

var namePattern = regexp.MustCompile("^[A-Za-z]{7,}$")
var passwordPattern = regexp.MustCompile("^[A-Za-z0-9@$!%*?&]{7,}$")

func userValidate(user models.User) error {
	if !validateName(user.Name) {
		return errors.NewInputError("Error Validando nombre")
	}
	if !validatePassword(user.Password) {
		return errors.NewInputError("Error Validando password")
	}
	return nil
}

func validatePassword(password string) bool {
	return passwordPattern.MatchString(password)
}

func validateName(name string) bool {
	return namePattern.MatchString(name)
}
