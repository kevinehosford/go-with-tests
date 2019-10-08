package server

import (
	"fmt"
	"net/http"
)

// PlayerServer ...
type PlayerServer struct {
	Store PlayerStore
}

// InMemoryPlayerStore ...
type InMemoryPlayerStore struct{}

// GetPlayerScore ...
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	player := req.URL.Path[len("/players/"):]

	fmt.Fprintf(w, fmt.Sprintf("%d", p.Store.GetPlayerScore(player)))
}

// PlayerStore ...
type PlayerStore interface {
	GetPlayerScore(string) int
}
