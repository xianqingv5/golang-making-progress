package main

import (
	"golang-making-progress/并发模式/work"
	"log"
	"sync"
	"time"
)

// names提供了一组用来显示的名字
var names = []string{
	"steve",
	"bob",
	"mary",
	"therese",
	"jason",
}

// namePrinter 使用特定方式打印名字
type namePrinter struct {
	name string
}

func (m *namePrinter) Task() {
	log.Println(m.name)
	time.Sleep(time.Second)
}

// main是所有Go程序的入口
func main() {
	p := work.New(2) // 使用两个goroutine来创建工作池

	var wg sync.WaitGroup
	wg.Add(100 * len(names))

	for i := 0; i < 100; i++ {
		for _, name := range names { // 迭代names切片
			np := namePrinter{
				name: name,
			}

			go func() {
				p.Run(&np)
				wg.Done()
			}()
		}
	}

	wg.Wait()

	p.Shutdown()
}
