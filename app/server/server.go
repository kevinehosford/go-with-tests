package server

import (
	"fmt"
	"net/http"
)

// PlayerServer ...
func PlayerServer(w http.ResponseWriter, req *http.Request) {
	player := req.URL.Path[len("/players/"):]

	fmt.Fprintf(w, getPlayerScore(player))
}

func getPlayerScore(player string) string {
	if player == "Pepper" {
		return "20"
	}

	if player == "Floyd" {
		return "10"
	}

	return ""
}
