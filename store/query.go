package store

const (
	// --------------------------- TASK QUERIES ---------------------------------
	createTaskQ = `INSERT INTO "task" (user_id, title, description, due_date, created_at) VALUES ($1, $2, $3, $4, $5)`
	getTaskQ    = `SELECT user_id, title, description, due_date, created_at from "task" WHERE id=$1`
	getAllTaskQ = `SELECT id, user_id, title, description, due_date, created_at from "task"`
	updateTaskQ = `UPDATE "task" SET title=$1, description=$2, due_date=$3 WHERE id=$4`

	// --------------------------- USER QUERIES ---------------------------------
	createUserQ  = `INSERT INTO "user" (username, password, full_name, email, role, created_at) VALUES ($1, $2, $3, $4, $5, $6)`
	getUserQ     = `SELECT username, full_name, email, role, created_at from "user" WHERE id=$1`
	getAllUserQ  = `SELECT id, username, full_name, email, role, created_at from "user"`
	updateUserQ  = `UPDATE "user" SET full_name=$1, email=$2 WHERE id=$3`
	deleteUserQ  = `DELETE FROM "user" WHERE id=$1`
	loginSelectQ = `SELECT id, username, email, password FROM "user" WHERE username=$1`
)
