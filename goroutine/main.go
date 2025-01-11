package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("hello 1")
	time.Sleep(1 * time.Second)
	go hello()
	fmt.Println("hello 2")
}

func hello() {
	for {
		fmt.Println("Another thread")
	}
}
