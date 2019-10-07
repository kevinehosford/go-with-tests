package main

import (
	"log"
	"net/http"

	"github.com/kevinehosford/go-with-tests/app/server"
)

func main() {
	handlerFunc := http.HandlerFunc(server.PlayerServer)

	if err := http.ListenAndServe(":5000", handlerFunc); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
