package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string, 50)
	go count("sheep", c)

	for msg := range c { /// msg<-c recieve messages from channel c untill channel closes ..blocking read(synchronous since no channel buffers)
		fmt.Println(msg)
	}
}

func count(thing string, c chan string) {
	defer close(c) // close channel at the end of func
	for i := 1; i <= 5; i++ {
		c <- thing //send strings to the channel c
		time.Sleep(time.Millisecond * 500)
	}
}
