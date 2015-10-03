package main

import (
	cmd "github.com/codegangsta/cli"
	"log"
	"errors"
	"io/ioutil"
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"os"
)

// Cli provides the context (flags and values) with which to run a process for
// generating over a front matter file.
func NewFrontMatterProcessor(c *cmd.Context) {
	p := &Process{
		Ctx: c,
		Portions: &Portions{},
	}
	err := p.Validate()
	if err != nil {
		log.Fatal(err)
	}
	p.Run()
}

type Process struct {
	*Portions
	Ctx *cmd.Context
}

func (c *Process) HasString(s string) bool {
	return c.Ctx.String(s) != ""
}
func (p *Process) Validate() error {
	if p.HasString("input") && p.HasString("output") {
		return nil
	} else {
		return errors.New("Both --input and --output are required")
	}
}
func (p *Process) InputFile() string {
	return p.Ctx.String("input")
}
func (p *Process) OutputFile() string {
	return p.Ctx.String("output")
}
func (p *Process) Debug() bool {
	return p.Ctx.Bool("debug")
}
func (p *Process) NoSource() bool {
	return p.Ctx.Bool("no-source")
}

func (p *Process) dumpEnv() {
	fmt.Println()
	fmt.Println("ENV: ")
	fmt.Println(strings.Repeat("-", 80))
	ShowEnvironment()
	fmt.Println()
	fmt.Println("File To be Written: ", p.OutputFile())
	fmt.Println(strings.Repeat("-", 80))

	if p.NoSource() {
		fmt.Println("Generated source code supressed with --no-source")
	} else {
		fmt.Println(p.Rendered)
	}
}

func (p *Process) Run() {
	b, err := ioutil.ReadFile(p.InputFile())
	if err != nil {
		log.Fatal(err)
	}
	reader := bufio.NewReader(bytes.NewBuffer(b))
	p.Read(reader)
	tpl,err := p.Render()
	if err != nil {
		p.dumpEnv()
		log.Fatal(err)
	}

	if !p.Debug() {
		file, err := os.Create(p.OutputFile())
		if err == nil {
			tpl.Execute(file, p.Settings())
		} else {
			log.Fatal(err)
		}
	} else {
		p.dumpEnv()
	}
}

