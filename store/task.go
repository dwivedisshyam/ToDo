package store

import (
	"database/sql"

	"github.com/dwivedisshyam/go-lib/pkg/errors"
	"github.com/dwivedisshyam/todo/model"
)

type taskStore struct {
	db *sql.DB
}

func NewTask(db *sql.DB) Task {
	return &taskStore{db: db}
}

func (us *taskStore) Create(t *model.Task) error {
	_, err := us.db.Exec(createTaskQ, t.UserID, t.Title, t.Description, t.DueDate, t.CreatedAt)
	if err != nil {
		return errors.Unexpected(err.Error())
	}

	return nil
}

func (us *taskStore) Update(t *model.Task) error {
	_, err := us.db.Exec(updateTaskQ, t.Title, t.Description, t.DueDate, t.ID)
	if err != nil {
		return errors.Unexpected(err.Error())
	}

	return nil
}

func (us *taskStore) List() (model.Tasks, error) {
	rows, err := us.db.Query(getAllTaskQ)
	if err != nil {
		return nil, errors.Unexpected(err.Error())
	}

	var tasks model.Tasks

	for rows.Next() {
		var t model.Task

		err = rows.Scan(&t.ID, &t.UserID, &t.Title, &t.Description, &t.DueDate, &t.CreatedAt)
		if err != nil {
			return nil, errors.Unexpected(err.Error())
		}

		tasks = append(tasks, t)
	}

	return tasks, nil
}

func (us *taskStore) Get(id int64) (*model.Task, error) {
	var t model.Task

	err := us.db.QueryRow(getTaskQ, id).Scan(&t.UserID, &t.Title, &t.Description, &t.DueDate, &t.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, errors.NotFound("task not found")
	}

	if err != nil {
		return nil, errors.Unexpected(err.Error())
	}

	return &t, nil
}
