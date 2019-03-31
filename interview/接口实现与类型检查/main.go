package main

import "fmt"

// People people
type People interface {
	Speak(string) string
}

// Stduent stduent
type Stduent struct{}

// Speak speak
func (stu *Stduent) Speak(name string) (talk string) {
	if name == "xiaoming" {
		talk = "xiaoming is a good boy."
	} else {
		talk = "who's that guy?"
	}
	return
}

func main() {
	// var peo People = Stduent{}
	var peo People = &Stduent{}
	fmt.Println(peo.Speak("hello"))
}

// 考点: 接口实现与类型检查
// 对于赋值给接口的类型，会有严格的类型检查，不允许做A到A*的转换
// 类型A的方法集（method set）是类型*A的一个子集
// 输出:
// ./p7.go:23: cannot use Stduent literal (type Stduent) as type People in assignment:
//	Stduent does not implement People (Speak method has pointer receiver)
