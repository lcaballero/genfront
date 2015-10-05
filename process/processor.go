package process

import (
	cmd "github.com/codegangsta/cli"
)

type Processor struct {
	Ctx *cmd.Context
}

func (c *Processor) HasString(s string) bool {
	return c.Ctx.String(s) != ""
}
func (p *Processor) InputFile() string {
	return p.Ctx.String("input")
}
func (p *Processor) OutputFile() string {
	return p.Ctx.String("output")
}
func (p *Processor) Debug() bool {
	return p.Ctx.Bool("debug")
}
func (p *Processor) NoSource() bool {
	return p.Ctx.Bool("no-source")
}
