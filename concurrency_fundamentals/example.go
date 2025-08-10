package concurrencyfundamentals

import (
	"fmt"
	"time"
)

type Job struct {
	ID      int
	Payload string
}

func Worker(id int, jobs <-chan Job, results chan<- string) {
	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job.ID)
		time.Sleep(1 * time.Second)
		fmt.Printf("Worker %d finished job %d\n", id, job.ID)
		results <- fmt.Sprintf("Job %d processed by worker %d", job.ID, id)
	}
}

func Factory() {
	numWorkers := 3
	numJobs := 5

	jobs := make(chan Job, numJobs)
	results := make(chan string, numJobs)

	// Start workers
	for w := 1; w <= numWorkers; w++ {
		go Worker(w, jobs, results)
	}

	// Dispatch jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- Job{ID: j, Payload: fmt.Sprintf("Payload %d", j)}
	}
	close(jobs)

	// Collect results
	for r := 1; r <= numJobs; r++ {
		result := <-results
		fmt.Println("Result:", result)
	}

	fmt.Println("All jobs completed.")
}
