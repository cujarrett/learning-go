package main

import (
	"context"
	"fmt"
	"time"
)

func doWork(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("worker %d stopped: %v\n", id, ctx.Err())
			return
		default:
			fmt.Printf("worker %d doing work\n", id)
			time.Sleep(200 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go doWork(ctx, 1)
	go doWork(ctx, 2)

	<-ctx.Done()
	time.Sleep(250 * time.Millisecond)
	fmt.Println("main exited cleanly")
}
