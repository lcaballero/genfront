package cli

import (
	cmd "github.com/codegangsta/cli"
	"genfront/process"
)

var usage = "Converts processes a front matter file with yaml data and handlebars template."


func NewCli() *cmd.App {
	app := cmd.NewApp()
	app.Name = "genfront"
	app.Version = "0.0.1"
	app.Usage = usage
	app.Commands = []cmd.Command{
		front(),
		fields(),
	}
	return app
}

func fields() cmd.Command {
	custom := []cmd.Flag{
		cmd.StringFlag{
			Name: "output",
			Usage: "Name of source-code output file",
		},
		cmd.IntFlag{
			Name: "line",
			Usage: "Line number of this instance.",
		},
	}
	return cmd.Command{
		Name: "fields",
		Action: process.NewFieldProcessor,
		Flags: flags(debug(), custom...),
	}
}

func front() cmd.Command {
	custom := []cmd.Flag{
		cmd.StringFlag{
			Name: "input",
			Usage: "Front-matter file to process.",
		},
		cmd.StringFlag{
			Name: "output",
			Usage: "Name of source-code output file.",
		},
	}
	return cmd.Command{
		Name: "front",
		Flags: flags(debug(), custom...),
		Action: process.NewFrontMatterProcessor,
	}
}

func debug() []cmd.Flag {
	return []cmd.Flag{
		cmd.BoolFlag{
			Name: "no-source",
			Usage: "Hides generated source when using debug flag.",
		},
		cmd.BoolFlag{
			Name: "debug",
			Usage: "Process file, output to std-out, and show data points",
		},
	}
}

func flags(b []cmd.Flag, flags ...cmd.Flag) []cmd.Flag {
	return append(b, flags...)
}

