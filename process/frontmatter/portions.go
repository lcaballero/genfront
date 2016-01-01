package frontmatter

import (
	"bytes"
	"bufio"
	"html/template"
	"io"
	"strings"
	"github.com/spf13/viper"
	"github.com/lcaballero/genfront/process"
)



type Portions struct {
	FrontMatter string
	Template string
	Error error
}

func (p *Portions) Settings() map[string]interface{} {
	v := viper.New()
	v.SetConfigType("yaml")
	v.ReadConfig(bytes.NewBufferString(p.FrontMatter))

	return process.BuildData(v.AllSettings())
}

func (p *Portions) CreateTemplate() (*template.Template, error) {
	return template.New("FrontMatterProcessor").Funcs(process.BuildFuncMap()).Parse(p.Template)
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
			return err
		}
	}

	p.FrontMatter = strings.Trim(fm.String(), " \t\r\n")
	p.Template = strings.Trim(tp.String(), " \t\r\n")

	return nil
}

