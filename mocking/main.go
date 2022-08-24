package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	Countdown(os.Stdout)
}

// Countdown is cool.
func Countdown(w io.Writer) {
	for i := startCount; i > 0; i-- {
		fmt.Fprintln(w, i)
		time.Sleep(1 * time.Second)
	}
	fmt.Fprint(w, finalWord)
}

const (
	startCount = 3
	finalWord  = "Go!"
)
