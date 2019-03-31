package main

import "fmt"

func loop() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
}

func main() {
	go loop()
	loop()
}

// ➜  goroutine git:(master) ✗ go run main.go
// 0 1 2 3 4 5 6 7 0 1 2 3 4 5 6 7 8 9 8 9 %
// ➜  goroutine git:(master) ✗ go run main.go
// 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 %
// ➜  goroutine git:(master) ✗ go run main.go
// 0 1 2 3 4 5 6 7 0 1 2 3 4 5 6 7 8 9 8 9 %
// ➜  goroutine git:(master) ✗ go run main.go
// 0 1 2 3 4 5 6 7 8 9 %
// ➜  goroutine git:(master) ✗ go run main.go
// 0 1 2 3 4 5 6 7 8 9 %
// ➜  goroutine git:(master) ✗ go run main.go
// 0 1 2 3 4 5 6 7 8 9 %
// ➜  goroutine git:(master) ✗ go run main.go
// 0 1 2 3 4 5 6 7 8 9 0 1 %
// ➜  goroutine git:(master) ✗ go run main.go
// 0 1 2 3 4 5 6 7 8 9 0 1 %
// ➜  goroutine git:(master) ✗ go run main.go
// 0 1 2 3 4 5 6 7 8 0 1 2 3 4 5 6 7 9 8 9 %
