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

	app.Commands = []cli.Command{
		cli.Command{
			Name:  "install",
			Usage: "Downloads the specified package and outputs it onto the specified directory",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "version, v",
					Usage: "Version of the package to be installed",
				},

				cli.StringFlag{
					Name:  "outputFolder, o",
					Usage: "Path to the output folder for the package to be installed",
					Value: "./",
				},

				cli.StringFlag{
					Name:  "source, s",
					Usage: "Base URL for the desired zippy endpoint to be looked at for this install operation",
				},
			},
		},
	}

	app.Action = func(c *cli.Context) error {
		fmt.Print("hello and byee!\n")
		return cli.NewExitError("no pun intended", -1)
	}

	app.Run(os.Args)
}
