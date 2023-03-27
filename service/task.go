package service

import (
	"time"

	"github.com/dwivedisshyam/todo/model"

	"github.com/dwivedisshyam/todo/store"
)

type taskService struct {
	store store.Task
}

func NewTask(st store.Task) Task {
	return &taskService{store: st}
}

func (ts *taskService) Create(t *model.Task) error {
	err := t.Validate()
	if err != nil {
		return err
	}

	currentTime := time.Now()
	t.CreatedAt = &currentTime

	err = ts.store.Create(t)
	if err != nil {
		return err
	}

	return nil
}

func (ts *taskService) Update(t *model.Task) error {
	err := t.Validate()
	if err != nil {
		return err
	}

	err = ts.store.Update(t)
	if err != nil {
		return err
	}

	return nil
}

func (ts *taskService) Get(id int64) (*model.Task, error) {
	task, err := ts.store.Get(id)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (ts *taskService) List() (model.Tasks, error) {
	tasks, err := ts.store.List()
	if err != nil {
		return nil, err
	}

	return tasks, nil
}
