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
type InMemoryPlayerStore struct {
	Scores map[string]int
}

// NewInMemoryPlayerStore ...
func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

// GetPlayerScore ...
func (i *InMemoryPlayerStore) GetPlayerScore(name string) (int, bool) {
	val, ok := i.Scores[name]

	return val, ok
}

// RecordWin ...
func (i *InMemoryPlayerStore) RecordWin(name string) {
	val, _ := i.Scores[name]

	i.Scores[name] = val + 1
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		p.showScore(w, req)
	case http.MethodPost:
		p.processWin(w, req)
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

// PlayerStore ...
type PlayerStore interface {
	GetPlayerScore(string) (int, bool)
	RecordWin(string)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusAccepted)
	player := req.URL.Path[len("/players/"):]
	p.Store.RecordWin(player)
}

func (p *PlayerServer) showScore(w http.ResponseWriter, req *http.Request) {
	player := req.URL.Path[len("/players/"):]

	score, _ := p.Store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprintf(w, fmt.Sprintf("%d", score))
}
