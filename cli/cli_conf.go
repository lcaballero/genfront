package cli

import (
	"fmt"

	cmd "github.com/codegangsta/cli"
)

const (
	line   = "line"
	input  = "input"
	output = "output"
	debug  = "debug"
	noop   = "noop"
)

type CliConf struct {
	ctx *cmd.Context
}

func NewCliConf(c *cmd.Context) *CliConf {
	return &CliConf{
		ctx: c,
	}
}

func (c *CliConf) Line() int {
	return c.ctx.Int(line)
}
func (p *CliConf) InputFile() string {
	return p.ctx.String(input)
}
func (p *CliConf) OutputFile() string {
	return p.ctx.String(output)
}

func (p *CliConf) Debug() bool {
	return p.ctx.Bool(debug)
}
func (p *CliConf) Noop() bool {
	return p.ctx.Bool(noop)
}

func (p *CliConf) HasOutputFile() bool {
	return p.ctx.IsSet(output)
}
func (p *CliConf) HasInputFile() bool {
	return p.ctx.IsSet(input)
}
func (p *CliConf) HasDebug() bool {
	return p.ctx.IsSet(debug)
}
func (p *CliConf) HasNoop() bool {
	return p.ctx.IsSet(noop)
}

func (p *CliConf) String() string {
	return fmt.Sprintf(`Line: %d
InputFile: %s
OutputFile: %s
Noop: %t
Debug: %t`,
		p.Line(),
		p.InputFile(),
		p.OutputFile(),
		p.Noop(),
		p.Debug())
}
