package cli

import (
	"fmt"

	cmd "github.com/codegangsta/cli"
	"strings"
)

const (
	line     = "line"
	input    = "input"
	output   = "output"
	debug    = "debug"
	noop     = "noop"
	template = "template"
	datafile = "data-file"
)

type CliConf struct {
	ctx *cmd.Context
}

func NewCliConf(c *cmd.Context) *CliConf {
	return &CliConf{
		ctx: c,
	}
}

func (c *CliConf) DataFile() (string, string, error) {
	spec := c.ctx.String(datafile)
	split := strings.Split(":", spec)
	if len(split) != 2 {
		return "", "", fmt.Errorf("Expected key:data-file flag value, but found '%s'", spec)
	}
	return split[0], split[1], nil
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
func (p *CliConf) Template() string {
	return p.ctx.String(template)
}

func (p *CliConf) Debug() bool {
	return p.ctx.Bool(debug)
}
func (p *CliConf) Noop() bool {
	return p.ctx.Bool(noop)
}

func (p *CliConf) HasDataFile() bool {
	return p.ctx.IsSet(datafile)
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
func (p *CliConf) HasTemplate() bool {
	return p.ctx.IsSet(template)
}

func (p *CliConf) String() string {
	return fmt.Sprintf(`Line: %d
InputFile: %s
OutputFile: %s
Noop: %t
Debug: %t
Template: %s`,
		p.Line(),
		p.InputFile(),
		p.OutputFile(),
		p.Noop(),
		p.Debug(),
		p.Template())
}
