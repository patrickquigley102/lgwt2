package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	sleeper := &ConfigurableSleeper{
		duration: (1 * time.Second),
		sleep:    time.Sleep,
	}
	Countdown(os.Stdout, sleeper)
}

// Countdown is cool.
func Countdown(w io.Writer, s Sleeper) {
	for i := startCount; i > 0; i-- {
		fmt.Fprintln(w, i)
		s.Sleep()
	}
	fmt.Fprint(w, finalWord)
}

// Sleeper sleeps.
type Sleeper interface {
	Sleep()
}

// Sleep sleeps() for the configured duration.
func (s *ConfigurableSleeper) Sleep() {
	s.sleep(s.duration)
}

// ConfigurableSleeper stores the duration and the sleep() function.
type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

const (
	startCount = 3
	finalWord  = "Go!"
)
