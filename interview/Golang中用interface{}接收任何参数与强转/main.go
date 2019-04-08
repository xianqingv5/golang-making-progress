package main

import "fmt"

// func Any(v interface{}) {
// 	v1 := int(v)
// 	println(v1)
// }
// cannot convert v (type interface {}) to type int: need type assertion

// Any any
func Any(v interface{}) {
	vv := v.(string)
	fmt.Println(vv)

	if v2, ok := v.(string); ok {
		println(v2)
	} else if v3, ok2 := v.(int); ok2 {
		println(v3)
	}
}

func main() {
	Any(2)
	Any("666")
}
