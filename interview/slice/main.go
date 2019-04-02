package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}

func main() {
	s := []string{"w", "a", "g"}
	d := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	fmt.Println(*d)
}

// 切片的内部实现
// https://zhuanlan.zhihu.com/p/28399762
