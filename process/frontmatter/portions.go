package frontmatter

import (
	"bufio"
	"bytes"
	"io"
	"strings"
	"log"
)

type Portions struct {
	FrontMatter string
	Template    string
}

func (p *Portions) Read(r *bufio.Reader) error {
	fm := bytes.NewBuffer([]byte{})
	tp := bytes.NewBuffer([]byte{})
	state := Initial
	var buf *bytes.Buffer = fm
	line, prefix, err := r.ReadLine()

	if err != nil {
		return err
	}

	log.Println("Reading FrontMatter in FrontMatter file.")
	n := 0
	for err == nil && line != nil {
		n++
		log.Printf("Line number: %d\n", n)
		if string(line) == "---" {
			switch state {
			case Initial:
				if n > 1 {
					log.Printf("Normally usage requires '---' on first line, but found on line: %d\n", n)
				}
				state = FrontMatter
				buf = fm
			case FrontMatter:
				log.Println("Reading Template in FrontMatter file.")
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
