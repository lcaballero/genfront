package main

import (
	"bytes"
	"log"
	"bufio"
	"github.com/spf13/viper"
	"io"
	"strings"
	"html/template"
)


type Portions struct {
	FrontMatter string
	Template string
	Rendered string
	Error error
}

const (
	Initial uint = iota
	FrontMatter
	Template
)

func BuildData(pairs map[string]interface{}) map[string]interface{} {
	env := make(map[string]interface{})
	for k,v := range BuildEnv() {
		env[k] = v
	}
	pairs["ENV"] = env
	return pairs
}

func (p *Portions) Settings() map[string]interface{} {
	v := viper.New()
	v.SetConfigType("yaml")
	v.ReadConfig(bytes.NewBufferString(p.FrontMatter))

	return BuildData(v.AllSettings())
}

func (p *Portions) Render() (*template.Template, error) {
	return template.New("").Funcs(BuildFuncMap()).Parse(p.Template)
}

func (p *Portions) Read(r *bufio.Reader) error {

	fm := bytes.NewBuffer([]byte{})
	tp := bytes.NewBuffer([]byte{})
	state := Initial
	var buf *bytes.Buffer = fm
	line, prefix, err := r.ReadLine()

	for err == nil && line != nil {
		if string(line) == "---" {
			switch state {
			case Initial:
				state = FrontMatter
				buf = fm
			case FrontMatter:
				state = Template
				buf = tp
			default:
				break
			}
			line, prefix, err = r.ReadLine()
			continue
		}
		buf.Write(line)
		if !prefix {
			buf.WriteByte('\n')
		}

		line, prefix, err = r.ReadLine()
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
	}

	p.FrontMatter = strings.Trim(fm.String(), " \t\r\n")
	p.Template = strings.Trim(tp.String(), " \t\r\n")

	return nil
}
