package main

import (
	"os"
	"time"

	"github.com/patrickquigley102/lgwt2/mocking/countdown"
)

func main() {
	sleeper := &countdown.ConfigurableSleeper{
		Duration:  (1 * time.Second),
		SleepFunc: time.Sleep,
	}
	countdown.Countdown(os.Stdout, sleeper)
}
