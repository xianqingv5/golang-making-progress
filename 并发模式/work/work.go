package work

import "sync"

// Worker 必须满足接口类型，才能使用工作池
type Worker interface {
	Task()
}

// Pool 提供一个goroutine池，这个池可以完成任何已提交的Worker任务
type Pool struct {
	work chan Worker
	wg   sync.WaitGroup
}

// New 创建一个新工作池
func New(maxGoroutines int) *Pool {
	p := Pool{
		work: make(chan Worker),
	}

	p.wg.Add(maxGoroutines)
	for i := 0; i < maxGoroutines; i++ {
		go func() {
			for w := range p.work { // for range循环会一直阻塞，直到从work通道收到一个Worker接口值。如果收到一个值，就会执行这个值的Task方法。一旦work通道被关闭，for range循环就会结束，并调用WaitGroup的Done方法。然后goroutine终止。
				w.Task()
			}
			p.wg.Done()
		}()
	}

	return &p
}

// Run 提交工作到工作池
func (p *Pool) Run(w Worker) {
	p.work <- w
}

// Shutdown 等待所有goroutine停止工作
func (p *Pool) Shutdown() {
	close(p.work)
	p.wg.Wait()
}
