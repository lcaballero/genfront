package frontmatter

import (
	cmd "github.com/codegangsta/cli"
	"log"
	"errors"
	"io/ioutil"
	"bufio"
	"bytes"
	"fmt"
	"os"
	"html/template"
	"github.com/lcaballero/genfront/process"
	"github.com/lcaballero/genfront/cli"
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
		CliConf: cli.NewCliConf(c),
		portions: &Portions{},
	}

	err := p.Validate()
	if err != nil {
		log.Fatal(err)
	}
	p.Run()
}

type FrontMatterProcess struct {
	portions *Portions
	*cli.CliConf
}

func (p *FrontMatterProcess) Validate() error {
	if p.HasInputFile() && p.HasOutputFile() {
		return nil
	} else {
		return errors.New("Both --input and --output are required")
	}
}

func (p *FrontMatterProcess) debug(tpl *template.Template) {
	fmt.Println()
	fmt.Println("ENV: ")
	fmt.Println(process.Sep())
	process.ShowEnvironment()
	fmt.Println()
	fmt.Println(p.CliConf)
	fmt.Println(process.Sep())

	if p.Noop() {
		fmt.Println("Generated source code supressed with --noop")
	} else {
		tpl.Execute(os.Stdout, p.portions.Settings())
	}
}

func (p *FrontMatterProcess) Run() {
	b, err := ioutil.ReadFile(p.InputFile())
	if err != nil {
		log.Fatal(err)
	}
	reader := bufio.NewReader(bytes.NewBuffer(b))

	err = p.portions.Read(reader)
	if err != nil {
		log.Fatal(err)
	}

	tpl,err := p.portions.CreateTemplate()
	if err != nil {
		log.Fatal(err)
	}

	if p.Debug() {
		p.debug(tpl)
	} else {
		file, err := os.Create(p.OutputFile())
		if err == nil {
			defer file.Close()
			tpl.Execute(file, p.portions.Settings())
		} else {
			log.Fatal(err)
		}
	}
}

