package main

import (
	"fmt"
	"reflect"
)

func hello() {
	fmt.Println("Hello, world!")
}

func main() {
	hl := hello
	hl()

	fv := reflect.ValueOf(hl)
	fmt.Println("fv is reflect.Func ?", fv.Kind() == reflect.Func)
	fv.Call(nil)
}
