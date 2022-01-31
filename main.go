package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/jimmykodes/mocksponse/internal/server"
)

func main() {
	port := flag.Int("p", 5000, "port to run server on")
	file := flag.String("f", "", "specify a recipe file other than recipe.(yaml|yml)")
	flag.Parse()

	filename, err := getFile(*file)
	if err != nil {
		log.Println(err)
		return
	}
	svr, err := server.New(filename, *port)
	if err != nil {
		log.Println("error initializing server", err)
		return
	}
	if err := svr.Run(); err != nil {
		log.Println("error running server", err)
		return
	}
}

func getFile(fromFlag string) (string, error) {
	if fromFlag != "" {
		_, err := os.Stat(fromFlag)
		if err != nil {
			return "", fmt.Errorf("could not find specified recipe file")
		} else {
			return fromFlag, nil
		}
	} else {
		if _, err := os.Stat("recipe.yaml"); err == nil {
			return "recipe.yaml", nil
		}
		if _, err := os.Stat("recipe.yml"); err == nil {
			return "recipe.yml", nil
		}
	}
	return "", fmt.Errorf("no valid recipe file found")
}
