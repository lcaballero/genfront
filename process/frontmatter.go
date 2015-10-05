package process

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
	"html/template"
)

const (
	Initial uint = iota
	FrontMatter
	Template
)

// Cli provides the context (flags and values) with which to run a process for
// generating over a front matter file.
func NewFrontMatterProcessor(c *cmd.Context) {
	p := &FrontMatterProcess{
		Processor: &Processor{ Ctx: c },
		Portions: &Portions{},
	}

	err := p.Validate()
	if err != nil {
		log.Fatal(err)
	}
	p.Run()
}

type FrontMatterProcess struct {
	*Portions
	*Processor
}

func (p *FrontMatterProcess) Validate() error {
	if p.HasString("input") && p.HasString("output") {
		return nil
	} else {
		return errors.New("Both --input and --output are required")
	}
}

func (p *FrontMatterProcess) dumpEnv(tpl *template.Template) {
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
		tpl.Execute(os.Stdout, p.Settings())
	}
}

func (p *FrontMatterProcess) Run() {
	b, err := ioutil.ReadFile(p.InputFile())
	if err != nil {
		log.Fatal(err)
	}
	reader := bufio.NewReader(bytes.NewBuffer(b))
	p.Read(reader)
	tpl,err := p.Render()
	if err != nil {
		p.dumpEnv(tpl)
		log.Fatal(err)
	}

	if !p.Debug() {
		file, err := os.Create(p.OutputFile())
		if err == nil {
			defer file.Close()
			tpl.Execute(file, p.Settings())
		} else {
			log.Fatal(err)
		}
	} else {
		p.dumpEnv(tpl)
	}
}
