package main

import (
	"bytes"
	"reflect"
	"testing"
)

// OperationsSpy ...
type OperationsSpy struct {
	Calls []string
}

// Sleep ...
func (s *OperationsSpy) Sleep() {
	s.Calls = append(s.Calls, "sleep")
}

// Write ...
func (s *OperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, "write")
	return
}

func TestCountdown(t *testing.T) {
	t.Run("Testing output", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		operationsSpy := &OperationsSpy{}

		Countdown(buffer, operationsSpy)
		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("Got %q, want %q", got, want)
		}
	})

	t.Run("Testing operations", func(t *testing.T) {
		operationsSpy := &OperationsSpy{}

		Countdown(operationsSpy, operationsSpy)
		got := operationsSpy.Calls
		want := []string{
			"sleep",
			"write",
			"sleep",
			"write",
			"sleep",
			"write",
			"sleep",
			"write",
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %q, want %q", got, want)
		}
	})
}
