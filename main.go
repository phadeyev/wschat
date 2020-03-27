package main

import (
	"log"

	"github.com/phadeyev/wschat/server"
)

func main() {
	serv := server.New()
	if err := serv.Start(); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
