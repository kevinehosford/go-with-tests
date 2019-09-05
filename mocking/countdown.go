package main

import (
	"fmt"
	"io"
	"time"
)

// Sleeper ...
type Sleeper interface {
	Sleep()
}

// DefaultSleeper ...
type DefaultSleeper struct{}

// Sleep ...
func (d DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

// Countdown counts down
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := 3; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}

	sleeper.Sleep()
	fmt.Fprintf(out, "Go!")
}

func main() {
	// sleeper := &DefaultSleeper{}
	// Countdown(os.Stdout, sleeper)
}
