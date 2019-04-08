package main

import "fmt"

const (
	x = iota
	y
	z = "zz"
	k
	p = iota
)

// ...
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
)

const (
	_, _ = iota, iota * 10 // 0, 0 * 10
	a, b                   // 1, 1 * 10
	c, d                   // 2, 2 * 10
)

func main() {
	fmt.Println(x, y, z, k, p)
}
