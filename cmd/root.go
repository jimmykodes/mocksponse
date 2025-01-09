package cmd

import (
	"github.com/jimmykodes/gommand"
	"github.com/jimmykodes/mocksponse/cmd/run"
)

func Cmd() *gommand.Command {
	cmd := &gommand.Command{
		Name: "mocksponse",
	}
	cmd.SubCommand(
		run.Cmd(),
	)
	return cmd
}
