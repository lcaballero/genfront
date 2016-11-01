package main

import (
	"os"

	"github.com/lcaballero/genfront/cli"
	"github.com/lcaballero/genfront/process/doctable"
	"github.com/lcaballero/genfront/process/fields"
	"github.com/lcaballero/genfront/process/frontmatter"
	"github.com/lcaballero/genfront/process/plain"
)

func main() {
	procs := cli.Processors{
		FieldProcessor:    fields.RunFieldProcessor,
		FrontMatter:       frontmatter.RunFrontMatterProcessor,
		PlainProcessor:    plain.RunPlainProcessor,
		DocTableProcessor: doctable.RunDocFinder,
	}
	cli.NewCli(procs).Run(os.Args)
}
