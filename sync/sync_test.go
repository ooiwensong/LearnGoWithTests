package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})

	t.Run("it runs concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := NewCounter()

		var wg sync.WaitGroup // a WaitGroup waits for a collection of goroutines to finish
		wg.Add(wantedCount)   // Add() sets the number of goroutines to wait for

		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Inc()
				wg.Done() // each goroutine calls Done() when finished
			}()
		}
		wg.Wait() // Wait() is used to block until all goroutines have finished

		assertCounter(t, counter, wantedCount)
	})
}

// We have to pass Counter by reference rather than by value
// else a copy of the mutex will be created which is not good
func assertCounter(t testing.TB, got *Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d, want %d", got.Value(), want)
	}
}
