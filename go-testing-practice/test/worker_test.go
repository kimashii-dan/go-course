package test

import (
	"sync"
	"testing"
)

func TestWorkerPool(t *testing.T) {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	var wg sync.WaitGroup
	for range 3 { // 3 workers
		wg.Go(func() { // for each worker perform 1 goroutine
			for job := range jobs { // 2. accept jobs from jobs channel
				results <- job * 2 // 3. send each job to results channel
			}
		})
	}

	jobCount := 10
	for j := range jobCount {
		jobs <- j // 1. send job to jobs channel
	}
	close(jobs) // 4. close jobs channel

	go func() {
		// 5. wait for workers and close results
		wg.Wait()
		close(results)
	}()

	sum := 0
	for result := range results { // 6. accept all jobs from results channel
		sum += result // 7. sum each job
	}

	expected := 90 // 8. set expected sum: 0*2 + 1*2 + ... + 9*2
	if sum != expected {
		t.Errorf("got %d; want %d", sum, expected)
	}
}
