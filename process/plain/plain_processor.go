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
	"gopkg.in/ini.v1"
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
	log.Printf("Reading template file: %s\n", maybe.JoinCwd(p.Template()))
	b, err := ioutil.ReadFile(p.Template())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Creating template")
	tpl, err := p.Env.CreateTemplate(string(b))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Template created")

	ext, key, datafile, err := p.ExtKeyAndFile()
	if err != nil {
		log.Println("No data-file provided")
	} else {
		log.Printf("Rendering file: '%s', ext: '%s', key: '%s'\n", datafile, ext, key)
	}

	switch ext {
	case ".json":
		log.Printf("processing json, with key: %s", key)
		p.AddJsonValues(ext, key, datafile)
	case ".csv", ".tsv":
		log.Printf("processing csv, with key: %s", key)
		p.AddCsvValues(ext, key, datafile)
	case ".ini":
		log.Printf("processing ini, with key: %s", key)
		p.AddIniValues(ext, key, datafile)
	}

	p.AddFileValues()
	p.Env.MaybeExit(p.CliConf, tpl, "")

	log.Printf("Writing output file: %s\n", p.OutputFile())
	file, err := os.Create(p.OutputFile())
	if err == nil {
		defer file.Close()
		tpl.Execute(file, p.Env.ToMap())
	} else {
		log.Fatal(err)
	}
}

func (p *PlainProcessor) AddFileValues() {
	out_file := p.OutputFile()
	dir := filepath.Dir(out_file)
	base_dir := filepath.Base(dir)
	ext := filepath.Ext(out_file)
	out_name := filepath.Base(out_file)

	p.Env.Add("OUT_FILE", out_file)
	p.Env.Add("OUT_DIR", dir)
	p.Env.Add("OUT_BASE_DIR", base_dir)
	p.Env.Add("OUT_EXT", ext)
	p.Env.Add("OUT_NAME", out_name[:len(out_name)-len(ext)])
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
	log.Printf("is tab delimited: %t\n", p.CliConf.IsTabDelimited())
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

func (p *PlainProcessor) AddIniValues(ext, key, file string) {
	mapping := iniToMap(file)
	p.Env.Add(key, mapping)
}

// ExtKeyAndFile checks for a data-file flag and if there is not a flag
// returns and error, else it parse the value for the a key, naming the
// data in the file for use in a template, the name of the file itself
// and teh extension of the file so that a correct parser can be used
// to transform the content into data usable by a template.
func (p *PlainProcessor) ExtKeyAndFile() (ext, key, file string, err error) {
	if !p.CliConf.HasDataFile() {
		return "", "", "", errors.New("No data file")
	}
	key, file, err = p.CliConf.KeyAndFile()
	if err != nil {
		return "", "", "", err
	}
	return filepath.Ext(file), key, file, nil
}

func iniToMap(filename string) interface{} {
	file, err := ini.LoadSources(ini.LoadOptions{IgnoreInlineComment: true}, filename)

	if err != nil {
		panic(err)
	}

	res := map[string]interface{}{}
	sections := file.Sections()

	for i := 0; i < len(sections); i++ {
		sec := sections[i]
		keys := sec.Keys()
		kvp := map[string]interface{}{}

		for j := 0; j < len(keys); j++ {
			key := keys[j]
			name := key.Name()
			val := key.Value()

			kvp[name] = val
		}

		res[sec.Name()] = kvp
	}

	return res
}