package main

import (
	"fmt"
	"os"

	cli "gopkg.in/urfave/cli.v1"
)

// Dir will set the GOLNS_DIR environment variable to the new argument-passed subdirectory.
func Dir(cli *cli.Context) error {
	if cli.NArg() != 1 {
		fmt.Println(cli.Command.ArgsUsage)
	}

	os.Setenv(GolnsDirKey, cli.Args().First())

	return nil
}

// NewCommandDir appends the Dir function wrapped in a cli.Command struct.
func NewCommandDir(app *cli.App) {
	CommandDir := cli.Command{
		Name:      "dir",
		Aliases:   []string{"d"},
		Usage:     "Set the environment variable GOLNS_DIR to a given directory",
		ArgsUsage: "dir path-to-directory",
		Action:    Dir,
	}
	app.Commands = append(app.Commands, CommandDir)
}
