package main

import (
	cmd "github.com/codegangsta/cli"
)

var usage = "Converts processes a front matter file with yaml data and handlebars template."


func NewCli() *cmd.App {
	app := cmd.NewApp()
	app.Name = "ggen"
	app.Version = "0.0.1"
	app.Usage = usage
	app.Action = NewProcess
	app.Flags = []cmd.Flag{
		cmd.StringFlag{
			Name: "input",
			Usage: "Front-matter file to process",
		},
		cmd.StringFlag{
			Name: "output",
			Usage: "Name of source-code output file",
		},
		cmd.BoolFlag{
			Name: "debug",
			Usage: "Process file, output to std-out, and show data points",
		},
	}
	return app
}
