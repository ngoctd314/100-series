package main

import (
	"fmt"
	"strings"
)

type store struct {
	m map[int]*int
}

func (s store) put(v []int) {
	fmt.Println(strings.Repeat("~", 10))
	for k := range v {
		fmt.Printf("v[%d] = %p\n", k, &v[k])
		s.m[k] = &v[k]
	}
}

func main() {
	m := map[int]bool{
		0: true,
		1: false,
		2: true,
	}

	for k := range m {
		m[10+k] = true
	}

	fmt.Println(m)
}
