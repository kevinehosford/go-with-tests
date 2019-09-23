package context

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response string
	t        *testing.T
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, i := range s.response {
			select {
			case <-ctx.Done():
				s.t.Log("Spy was canceled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(i)
			}
		}

		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

type SpyRecorder struct {
	written bool
}

func (r *SpyRecorder) Header() http.Header {
	r.written = true
	return nil
}

func (r *SpyRecorder) Write(b []byte) (int, error) {
	r.written = true
	return 0, errors.New("not implemented")
}

func (r *SpyRecorder) WriteHeader(statusCode int) {
	r.written = true
}

func TestContext(t *testing.T) {
	t.Run("fetches the data", func(t *testing.T) {
		data := "response data"
		// create store
		store := &SpyStore{response: data, t: t}

		// create server
		svr := Server(store)

		// create httptest
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		// create httprecorder
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		// check if httprecorder.Response.Body == store.Data
		if response.Body.String() != data {
			t.Errorf("Expected %q but got %q", data, response.Body.String())
		}
	})

	t.Run("cancels after a given time", func(t *testing.T) {
		svr := Server(&SpyStore{response: "response data"})

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		cancellingCtx, cancel := context.WithCancel(request.Context())

		time.AfterFunc(5*time.Millisecond, cancel)

		request = request.WithContext(cancellingCtx)
		response := &SpyRecorder{}

		svr.ServeHTTP(response, request)

		if response.written {
			t.Errorf("store should not have written response")
		}
	})
}
