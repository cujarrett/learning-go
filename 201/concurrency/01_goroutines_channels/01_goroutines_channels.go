package main

import (
	"fmt"
	"sync"
)

type squareResult struct {
	Number int
	Square int
}

func square(n int, out chan<- squareResult, wg *sync.WaitGroup) {
	defer wg.Done()
	out <- squareResult{Number: n, Square: n * n}
}

func main() {
	numbers := []int{2, 3, 4, 5}
	results := make(chan squareResult, len(numbers))

	var wg sync.WaitGroup
	for _, n := range numbers {
		wg.Add(1)
		go square(n, results, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Printf("%d squared is %d\n", result.Number, result.Square)
	}
}
