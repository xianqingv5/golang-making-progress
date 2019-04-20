package main

import (
	"fmt"
	"sort"
)

func main() {
	// To create a map as input
	m := make(map[float64]string)
	m[1.2] = "a"
	m[2.3] = "c"
	m[0.1] = "b"

	// To store the keys in slice in sorted order
	var keys []float64
	for k := range m {
		keys = append(keys, k)
	}
	sort.Float64s(keys)

	// To perform the opertion you want
	for _, k := range keys {
		fmt.Println("Key:", k, "Value:", m[k])
	}
}
