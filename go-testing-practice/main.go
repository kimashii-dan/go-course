package main

import (
	"fmt"
	"time"
)

// worker function
func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job)
		time.Sleep(500 * time.Millisecond) // simulate work
		fmt.Printf("Worker %d finished job %d\n", id, job)
		results <- job * 2 // send result back
	}
}

// let's implement worker pool functionality :)
func main() {
	const numJobs = 5
	const numWorkers = 3

	// create channels
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// start worker-goroutines
	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
	}

	// send jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // close channel jobs

	// collect results
	for r := 1; r <= numJobs; r++ {
		result := <-results
		fmt.Printf("Received result: %d\n", result)
	}
}
