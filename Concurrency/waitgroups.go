package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	// waitgroup tracks or waits for the goroutines started within the main go routine
	go func() {
		count("sheep")
		wg.Done() // signals the go routine is ended and reduces the number of go routines to wait for
	}()

	for i := 1; i <= 5; i++ {
		wg.Add(1) // adding go routines to the waitgroup in main
		tex := strconv.Itoa(i)
		go worker(tex, &wg) // creating a new go routine with a pointer to the waitgroup used for waiting
	}

	wg.Wait() // main function will wait for the waitgroup to have no groutines to tract before continuing execution
}

func count(thing string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}

func worker(thing string, wg *sync.WaitGroup) {
	defer wg.Done() // defering the done to the waitgroup inside the sub function
	fmt.Printf("Worker %s starting\n", thing)
	time.Sleep(time.Millisecond * 500)
	fmt.Printf("Worker %s done\n", thing)

}
