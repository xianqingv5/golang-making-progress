package main

import "fmt"

type Strategier interface {
	Compute(num1, num2 int) int
}

type Division struct {
}

func (p Division) Compute(num1, num2 int) int {
	defer func() {
		if f := recover(); f != nil {
			fmt.Println(f)
			return
		}
	}()

	if num2 == 0 {
		panic("num2 must not be 0!")
	}

	return num1 / num2
}

func main() {

}
