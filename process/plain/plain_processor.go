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

	ext, key, datafile, err := p.Ext()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Rendering:", datafile, ext, key)
	switch ext {
	case ".json":
		log.Println("processing json")
		p.AddJsonValues(ext, key, datafile)
	case ".csv", ".tsv":
		log.Println("processing csv")
		p.AddCsvValues(ext, key, datafile)
	}

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

func (p *PlainProcessor) AddJsonValues(ext, key, file string) {
	json, err := datafiles.NewJsonData(key, file).Parse()
	if err != nil {
		log.Fatal(err)
	}
	p.Env.Add(key, json.Data)
}

func (p *PlainProcessor) AddCsvValues(ext, key, file string) {
	delimiter := ','
	log.Println(p.CliConf.IsTabDelimited())
	if p.CliConf.IsTabDelimited() {
		delimiter = '\t'
	}

	csv, err := datafiles.NewCsvData(key, file, delimiter).Parse()
	if err != nil {
		log.Fatal(err)
	}

	data, err := csv.MapFieldNames()
	if err != nil {
		log.Fatal(err)
	}

	p.Env.Add(csv.Key, data)
}

func (p *PlainProcessor) Ext() (string, string, string, error) {
	if !p.CliConf.HasDataFile() {
		return "", "", "", errors.New("No data file")
	}
	key, file, err := p.CliConf.DataFile()
	if err != nil {
		return "", "", "", err
	}
	return filepath.Ext(file), key, file, nil
}
