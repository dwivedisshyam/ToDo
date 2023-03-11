package user

import (
	"time"

	"github.com/dwivedisshyam/go-lib/pkg/errors"
	"github.com/dwivedisshyam/todo/model"
	"github.com/dwivedisshyam/todo/service"
	"github.com/dwivedisshyam/todo/store"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	store store.User
}

func New(st store.User) service.User {
	return &userService{store: st}
}

func (us *userService) Create(u *model.User) error {
	err := u.Validate()
	if err != nil {
		return err
	}

	pswd, hashErr := HashPassword(u.Password)
	if hashErr != nil {
		return errors.Unauthenticated(hashErr.Error())
	}

	u.Password = pswd
	currentTime := time.Now()

	u.CreatedAt = &currentTime

	err = us.store.Create(u)
	if err != nil {
		return err
	}

	return nil
}

func (us *userService) Update(u *model.User) error {
	err := u.Validate()
	if err != nil {
		return err
	}

	err = us.store.Update(u)
	if err != nil {
		return err
	}

	return nil
}

func (us *userService) Get(id int64) (*model.User, error) {
	user, err := us.store.Get(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (us *userService) List() (model.Users, error) {
	user, err := us.store.List()
	if err != nil {
		return nil, err
	}

	return user, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
