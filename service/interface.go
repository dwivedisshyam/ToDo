package service

import (
	"github.com/dwivedisshyam/todo/model"
	"github.com/gin-gonic/gin"
)

type User interface {
	Create(u *model.User) (*model.User, error)
	Update(u *model.User) error
	List() (model.Users, error)
	Get(id int64) (*model.User, error)
	Delete(id int64) error

	Login(l *model.Login) (string, error)
	ValidateToken(ctx *gin.Context, tokenString string) (*model.Payload, error)
}

type Task interface {
	Create(u *model.Task) error
	Update(u *model.Task) error
	List() (model.Tasks, error)
	Get(id int64) (*model.Task, error)
}
