package plain

import (
	cmd "github.com/codegangsta/cli"
	"github.com/lcaballero/genfront/cli"
	"github.com/lcaballero/genfront/process"
	"log"
	"errors"
	"io/ioutil"
	"os"
	. "github.com/lcaballero/genfront/maybe"
)


type PlainProcessor struct {
	*cli.CliConf
	*process.Env
}

func RunPlainProcessor(c *cmd.Context) {
	p := &PlainProcessor{
		CliConf: cli.NewCliConf(c),
		Env:     process.NewEnv(),
	}
	err := p.Validate()
	if err != nil {
		log.Fatal(err)
	}
	p.Run()
}

func (p *PlainProcessor) Validate() error {
	if p.HasOutputFile() && p.HasTemplate() {
		return nil
	} else {
		return errors.New("output and template filenames are required")
	}
}

func (p *PlainProcessor) Run() {
	log.Println("Reading template file", JoinCwd(p.Template()))
	b, err := ioutil.ReadFile(p.Template())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Creating template")
	tpl, err := p.Env.CreateTemplate(string(b))
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
