package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = "zippy"
	app.Author = "Tugberk Ugurlu"
	app.Usage = "General purpose package distribution tool for your stuff (yes, all your stuff™)"

	app.Action = func(c *cli.Context) error {
		fmt.Print("hello and byee!\n")
		return cli.NewExitError("no pun intended", -1)
	}

	app.Run(os.Args)
}
