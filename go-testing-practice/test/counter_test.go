package test

import (
	"go-testing-practice/service"
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	// create a subtest
	t.Run("it runs safely concurrently", func(t *testing.T) {
		// set expected count
		wantedCount := 1000

		// create a new counter
		counter := service.NewCounter()

		// add amount of all goroutines to waitgroup
		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for range wantedCount {
			// run 1000 goroutines
			go func() {
				// increment counter
				counter.Inc()

				// delete 1 goroutine from waitgroup
				wg.Done()
			}()
		}

		// wait all goroutines to finish
		wg.Wait()

		// helper function to check results
		assertCounter(t, counter, wantedCount)
	})
}

func assertCounter(t testing.TB, got *service.Counter, want int) {
	t.Helper()
	// if counter's value doesn't equal to expected value, then its failed.
	if got.Value() != want {
		t.Errorf("got %d, want %d", got.Value(), want)
	}
}
