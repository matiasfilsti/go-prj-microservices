package services

import (
	"authentication/src/api/config"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type MyCustomClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func (s *UserService) GenerateJwtToken(username string) (string, error) {
	claims := MyCustomClaims{
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "test",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	fmt.Printf("token claims added: %v \n", claims)
	tokenString, err := token.SignedString([]byte(config.SecretJwtKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (s *UserService) ValidateJwtToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.SecretJwtKey), nil
	})
	if err != nil {
		return nil, errors.New("error parsing jwt")
	}
	if !token.Valid {
		return nil, errors.New("jwt token invalid")
	}
	return token, nil
}
