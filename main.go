package main

import (
	"os"
	"genfront/cli"
)

func main() {
	cli.NewCli().Run(os.Args)
}

