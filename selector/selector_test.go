package selector

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("check it works", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		want := fastServer.URL
		got, _ := Racer(want, slowServer.URL)

		if want != got {
			t.Errorf("Got %q, want %q", got, want)
		}
	})

	t.Run("check it errors", func(t *testing.T) {
		server := makeDelayedServer(20 * time.Second)

		defer server.Close()

		_, got := ConfigurableRacer(server.URL, server.URL, 10*time.Millisecond)

		if got == nil {
			t.Errorf("expected an error but didn't get one")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
