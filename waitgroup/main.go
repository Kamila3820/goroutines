package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)

	sleepDuration := time.Duration(rand.Intn(1000) * int(time.Millisecond))
	time.Sleep(sleepDuration)

	fmt.Printf("Worker %d completed\n", id)
}

func main() {
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait() //Help goroutine to wait other goroutine to work efficient
	fmt.Println("All works completed!!")
}
