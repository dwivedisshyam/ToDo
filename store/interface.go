package store

import (
	"github.com/dwivedisshyam/todo/model"
)

type User interface {
	Create(u *model.User) error
	Update(u *model.User) error
	List() (model.Users, error)
	Get(id int64) (*model.User, error)
	GetByUsername(username string) (*model.User, error)
}

type Task interface {
	Create(*model.Task) error
	Update(u *model.Task) error
	List() (model.Tasks, error)
	Get(id int64) (*model.Task, error)
}
