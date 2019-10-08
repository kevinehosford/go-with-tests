package main

import (
	"log"
	"net/http"

	"github.com/kevinehosford/go-with-tests/app/server"
)

func main() {
	store := server.InMemoryPlayerStore{}
	srv := &server.PlayerServer{&store}

	if err := http.ListenAndServe(":5000", srv); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
