package main

import (
	"log"

	"github.com/robert-w/go-server/internal/server"
)

func main() {
	err := server.Run()

	if err != nil {
		log.Fatal("Server failed to start with error: ", err)
	}
}
