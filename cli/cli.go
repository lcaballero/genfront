package cli

import (
	cmd "github.com/codegangsta/cli"
)

const (
	DefaultFieldTemplate = "struct_sql_tomap.fm"
	usage = "Provides various Go generation utilities."
)

type Processor func(c *cmd.Context)
type Processors struct {
	DocTableProcessor, FrontMatter, FieldProcessor, PlainProcessor Processor
}

func NewCli(p *Processors) *cmd.App {
	app := cmd.NewApp()
	app.Name = "genfront"
	app.Version = "0.0.1"
	app.Usage = usage
	app.Commands = []cmd.Command{
		frontCommand(p.FrontMatter),
		fieldsCommand(p.FieldProcessor),
		plainCommand(p.PlainProcessor),
		doctableCommand(p.DocTableProcessor),
	}
	return app
}

func doctableCommand(p Processor) cmd.Command {
	custom := []cmd.Flag{
		cmd.StringFlag{
			Name:  "input",
			Usage: "Optional input .go file to process.",
		},
		cmd.StringFlag{
			Name:  "output",
			Usage: "Name of json file to output.",
		},
		cmd.StringFlag{
			Name:  "template",
			Usage: "The name of the template file to render.",
		},
		cmd.StringFlag{
			Name: "var-name",
			Usage: "Variable name for use in template.",
		},
		cmd.IntFlag{
			Name:  "line",
			Usage: "Line number of this instance.",
		},
	}
	return cmd.Command{
		Name:   "doctable",
		Usage:  "Process a template with Go environment with fields and comments.",
		Action: p,
		Flags:  flags(debugFlag(), custom...),
	}
}

func plainCommand(p Processor) cmd.Command {
	custom := []cmd.Flag{
		cmd.StringFlag{
			Name:  "output",
			Usage: "Name of source-code output file",
		},
		cmd.StringFlag{
			Name:  "template",
			Usage: "Optional value that specifies alternative template for processing",
			Value: DefaultFieldTemplate,
		},
		cmd.IntFlag{
			Name:  "line",
			Usage: "Line number of this instance.",
		},
		cmd.BoolFlag{
			Name:  "tab-delimited",
			Usage: "File is tab delimited",
		},
	}
	return cmd.Command{
		Name:   "plain",
		Usage:  "Process a template with Go environment.",
		Action: p,
		Flags:  flags(debugFlag(), custom...),
	}
}

func fieldsCommand(p Processor) cmd.Command {
	custom := []cmd.Flag{
		cmd.StringFlag{
			Name:  "output",
			Usage: "Name of source-code output file",
		},
		cmd.StringFlag{
			Name:  "template",
			Usage: "Optional value that specifies alternative template for processing",
			Value: DefaultFieldTemplate,
		},
		cmd.IntFlag{
			Name:  "line",
			Usage: "Line number of this instance.",
		},
	}
	return cmd.Command{
		Name:   "fields",
		Usage:  "Process struct fields for sql io.",
		Action: p,
		Flags:  flags(debugFlag(), custom...),
	}
}

func frontCommand(p Processor) cmd.Command {
	custom := []cmd.Flag{
		cmd.StringFlag{
			Name:  "input",
			Usage: "Front-matter file to process.",
		},
		cmd.StringFlag{
			Name:  "output",
			Usage: "Name of source-code output file.",
		},
	}
	return cmd.Command{
		Name:   "front",
		Usage:  "Runs generator based on a front-matter file.",
		Flags:  flags(debugFlag(), custom...),
		Action: p,
	}
}

func debugFlag() []cmd.Flag {
	return []cmd.Flag{
		cmd.BoolFlag{
			Name:  "noop",
			Usage: "Doesn't generate source. Use with the --debug flag.",
		},
		cmd.BoolFlag{
			Name:  "debug",
			Usage: "Process file, output to std-out, and show data points",
		},
		cmd.StringFlag{
			Name:  "data-file",
			Usage: "Provide data from a file.  Value should be name:file.ext",
		},
	}
}

func flags(b []cmd.Flag, flags ...cmd.Flag) []cmd.Flag {
	return append(b, flags...)
}
