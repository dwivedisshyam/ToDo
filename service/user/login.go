package user

import (
	"fmt"
	"time"

	"github.com/dwivedisshyam/go-lib/pkg/errors"
	"github.com/dwivedisshyam/todo/model"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var hmacSampleSecret = "sample_secret"

func (us *userService) Login(l *model.Login) (string, error) {
	if err := l.Validate(); err != nil {
		return "", err
	}

	user, err := us.store.GetByUsername(l.Username)
	if err != nil {
		return "", errors.Unauthenticated("invalid credentials")
	}

	if !CheckPasswordHash(l.Password, user.Password) {
		return "", errors.Unauthenticated("invalid credentials")
	}

	user.Password = ""

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, model.Payload{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
		User: user,
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, tkErr := token.SignedString([]byte(hmacSampleSecret))
	if tkErr != nil {
		return "", errors.Unexpected(tkErr.Error())
	}

	return tokenString, nil
}

func (us *userService) ValidateToken(ctx *gin.Context, tokenString string) (*model.Payload, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(hmacSampleSecret), nil
	})

	if err != nil {
		return nil, errors.Unauthenticated("invalid credentials")
	}

	claims, ok := token.Claims.(model.Payload)
	if !ok || !token.Valid {
		return nil, errors.Unauthenticated("invalid credentials")
	}

	return &claims, nil
}
