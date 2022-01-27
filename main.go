package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/jimmykodes/mock-sponse/internal/server"
)

func main() {
	port := flag.Int("port", 5000, "port to run server on")
	flag.Parse()

	args := flag.Args()
	var filename string

	switch len(args) {
	case 1:
		filename = args[0]
	default:
		fmt.Println("invalid number of args")
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
