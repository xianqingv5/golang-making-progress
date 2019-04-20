package main

import (
	"fmt"
	"sync"
)

func main() {
	// var mu sync.Mutex
	// mu.Lock()
	// go func() {
	// 	fmt.Println("nihao, shijie")
	// 	mu.Lock()
	// }()
	// mu.Unlock()

	// done := make(chan int)
	// go func() {
	// 	fmt.Println("nihao, shijie")
	// 	<-done
	// }()
	// done <- 1

	// done := make(chan int, 1)
	// go func() {
	// 	fmt.Println("nihao,shijie")
	// 	done <- 1
	// }()
	// <-done

	// done := make(chan int, 10)
	// for i := 0; i < cap(done); i++ {
	// 	go func(count int) {
	// 		fmt.Println("nihao, shijie")
	// 		done <- count
	// 	}(i)
	// }
	// for i := 0; i < cap(done); i++ {
	// 	fmt.Println(<-done)
	// }

	done := make(chan int, 10)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(count int) {
			fmt.Println("nihao, shijie")
			wg.Done()
			done <- count
		}(i)
	}

	wg.Wait()

	for j := 0; j < cap(done); j++ {
		fmt.Println(<-done)
	}
}
