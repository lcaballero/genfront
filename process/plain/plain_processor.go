package plain

import (
	"errors"
	"io/ioutil"
	"log"
	"os"

	"github.com/lcaballero/genfront/cli"
	"github.com/lcaballero/genfront/maybe"
	"github.com/lcaballero/genfront/process"
	"github.com/lcaballero/genfront/process/datafiles"
	"path/filepath"
)

type PlainProcessor struct {
	*cli.CliConf
	*process.Env
}

func RunPlainProcessor(c *cli.CliConf) {
	p := &PlainProcessor{
		CliConf: c,
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
	log.Println("Reading template file", maybe.JoinCwd(p.Template()))
	b, err := ioutil.ReadFile(p.Template())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Creating template")
	tpl, err := p.Env.CreateTemplate(string(b))
	if err != nil {
		log.Fatal(err)
	}

	p.AddDataFileValues()

	p.Env.MaybeExit(p.CliConf, tpl, "")

	log.Printf("Writing output file: %s", maybe.JoinCwd(p.OutputFile()))
	file, err := os.Create(p.OutputFile())
	if err == nil {
		defer file.Close()
		tpl.Execute(file, p.Env.ToMap())
	} else {
		log.Fatal(err)
	}
}

func (p *PlainProcessor) AddDataFileValues() {
	if p.CliConf.HasDataFile() {
		keyed, err := p.CliConf.DataFile()
		if err != nil {
			log.Fatal(err)
		}
		p.AddSingleFile(keyed)
		return
	}
	if p.CliConf.HasDataFiles() {
		p.AddFiles()
		return
	}
}

func (p *PlainProcessor) AddFiles() {
	keyed, err := p.CliConf.DataFiles()
	if err != nil {
		log.Fatal(err)
	}
	for _, k := range keyed {
		p.AddJsonValues(k)
	}
}

func (p *PlainProcessor) AddSingleFile(keyed cli.DataFile) {
	ext := filepath.Ext(keyed.File)
	switch ext {
	case ".tsv", ".csv":
		p.AddTabSepValues(keyed)
	case ".json":
		p.AddJsonValues(keyed)
	}
}

func (p *PlainProcessor) AddTabSepValues(keyed cli.DataFile) {
	delimiter := ','
	log.Println(p.CliConf.IsTabDelimited())
	if p.CliConf.IsTabDelimited() {
		delimiter = '\t'
	}

	csv, err := datafiles.NewCsvData(keyed, delimiter).Parse()
	if err != nil {
		log.Fatal(err)
	}

	data, err := csv.MapFieldNames()
	if err != nil {
		log.Fatal(err)
	}

	p.Env.Add(keyed.Key, data)
}

func (p *PlainProcessor) AddJsonValues(keyed cli.DataFile) {
	json := datafiles.NewJsonData(keyed)
	data, err := json.Unmarshal()
	if err != nil {
		log.Fatal(err)
	}

	p.Env.Add(keyed.Key, data)
}
