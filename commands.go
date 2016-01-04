package main

import (
	"fmt"
	"os"

	"github.com/ariarijp/togoo/command"
	"github.com/codegangsta/cli"
)

// GlobalFlags is global option for all command.
var GlobalFlags = []cli.Flag{}

var forceFlags = []cli.Flag{
	cli.BoolFlag{
		Name:  "force, f",
		Usage: "Force initialize even if database already existed.",
	},
}

var allListFlags = []cli.Flag{
	cli.BoolFlag{
		Name:  "all, a",
		Usage: "List all tasks.",
	},
}

var allDeleteFlags = []cli.Flag{
	cli.BoolFlag{
		Name:  "all, a",
		Usage: "Delete all tasks.",
	},
}

// Commands are slice of command.
var Commands = []cli.Command{
	{
		Name:   "init",
		Usage:  "Initialize SQLite database",
		Action: command.CmdInit,
		Flags:  forceFlags,
	},
	{
		Name:    "add",
		Aliases: []string{"a"},
		Usage:   "Add task",
		Action:  command.CmdAdd,
		Flags:   []cli.Flag{},
	},
	{
		Name:    "list",
		Aliases: []string{"l"},
		Usage:   "Show all tasks",
		Action:  command.CmdList,
		Flags:   allListFlags,
	},
	{
		Name:    "done",
		Aliases: []string{"d"},
		Usage:   "Mark a task as done",
		Action:  command.CmdDone,
		Flags:   []cli.Flag{},
	},
	{
		Name:   "delete",
		Usage:  "Delete task",
		Action: command.CmdDelete,
		Flags:  allDeleteFlags,
	},
}

// CommandNotFound is custom error.
func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
