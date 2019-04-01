package main

import "fmt"

type data struct {
	z int
}

func check(a data) bool {
	return a == data{}
}

func main() {

	sn1 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}
	sn2 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}

	if sn1 == sn2 {
		fmt.Println("sn1 == sn2")
	}

	sm1 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}
	sm2 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}

	if sm1 == sm2 {
		fmt.Println("sm1 == sm2")
	}

	fmt.Println("----------------")

	var a struct {
		x    int
		y    int
		data struct {
			z int
		}
	}
	a.data = struct{ z int }{100}
	println(check(a.data))
}
