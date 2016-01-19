package frontmatter

import (
	"bufio"
	"bytes"
	"errors"
	"io/ioutil"
	"log"
	"os"

	cmd "github.com/codegangsta/cli"
	"github.com/lcaballero/genfront/cli"
	"github.com/lcaballero/genfront/process"
	. "github.com/lcaballero/genfront/maybe"
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

func (p *FrontMatterProcess) Run() {
	log.Printf("Reading input file: %s", JoinCwd(p.InputFile()))
	b, err := ioutil.ReadFile(p.InputFile())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Separating embedded front-matter")
	err = p.portions.Read(bufio.NewReader(bytes.NewBuffer(b)))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Parsing front-matter")
	_, err = p.AddSettings(p.portions.FrontMatter)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Creating template")
	tpl, err := p.Env.CreateTemplate(p.portions.Template)
	if err != nil {
		log.Fatal(err)
	}

	p.Env.Debug(tpl, p.CliConf)

	log.Printf("Writing output file: %s", JoinCwd(p.OutputFile()))
	file, err := os.Create(p.OutputFile())
	if err == nil {
		defer file.Close()
		tpl.Execute(file, p.Env.ToMap())
	} else {
		log.Fatal(err)
	}
}
