package command

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/codegangsta/cli"
)

// CmdDone mark task as finished.
func CmdDone(c *cli.Context) {

	if len(c.Args()) != 1 {
		return
	}

	id := c.Args()[0]

	db, err := sql.Open("sqlite3", dbPath())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("UPDATE todos SET is_done = 1 WHERE id = " + id)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Task %s is done\n", id)
}
