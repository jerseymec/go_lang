package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//d := time.Now().Add(time.Millisecond * 5)
	ctx := context.Background() //background parent top level context
	//ctx,cancel := context.WithCancel(ctx)
	//ctx,cancel := context.WithDeadline(ctx,d) // use a deadline from the time the context started to close it
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*5) /// use a timeout (time delta) to close a context
	ctx = context.WithValue(ctx, "requestId", "12345")          //context with value is used for communication
	// between API packages that dont know anything about each

	for i := range gen(ctx) { // loop over the receiving channel as long as there is input
		time.Sleep(time.Millisecond)
		fmt.Println("Got from gen: ", i)
		if i == 10 { //if received i values is 10 break main routine
			break
		}
	}
	cancel()
	//time.Sleep(time.Second)
} // main go routine ends but the go routine in gen func is still going on .. Leaky go routine.

func gen(ctx context.Context) <-chan int { // func takes a context and returns a receiving channel
	out := make(chan int) //create new channel
	i := 0
	go func() { //start a goroutine
		for { // infinite loop
			select { // continuously send to channel out
			case <-ctx.Done():
				fmt.Println("Request ID : ", ctx.Value("requestId"))
				fmt.Println("stopping goroutine...")
				fmt.Println(ctx.Err())
				close(out) // close out the channel otherwise out channel will sit and wait for input from ctx
				return
			case out <- i: // every time it sends to channel out increment variable i
				i++

			}
		}
	}()
	return out // returns the channel on which goroutine is send in inifinte loop
}
