package model

import (
	"github.com/dwivedisshyam/go-lib/pkg/errors"
	"github.com/golang-jwt/jwt/v4"
)

type Login struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

func (l *Login) Validate() error {
	if l.Username == "" {
		return errors.BadRequest("missing user_name")
	}

	if l.Password == "" {
		return errors.BadRequest("missing password")
	}

	return nil
}

type Payload struct {
	jwt.RegisteredClaims
	User User `json:"user"`
}
