package main

import (
	"bytes"
	"log"
	"bufio"
	ray "github.com/aymerick/raymond"
	"github.com/spf13/viper"
	"io"
	"strings"
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

func AddHelpers() {
	ray.RegisterHelper("toPascal", toPascal)
}

func (p *Portions) Render() error {
	v := viper.New()
	v.SetConfigType("yaml")
	v.ReadConfig(bytes.NewBufferString(p.FrontMatter))

	settings := BuildData(v.AllSettings())
	AddHelpers()
	p.Rendered, p.Error = ray.Render(p.Template, settings)

	return p.Error
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
