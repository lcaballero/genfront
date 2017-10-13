package cli

import (
	cmd "github.com/codegangsta/cli"
)

const (
	DefaultFieldTemplate = "struct_sql_tomap.fm"
	usage                = "Provides various Go generation utilities."
)

type Processor func(c *CliConf)

type Processors struct {
	DocTableProcessor Processor
	FrontMatter       Processor
	FieldProcessor    Processor
	PlainProcessor    Processor
}

func WithContext(p Processor) func(*cmd.Context) {
	return func(c *cmd.Context) {
		p(NewCliConf(c))
	}
}

func NewCli(p Processors) *cmd.App {
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
			Name:  "template",
			Usage: "The name of the template file to render.",
		},
//		cmd.StringFlag{
//			Name:  "var-name",
//			Usage: "Variable name for use in template.",
//		},
		cmd.IntFlag{
			Name:  "line",
			Usage: "Line number of this instance.",
		},
	}
	return cmd.Command{
		Name:   "doctable",
		Usage:  "Process a template with Go environment with fields and comments.",
		Action: WithContext(p),
		Flags:  Flags(DebugFlag(), DataFileFlag, custom),
	}
}

func plainCommand(p Processor) cmd.Command {
	custom := []cmd.Flag{
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
		cmd.StringFlag{
			Name:  "data-files",
			Usage: "Provide data from a files.  Value is comma sep of name:file.ext",
		},
	}
	return cmd.Command{
		Name:   "plain",
		Usage:  "Process a template with Go environment.",
		Action: WithContext(p),
		Flags:  Flags(DebugFlag(), DataFileFlag, custom),
	}
}

func fieldsCommand(p Processor) cmd.Command {
	custom := []cmd.Flag{
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
		Action: WithContext(p),
		Flags:  Flags(DebugFlag(), DataFileFlag, custom),
	}
}

func frontCommand(p Processor) cmd.Command {
	custom := []cmd.Flag{
		cmd.StringFlag{
			Name:  "input",
			Usage: "Front-matter file to process.",
		},
	}
	return cmd.Command{
		Name:   "front",
		Usage:  "Runs generator based on a front-matter file.",
		Action: WithContext(p),
		Flags:  Flags(DebugFlag(), custom),
	}
}

var DataFileFlag = []cmd.Flag{
	cmd.StringFlag{
		Name:  "data-file",
		Usage: "Provide data from a file. Value should be name:file.ext",
	},
}

func DebugFlag() []cmd.Flag {
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
			Name:  "output",
			Usage: "Name of source-code output file.",
		},
	}
}

func Flags(flags ...[]cmd.Flag) []cmd.Flag {
	rs := make([]cmd.Flag, 0)
	for _,gs := range flags {
		rs = append(rs, gs...)
	}
	return rs
}
