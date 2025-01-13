package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu sync.Mutex
	cond := sync.NewCond(&mu)

	ready := false

	go func() {
		fmt.Println("Waiting for condition")

		mu.Lock()
		for !ready {
			cond.Wait()
		}
		fmt.Println("Proceeding...")
		mu.Unlock()
	}()

	time.Sleep(2 * time.Second)

	mu.Lock()
	ready = true
	cond.Signal()
	mu.Unlock()
	fmt.Println("Pushing Signal")

	time.Sleep(1 * time.Second)

	fmt.Println("DONE")

}
