package main

import (
	"fmt"
	"time"
)

func main() {
	channel1 := make(chan string)

	// Publisher
	go func() {
		for i := 1; i <= 10; i++ {
			channel1 <- fmt.Sprintf("Hello from %d", i)
			time.Sleep(time.Second)
		}
	}()

	// Subscriber
	go func() {
		for {
			msg := <-channel1
			fmt.Println(msg)
		}
	}()

	time.Sleep(5 * time.Second)

}
