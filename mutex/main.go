package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	mu      sync.Mutex
	balance int = 1000
)

func main() {
	doneCh := make(chan bool, 3)

	go UpdateBalance(doneCh, 100)
	go UpdateBalance(doneCh, 200)
	go UpdateBalance(doneCh, 100)
	<-doneCh
	<-doneCh
	<-doneCh

	fmt.Println(balance)
}

func UpdateBalance(doneCh chan<- bool, amount int) {
	mu.Lock()
	fmt.Println("Updating Lock")
	time.Sleep(time.Second)

	balance -= amount
	doneCh <- true

	mu.Unlock()
	fmt.Println("Updated UnLock")
}
