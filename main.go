package main

import (
	"os"

	"github.com/daniel-888/golang-blockchain/cli"
)

func main() {
	defer os.Exit(0)
	cmd := cli.CommandLine{}
	cmd.Run()
}
