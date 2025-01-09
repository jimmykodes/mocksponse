package main

import (
	"os"

	"github.com/jimmykodes/mocksponse/cmd"
)

func main() {
	root := cmd.Cmd()
	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
