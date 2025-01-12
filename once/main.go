package main

import (
	"fmt"
	"sync"
)

var (
	once sync.Once
	wg   sync.WaitGroup
)

func main() {
	initialize := func() {
		fmt.Println("Initial Project")
	}

	doWork := func(id int) {
		defer wg.Done()
		fmt.Printf("Worker %d starting\n", id)
		once.Do(initialize) // This will only be executed once
		fmt.Printf("Worker %d completed\n", id)
	}

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go doWork(i)
	}

	wg.Wait()
}
