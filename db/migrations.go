package db

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

func migration(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS migration(
		id serial4 NOT NULL,
		app varchar(30) NOT NULL,
    	sql_file varchar(30) NOT NULL,
    	created_at timestamptz NOT NULL
	);`)

	if err != nil {
		return err
	}

	files, err := ioutil.ReadDir("./db/scripts")
	if err != nil {

		return err
	}

	sqlFiles := make(map[string]string)

	for _, file := range files {
		name := strings.Trim(strings.Trim(file.Name(), "db_"), ".sql")
		var exist bool

		err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM migration WHERE sql_file=$1 and app=$2)", name, "todo").Scan(&exist)
		if err != nil {
			return err
		}

		if !exist {
			sqlFiles[name] = "./db/scripts/" + file.Name()
		}
	}

	return runMigration(db, sqlFiles)
}

func runMigration(db *sql.DB, sqlFiles map[string]string) error {
	var success []string
	for m, f := range sqlFiles {
		file, err := ioutil.ReadFile(f)
		if err != nil {
			return err
		}

		if _, err := db.Exec(string(file)); err != nil {
			return err
		}

		_, err = db.Exec("INSERT INTO migration (app,sql_file,created_at) values($1,$2,$3)", "todo", m, time.Now())
		if err != nil {
			return err
		}

		success = append(success, m)
	}

	fmt.Printf("migration successul for %v\n", success)

	return nil
}
