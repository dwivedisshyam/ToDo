package user

import (
	"database/sql"

	"github.com/dwivedisshyam/go-lib/pkg/errors"
	"github.com/dwivedisshyam/todo/model"
	"github.com/dwivedisshyam/todo/store"
)

const (
	createQuery = `INSERT INTO "user" (username, password, full_name, email, role, created_at) VALUES ($1, $2, $3, $4, $5, $6)`
	getQuery    = `SELECT username, full_name, email, role, created_at from "user" WHERE id=$1`
	getAllQuery = `SELECT id, username, full_name, email, role, created_at from "user"`
	updateQuery = `UPDATE "user" SET username=$1, full_name=$2, role=$3 WHERE id=$4`
)

type userStore struct {
	db *sql.DB
}

func New(db *sql.DB) store.User {
	return &userStore{db: db}
}

func (us *userStore) Create(u *model.User) error {
	_, err := us.db.Exec(createQuery, u.Username, u.Password, u.FullName, u.Email, u.Role, u.CreatedAt)
	if err != nil {
		return errors.Unexpected(err.Error())
	}

	return nil
}

func (us *userStore) Update(u *model.User) error {
	_, err := us.db.Exec(updateQuery, u.Username, u.FullName, u.Role, u.ID)
	if err != nil {
		return errors.Unexpected(err.Error())
	}

	return nil
}

func (us *userStore) List() (model.Users, error) {
	rows, err := us.db.Query(getAllQuery)
	if err != nil {
		return nil, errors.Unexpected(err.Error())
	}

	var users model.Users

	for rows.Next() {
		var u model.User

		err = rows.Scan(&u.ID, &u.Username, &u.FullName, &u.Email, &u.Role, &u.CreatedAt)
		if err != nil {
			return nil, errors.Unexpected(err.Error())
		}

		users = append(users, u)
	}

	return users, nil
}

func (us *userStore) Get(id int64) (*model.User, error) {
	var u model.User

	err := us.db.QueryRow(getQuery, id).Scan(&u.Username, &u.FullName, &u.Email, &u.Role, &u.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, errors.NotFound("user not found")
	}

	if err != nil {
		return nil, errors.Unexpected(err.Error())
	}

	return &u, nil
}

func (us *userStore) GetByUsername(username string) (*model.User, error) {
	var u model.User
	err := us.db.QueryRow(`SELECT id, username, email, password FROM "Users" WHERE username=$1`, username).
		Scan(&u.ID, &u.Username, &u.Email, &u.Password)
	if err != nil {
		return nil, errors.Unexpected(err.Error())
	}

	return &u, nil
}
