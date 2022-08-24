package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	Countdown(os.Stdout, defaultSleep{})
}

// Countdown is cool.
func Countdown(w io.Writer, s Sleeper) {
	for i := startCount; i > 0; i-- {
		fmt.Fprintln(w, i)
		s.Sleep()
	}
	fmt.Fprint(w, finalWord)
}

type defaultSleep struct{}

func (s defaultSleep) Sleep() {
	time.Sleep(1 * time.Second)
}

// Sleeper sleeps.
type Sleeper interface {
	Sleep()
}

const (
	startCount = 3
	finalWord  = "Go!"
)
