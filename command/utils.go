package command

import (
	"os"

	"github.com/mitchellh/go-homedir"
)

func dbPath() string {
	home, err := homedir.Dir()
	if err != nil {
		panic("Home directory does not exist.")
	}

	return home + "/.togoo/togoo.db"
}

func exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
