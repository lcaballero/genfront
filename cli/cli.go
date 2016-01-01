package cli

import (
	cmd "github.com/codegangsta/cli"
)

var usage = "Provides various Go generation utilities."

type Processor func(c *cmd.Context)
type Processors struct {
	FrontMatter, FieldProcessor Processor
}

func NewCli(p *Processors) *cmd.App {
	app := cmd.NewApp()
	app.Name = "genfront"
	app.Version = "0.0.1"
	app.Usage = usage
	app.Commands = []cmd.Command{
		frontCommand(p.FrontMatter),
		fieldsCommand(p.FieldProcessor),
	}
	return app
}

func fieldsCommand(p Processor) cmd.Command {
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
		Usage: "Process struct fields for sql io.",
		Action: p,
		Flags: flags(debugFlag(), custom...),
	}
}

func frontCommand(p Processor) cmd.Command {
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
		Usage: "Runs generator based on a front-matter file.",
		Flags: flags(debugFlag(), custom...),
		Action: p,
	}
}

func debugFlag() []cmd.Flag {
	return []cmd.Flag{
		cmd.BoolFlag{
			Name: "noop",
			Usage: "Doesn't generate source. Use with the --debug flag.",
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
