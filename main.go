package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/jimmykodes/mock-sponse/internal/server"
)

func main() {
	args := os.Args
	var filename string

	switch len(args) {
	case 2:
		filename = args[1]
	case 1:
		fallthrough
	default:
		fmt.Println("invalid number of args")
		return
	}

	port := flag.Int("port", 5000, "port to run server on")
	flag.Parse()

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
