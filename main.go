package main

import (
	"os"

	"github.com/lcaballero/genfront/cli"
	"github.com/lcaballero/genfront/process/fields"
	"github.com/lcaballero/genfront/process/frontmatter"
)

func main() {
	procs := &cli.Processors{
		FieldProcessor: fields.RunFieldProcessor,
		FrontMatter:    frontmatter.NewFrontMatterProcessor,
	}
	cli.NewCli(procs).Run(os.Args)
}
