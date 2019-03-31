package main

import "fmt"

var ch1 chan int = make(chan int)
var ch2 chan int = make(chan int)

func worker1(s string) {
	fmt.Println(s)
	ch1 <- <-ch2
}

func worker2(s string) {
	ch1 <- <-ch2
	fmt.Println(s)
}

func main() {
	go worker1("work hard.")
	go worker2("just for fun!")
	<-ch1
}

// 考点：channel的执行顺序
// ch1 <- <-ch2会引起当前goroutine阻塞,
// <-ch1 会引起主线程阻塞,
// 所有goroutine都在阻塞，引起死锁
// 输出:
// work hard.
// fatal error: all goroutines are asleep - deadlock!
