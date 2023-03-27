package user

import (
	"database/sql"
	"fmt"

	"github.com/dwivedisshyam/go-lib/pkg/errors"
	"github.com/dwivedisshyam/todo/model"
	"github.com/dwivedisshyam/todo/store"
)

const (
	createQuery      = `INSERT INTO "user" (username, password, full_name, email, role, created_at) VALUES ($1, $2, $3, $4, $5, $6)`
	getQuery         = `SELECT username, full_name, email, role, created_at from "user" WHERE id=$1`
	getAllQuery      = `SELECT id, username, full_name, email, role, created_at from "user"`
	updateQuery      = `UPDATE "user" SET full_name=$1, email=$2 WHERE id=$3`
	deleteQuery      = `DELETE FROM "user" WHERE id=$1`
	loginSelectQuery = `SELECT id, username, email, password FROM "user" WHERE username=$1`
)

type userStore struct {
	db *sql.DB
}

func New(db *sql.DB) store.User {
	return &userStore{db: db}
}

func (us *userStore) Create(u *model.User) (int64, error) {
	tx, err := us.db.Begin()
	if err != nil {
		return 0, errors.Unexpected(err.Error())
	}

	_, err = tx.Exec(createQuery, u.Username, u.Password, u.FullName, u.Email, u.Role, u.CreatedAt)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return 0, errors.Unexpected(fmt.Sprintf("err: %v, rbErr: %v", err, rbErr))
		}
		return 0, errors.Unexpected(err.Error())
	}

	var id int64

	err = tx.QueryRow(`SELECT MAX(id) from "user"`).Scan(&id)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return 0, errors.Unexpected(fmt.Sprintf("err: %v, rbErr: %v", err, rbErr))
		}
		return 0, errors.Unexpected(err.Error())
	}

	err = tx.Commit()
	if err != nil {
		return 0, errors.Unexpected(err.Error())
	}

	return id, nil
}

func (us *userStore) Update(u *model.User) error {
	_, err := us.db.Exec(updateQuery, u.FullName, u.Email, u.ID)
	if err != nil {
		return errors.Unexpected(err.Error())
	}

	return nil
}

func (us *userStore) Delete(id int64) error {
	_, err := us.db.Exec(deleteQuery, id)
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
	err := us.db.QueryRow(loginSelectQuery, username).
		Scan(&u.ID, &u.Username, &u.Email, &u.Password)
	if err != nil {
		return nil, errors.Unexpected(err.Error())
	}

	return &u, nil
}
