package cli

import (
	"fmt"

	"strings"

	cmd "github.com/codegangsta/cli"
)

const (
	line = "line"
	input = "input"
	output = "output"
	debug = "debug"
	noop = "noop"
	template = "template"
	datafile = "data-file"
	datafiles = "data-files"
	tabDelimited = "tab-delimited"
)

type CliConf struct {
	ctx *cmd.Context
}

func NewCliConf(c *cmd.Context) *CliConf {
	return &CliConf{
		ctx: c,
	}
}

type DataFile struct {
	Key, File string
}

func (c *CliConf) DataFile() (DataFile, error) {
	spec := c.ctx.String(datafile)
	return c.splitKeyedData(spec)
}
func (c *CliConf) splitKeyedData(spec string) (DataFile, error) {
	split := strings.Split(spec, ":")
	if len(split) != 2 {
		return DataFile{}, fmt.Errorf("Expected key:data-file flag value, but found '%s'", spec)
	}
	return DataFile{Key:split[0], File:split[1]}, nil
}
func (c *CliConf) DataFiles() ([]DataFile, error) {
	spec := c.ctx.String(datafiles)
	split := strings.Split(spec, ",")
	keyed := make([]DataFile, 0)
	for _,split := range split {
		df, err := c.splitKeyedData(split)
		if err != nil {
			return nil, err
		}
		keyed = append(keyed, df)
	}
	return keyed, nil
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

func (p *CliConf) Debug() bool {
	return p.ctx.Bool(debug)
}
func (p *CliConf) Noop() bool {
	return p.ctx.Bool(noop)
}

func (p *CliConf) HasDataFile() bool {
	return p.ctx.IsSet(datafile)
}
func (p *CliConf) HasDataFiles() bool {
	return p.ctx.IsSet(datafiles)
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
