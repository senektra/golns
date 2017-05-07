package main

import (
	"fmt"
	"os"

	cli "gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "golns"
	app.Usage = "Links symbolically a directory from $GOPATH/src to a sub directory in $GOLNS_DIR"
	app.UsageText = "golns [global-options] (command [command-options] | " +
		"target-sub-directory gopath-src-directory)"

	NewCommandDir(app)

	app.Action = func(c *cli.Context) error {
		if c.NArg() != 2 {
			fmt.Println(c.App.Usage)
		}
		return nil
	}

	app.Run(os.Args)
}
