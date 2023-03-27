package model

import (
	"time"

	"github.com/dwivedisshyam/go-lib/pkg/errors"
)

type Task struct {
	ID          int64      `json:"id,omitempty"`
	UserID      int64      `json:"user_id,omitempty"`
	Title       string     `json:"title,omitempty"`
	Description string     `json:"description"`
	DueDate     *time.Time `json:"due_date"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
}

type Tasks []Task

func (t *Task) Validate() error {
	if t.Title == "" {
		return errors.BadRequest("missing title")
	}

	if t.UserID <= 0 {
		return errors.BadRequest("missing user_id")
	}

	if t.DueDate == nil {
		return errors.BadRequest("missing due_date")
	}

	return nil
}
