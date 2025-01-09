package run

import (
	"fmt"
	"os"

	"github.com/jimmykodes/gommand"
	"github.com/jimmykodes/gommand/flags"
	"github.com/jimmykodes/mocksponse/internal/server"
)

func Cmd() *gommand.Command {
	cmd := &gommand.Command{
		Name: "run",
		FlagSet: flags.NewFlagSet().AddFlags(
			flags.IntFlagS("port", 'p', 5000, "Port to listen on"),
			flags.StringFlagS("file", 'f', "", "file containing mock specs. default recipe.(yml|yaml)"),
		),
		Run: func(ctx *gommand.Context) error {
			filename, err := getFile(ctx.Flags().String("file"))
			if err != nil {
				return err
			}
			svr, err := server.New(filename, ctx.Flags().Int("port"))
			if err != nil {
				return err
			}
			return svr.Run()
		},
	}
	return cmd
}

func getFile(fromFlag string) (string, error) {
	if fromFlag == "" {
		if _, err := os.Stat("recipe.yaml"); err == nil {
			return "recipe.yaml", nil
		}
		if _, err := os.Stat("recipe.yml"); err == nil {
			return "recipe.yml", nil
		}
		return "", fmt.Errorf("no valid recipe file found")
	}

	stat, err := os.Stat(fromFlag)
	if err != nil {
		return "", fmt.Errorf("could not find specified recipe file")
	}
	if stat.IsDir() {
		return "", fmt.Errorf("%s is a directory: must provide a file", fromFlag)
	}

	return fromFlag, nil
}
