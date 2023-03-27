package model

import (
	"testing"
	"time"

	"github.com/dwivedisshyam/go-lib/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func Test_TaskValidate(t *testing.T) {
	dueDate := time.Now().Add(24 * time.Hour)
	tcs := []struct {
		name string
		task *Task
		exp  error
	}{
		{
			"UserIDMissing",
			&Task{Title: "MyTask", DueDate: &dueDate},
			errors.BadRequest("missing user_id"),
		},
		{
			"DueDateMissing",
			&Task{Title: "MyTask", UserID: 34},
			errors.BadRequest("missing due_date"),
		},
		{
			"TitleMissing",
			&Task{},
			errors.BadRequest("missing title"),
		},
		{
			"Success",
			&Task{Title: "MyTask", UserID: 34, DueDate: &dueDate},
			nil,
		},
	}

	for i := range tcs {
		tc := tcs[i]

		t.Run(tc.name, func(t *testing.T) {
			actual := tc.task.Validate()

			assert.Equal(t, tc.exp, actual)
		})
	}
}
