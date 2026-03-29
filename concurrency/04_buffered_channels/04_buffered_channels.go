package main

import (
	"fmt"
	"time"
)

func main() {
	buffered := make(chan string, 2)

	buffered <- "first"
	buffered <- "second"
	fmt.Println("sent two values without blocking")

	go func() {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("receiver started reading")
		fmt.Println("received:", <-buffered)
		fmt.Println("received:", <-buffered)
		fmt.Println("received:", <-buffered)
	}()

	fmt.Println("attempting third send (waits until receiver reads)")
	buffered <- "third"
	fmt.Println("third value sent")

	time.Sleep(1 * time.Second)
}
