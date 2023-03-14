package task

import (
	"database/sql"

	"github.com/dwivedisshyam/go-lib/pkg/errors"
	"github.com/dwivedisshyam/todo/model"
	"github.com/dwivedisshyam/todo/store"
)

const (
	createQuery = `INSERT INTO "task" (user_id, title, description, due_date, created_at) VALUES ($1, $2, $3, $4, $5)`
	getQuery    = `SELECT user_id, title, description, due_date, created_at from "task" WHERE id=$1`
	getAllQuery = `SELECT id, user_id, title, description, due_date, created_at from "task"`
	updateQuery = `UPDATE "task" SET title=$1, description=$2, due_date=$3 WHERE id=$4`
)

type taskStore struct {
	db *sql.DB
}

func New(db *sql.DB) store.Task {
	return &taskStore{db: db}
}

func (us *taskStore) Create(t *model.Task) error {
	_, err := us.db.Exec(createQuery, t.UserID, t.Title, t.Description, t.DueDate, t.CreatedAt)
	if err != nil {
		return errors.Unexpected(err.Error())
	}

	return nil
}

func (us *taskStore) Update(t *model.Task) error {
	_, err := us.db.Exec(updateQuery, t.Title, t.Description, t.DueDate, t.ID)
	if err != nil {
		return errors.Unexpected(err.Error())
	}

	return nil
}

func (us *taskStore) List() (model.Tasks, error) {
	rows, err := us.db.Query(getAllQuery)
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

	err := us.db.QueryRow(getQuery, id).Scan(&t.UserID, &t.Title, &t.Description, &t.DueDate, &t.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, errors.NotFound("task not found")
	}

	if err != nil {
		return nil, errors.Unexpected(err.Error())
	}

	return &t, nil
}
