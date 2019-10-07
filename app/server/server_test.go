package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetServer(t *testing.T) {
	t.Run("returns Pepper's score", func(t *testing.T) {
		request, _ := newGetScoreRequest("Pepper")
		response := httptest.NewRecorder()

		PlayerServer(response, request)

		got := response.Body.String()
		want := "20"

		assertResponseBody(t, got, want)
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		request, _ := newGetScoreRequest("Floyd")
		response := httptest.NewRecorder()

		PlayerServer(response, request)

		got := response.Body.String()
		want := "10"

		assertResponseBody(t, got, want)
	})
}

func newGetScoreRequest(player string) (*http.Request, error) {
	return http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", player), nil)
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
