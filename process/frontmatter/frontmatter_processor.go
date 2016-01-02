package frontmatter

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"

	cmd "github.com/codegangsta/cli"
	"github.com/lcaballero/genfront/cli"
	"github.com/lcaballero/genfront/process"
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
		CliConf:  cli.NewCliConf(c),
		portions: &Portions{},
		Env:      process.NewEnv(),
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
	*process.Env
}

func (p *FrontMatterProcess) Validate() error {
	if p.HasInputFile() && p.HasOutputFile() {
		return nil
	} else {
		return errors.New("Both --input and --output are required")
	}
}

func (p *FrontMatterProcess) debug(tpl *template.Template) {
	w := os.Stdout
	fmt.Fprintf(w, "%s\n", p.Sep())
	p.ShowEnvironment(w)
	fmt.Fprintf(w, "%s\n", p.Sep())
	fmt.Fprintln(w, p.CliConf)
	fmt.Fprintf(w, "%s\n", p.Sep())

	if p.Noop() {
		fmt.Fprintln(w, "Generated source code supressed with --noop")
	} else {
		tpl.Execute(w, p.Env.ToMap())
	}
}

func (p *FrontMatterProcess) Run() {
	log.Printf("Reading input file: %s", p.InputFile())

	b, err := ioutil.ReadFile(p.InputFile())
	if err != nil {
		log.Fatal(err)
	}
	reader := bufio.NewReader(bytes.NewBuffer(b))

	err = p.portions.Read(reader)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Creating template")

	p.AddSettings(p.portions.FrontMatter)

	tpl, err := p.Env.CreateTemplate(p.portions.Template)
	if err != nil {
		log.Fatal(err)
	}

	if p.Debug() {
		log.Printf("Skipping writing output file: %s", p.OutputFile())
		p.debug(tpl)
	} else {
		log.Printf("Writing output file: %s", p.OutputFile())

		file, err := os.Create(p.OutputFile())
		if err == nil {
			defer file.Close()
			tpl.Execute(file, p.Env.ToMap())
		} else {
			log.Fatal(err)
		}
	}
}
