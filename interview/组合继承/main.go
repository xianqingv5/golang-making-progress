package main

import "fmt"

type People struct{}
type Teacher struct {
	People
}

func (p *People) Walk() {
	fmt.Println("People Walk")
	p.Work()
}

func (p *People) Work() {
	fmt.Println("People Work")
}

func (t *Teacher) Work() {
	fmt.Println("Teacher Work")
}

func main() {
	t := Teacher{}
	t.Walk()
}

// 考点: golang的组合继承
// 虽然被组合的People类的方法升级成为外部Teacher组合类型的方法，
// 但p.Work()调用时接收者仍然是People, 因此打印"People Work"
// 输出:
// People Walk
// People Work
