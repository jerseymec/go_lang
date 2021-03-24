package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background() //background parent top level context

	for i := range gen(ctx) { // loop over the recieving channel as long as there is input
		fmt.Println("Got from gen: ", i)
		if i == 10 { //if received i values is 10 break main routine
			break
		}
	}
} // main go routine ends but the go routine in gen func is still going on .. Leaky go routine.

func gen(ctx context.Context) <-chan int { // func takes a context and returns a recieving channel
	out := make(chan int) //create new channel
	i := 0
	go func() { //start a goroutine
		for { // infinite loop
			select { // continuously send to channel out
			case out <- i: // every time it sends to channel out increment variable i
				i++

			}
		}
	}()
	return out // returns the channel on which goroutine is send in inifinte loop
}
