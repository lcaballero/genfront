package cli

import (
	"fmt"

	"strings"

	cmd "github.com/codegangsta/cli"
)

const (
	line         = "line"
	input        = "input"
	output       = "output"
	debug        = "debug"
	noop         = "noop"
	template     = "template"
	datafile     = "data-file"
	tabDelimited = "tab-delimited"
	varName      = "var-name"
)

type CliConf struct {
	ctx *cmd.Context
}

func NewCliConf(c *cmd.Context) *CliConf {
	return &CliConf{ctx: c}
}

// KeyAndFile translate the --data-file flag value into a key:file parts.
// If there was not a --data-file flag or if there are no key:file
// pairs, it will return an error.
func (c *CliConf) KeyAndFile() (key, file string, err error) {
	spec := c.ctx.String(datafile)
	split := strings.Split(spec, ":")
	if len(split) != 2 {
		return "", "", fmt.Errorf("Expected key:data-file flag value, but found '%s'", spec)
	}
	key, file, err = split[0], split[1], nil
	return key, file, err
}
func (c *CliConf) IsTabDelimited() bool {
	return c.ctx.IsSet(tabDelimited)
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
func (p *CliConf) VarName() string {
	return p.ctx.String(varName)
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
func (p *CliConf) HasVarName() bool {
	return p.ctx.IsSet(varName)
}

func (p *CliConf) String() string {
	return fmt.Sprintf(`Line: %d
InputFile: %s
OutputFile: %s
Noop: %t
Debug: %t
Template: %s
IsTabDelimited: %v`,
		p.Line(),
		p.InputFile(),
		p.OutputFile(),
		p.Noop(),
		p.Debug(),
		p.Template(),
		p.IsTabDelimited())
}
