package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			c1 <- "Every 500ms"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			c2 <- "Every 2 Seconds"
		}
	}()

	//
	//for {
	//	msg1 := <- c1 // Channel c1 sends messages at a faster rate
	//	fmt.Println(msg1)
	//	msg2 := <-c2 // blocks c1 as it is slower
	//	fmt.Println(msg2) // channel 2 sends messages at a slower rate
	//
	//}

	for {
		select { // select blocks picks any options thats available to execute
		case msg1 := <-c1: // faster channels sends and receives and therefore executes without being blocked by slower channel further in code
			fmt.Println(msg1)
		case msg2 := <-c2: // slower channel doesnt block faster channel
			fmt.Println(msg2)

		}
	}
}
