package model

import (
	"time"

	"github.com/dwivedisshyam/go-lib/pkg/errors"
)

type User struct {
	ID        int64      `json:"id,omitempty"`
	Username  string     `json:"username,omitempty"`
	Password  string     `json:"password,omitempty"`
	FullName  string     `json:"full_name,omitempty"`
	Email     string     `json:"email,omitempty"`
	Role      Role       `json:"role,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
}

type Users []User

func (u *User) ValidateCreate() error {
	if u.Username == "" {
		return errors.BadRequest("missing username")
	}

	if u.Password == "" {
		return errors.BadRequest("missing password")
	}

	if u.FullName == "" {
		return errors.BadRequest("missing full_name")
	}

	if u.Email == "" {
		return errors.BadRequest("missing email")
	}

	if u.Role == "" {
		return errors.BadRequest("missing role")
	}

	return nil
}

func (u *User) ValidateUpdate() error {
	if u.FullName == "" {
		return errors.BadRequest("missing full_name")
	}

	if u.Email == "" {
		return errors.BadRequest("missing email")
	}

	return nil
}
