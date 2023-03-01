package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	m := make(map[int]*person)
	m[0] = &person{}
	fmt.Println(m[0])
}

func f() {
	fmt.Println("a")
	panic("foo")
}
