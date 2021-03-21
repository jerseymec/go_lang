package main

import "fmt"

func main() {
	const Numjobs = 100
	jobs := make(chan int, Numjobs) // create channels with buffer size of 100
	results := make(chan int, Numjobs)

	for i := 1; i <= 4; i++ {
		go worker(jobs, results) // create a worker pool of 4. workers can receive on channel jobs, execute task based on jobs  and send output on channel results
	}

	for i := 0; i < 100; i++ {
		jobs <- i // send 100 messages to jobs channel
	}
	close(jobs) // close channel when done. Since 100 buffer size .. channel can be close without receiving

	for j := 0; j < 100; j++ {
		fmt.Println(<-results) // print the messages received on channel results. will receive all 100
	}

}

func worker(jobs <-chan int, result chan<- int) { // recieve messages on jobs and send results of execution on results channels
	for n := range jobs { // receive number on jobs
		result <- fib(n) // send results of fibonacci of number from jobs channel
	}
}

func fib(t int) int { // Basic Fibo function
	if t <= 1 {
		return t
	}
	return fib(t-1) + fib(t-2)
}
