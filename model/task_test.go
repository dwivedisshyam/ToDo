package model

import (
	"testing"

	"github.com/dwivedisshyam/go-lib/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func Test_TaskValidate(t *testing.T) {
	tcs := []struct {
		name string
		task *Task
		exp  error
	}{
		{
			"Failure",
			&Task{},
			errors.BadRequest("missing title"),
		},
		{
			"Success",
			&Task{Title: "MyTask"},
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
