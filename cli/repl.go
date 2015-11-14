package cli

import (
	"bufio"
	"os"
	"io"
	cmd "github.com/codegangsta/cli"
	"fmt"
)

type Repl struct {
	prompt string
}

func (r *Repl) Prompt() {
	fmt.Println(r.prompt)
}

func (p *Repl) Start() {
	r := bufio.NewReader(os.Stdin)
	for {
		p.Prompt()
		line, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				return
			}
		}
		fmt.Println(line)
	}
}

func NewRepl(c *cmd.Context) {
	r := &Repl{
		prompt: ">",
	}
	r.Start()
}
