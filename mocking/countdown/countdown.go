// Package countdown is to practice mocking.
package countdown

import (
	"fmt"
	"io"
	"time"
)

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
	s.SleepFunc(s.Duration)
}

// ConfigurableSleeper stores the duration and the sleep() function.
type ConfigurableSleeper struct {
	Duration  time.Duration
	SleepFunc func(time.Duration)
}

const (
	startCount = 3
	finalWord  = "Go!"
)
