package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	finalWord      = "Go!"
	countdownStart = 3

	write = "write"
	sleep = "sleep"
)

func Countdown(w io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(w, i)
		sleeper.Sleep()
	}
	fmt.Fprint(w, finalWord)
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}

/** IMPLEMENTATIONS **/

// Sleep() is a dependency interface we use in our test codes for testing
type Sleeper interface {
	Sleep()
}

// type SpySleeper struct {
// 	Calls int
// }

// // This method makes SpySleeper implement the Sleeper interface
// func (s *SpySleeper) Sleep() {
// 	s.Calls++
// }

// Creating a custom sleeper struct for the main function to run the real sleeper
// type DefaultSleeper struct{}

// func (s *DefaultSleeper) Sleep() {
// 	time.Sleep(1 * time.Second)
// }

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration) // The value in the "duration" field is automatically passed into this func
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}
