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

	//////////////////////
	jobCh := make(chan int, 10)
	resultCh := make(chan int, 10)

	for i := range 10 {
		jobCh <- i + 1
	}
	close(jobCh)

	go powerTwo(jobCh, resultCh)

	for range 10 { // range resultCh ไม่ได้เพราะ five มารับคัวแปรใน resultCh ไป
		five := <-resultCh
		fmt.Println(five)
	}

}

func powerTwo(jobCh <-chan int, resultCh chan<- int) {
	for i := range jobCh {
		resultCh <- i * 2
	}
}
