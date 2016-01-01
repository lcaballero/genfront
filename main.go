package main

import (
	"os"
	"github.com/lcaballero/genfront/cli"
)

func main() {
	cli.NewCli().Run(os.Args)
}

