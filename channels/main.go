package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		time.Sleep(time.Second)
		fmt.Println("Test 1")
		ch <- 10
		ch <- 20
		close(ch)
	}()

	fmt.Println("Test 2")
	v := <-ch // Turn into wait state

	fmt.Println(v)

}
