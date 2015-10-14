package main

import (
	"flag"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	// Define command line arguments.
	fibStart := flag.Int("maxfib", 40, "Defines the maximum (starting) value for the Fibonacci function.")
	numCores := flag.Int("cores", runtime.NumCPU(), "Overrides the automaticaly detected number of cores.")
	flag.Parse()
	// For each core run one worker.
	var wg sync.WaitGroup
	for i := 0; i < *numCores; i++ {
		wg.Add(1)
		go worker(*fibStart, &wg)
	}
	// Wait for all workers to finish.
	wg.Wait()
}

// Worker function. Measures execution time for Fibonacci function.
func worker(n int, wg *sync.WaitGroup) {
	// Signal the waitgroup when done.
	defer wg.Done()
	// Timestamp of execution start
	start := time.Now()
	fib(n)
	// Difference between current timestamp and execution start timestamp.
	fmt.Println(time.Since(start))
}

// Fibonacci function.
func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
