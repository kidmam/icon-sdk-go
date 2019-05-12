package main

import "fmt"
import "time"
import "sync"

// Define the job id type
type Id interface{}

// Define the job type
type Job func(id Id)

// Define the worker
func worker(job Job, queue chan Id, last_start_map *sync.Map, min_interval time.Duration) {
	// Get a job id from the queue
	for id := range queue {
		// Check time when a job last started
		last_start, ok := last_start_map.Load(id)

		// Calculate time when a job will start again
		should_start := time.Now()
		if ok {
			should_start = last_start.(time.Time).Add(min_interval)
		}

		// Run a job if should_start time is past due, otherwise skip it
		if time.Now().After(should_start) {
			// Update last start time
			last_start_map.Store(id, time.Now())
			job(id)
		} else {
			// Sleep for a one tenth of the interval to avoid extensive channel use
			time.Sleep(min_interval / 10)
		}

		// Put a job to the end of the queue
		queue <- id
	}
}

// Define the jobs and workers kikstarter
func kikstarter(job Job, ids []Id, max_workers int, min_interval time.Duration) {
	// Define the job queue
	queue := make(chan Id, len(ids))

	// Define the job last started hashmap
	// sync.Map is used to prevent race conditions
	var last_start_map sync.Map

	// Start workers
	for w := 0; w < max_workers; w++ {
		go worker(job, queue, &last_start_map, min_interval)
	}

	// Enqueue jobs
	for _, id := range ids {
		queue <- id
	}
}

// Main function
func main() {
	// Measure time for a debug purpose
	start := time.Now()

	// Define an amount of workers
	max_workers := 3

	// Define a mimimun interval for jobs
	min_interval := 6 * time.Second

	// Define the results channel
	results := make(chan Id)

	// Define the job
	job := func(id Id) {
		// Output debug information
		fmt.Println("started job", id, "on", time.Now().Sub(start).Round(time.Second))

		// Sleep id seconds
		time.Sleep(time.Duration(id.(int)) * time.Second)

		// Write result to channel
		results <- id

		// Output debug information
		fmt.Println("finished job", id, "on", time.Now().Sub(start).Round(time.Second))
	}

	// Schedule five jobs
	kikstarter(job, []Id{1, 2, 3, 4, 5}, max_workers, min_interval)

	// Wait for results
	// This will run program infietely
	for {
		<-results
	}
}
