package main

import (
	"fmt"
	"log"
	"os"

	"github.com/deild/myapp/commands"
	"github.com/urfave/cli"
)

var (
	version = "dev"
	date    = "unknown"
	commit  = "none"
)

func main() {
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Printf("%s %s\nBuild date: %s\nCommit: %s\n", c.App.Name, c.App.Version, date, commit)
	}

	app := cli.NewApp()
	app.Name = "Myapp"
	app.Version = version
	app.Author = "tolva <tolva@tuta.io>"
	app.Usage = "This is an myapp app written in Go"
	app.Action = func(c *cli.Context) error {
		err := commands.Foo()
		if err != nil {
			return err
		}
		fmt.Println("Done!")
		return nil
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
