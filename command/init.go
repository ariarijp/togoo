package command

import (
	"database/sql"
	"log"
	"os"
	"path"

	"github.com/codegangsta/cli"
)

// CmdInit create database and tables.
func CmdInit(c *cli.Context) {

	dbPath := dbPath()
	isDatabaseExists := exists(dbPath)
	isForceMode := c.String("force")

	if isForceMode == "false" && isDatabaseExists {
		println("Database is already existed")
		return
	}

	if isForceMode == "true" {
		os.Remove(dbPath)
	}

	togooDir := path.Dir(dbPath)
	if !exists(togooDir) {
		if err := os.Mkdir(togooDir, 0700); err != nil {
			println("Cannot create " + togooDir)
			return
		}
	}

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `
		CREATE TABLE todos (
			id integer NOT NULL PRIMARY KEY,
			title text NOT NULL,
			is_done integer NOT NULL,
			created_at integer NOT NULL,
			updated_at integer NOT NULL
		);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}

	println(dbPath + " created.")
}
