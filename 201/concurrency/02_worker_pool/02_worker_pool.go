package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		time.Sleep(300 * time.Millisecond)
		results <- fmt.Sprintf("worker %d processed job %d", id, job)
	}
}

func main() {
	jobs := make(chan int, 8)
	results := make(chan string, 8)

	const workerCount = 3
	var wg sync.WaitGroup

	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	for job := 1; job <= 8; job++ {
		jobs <- job
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println(result)
	}
}
