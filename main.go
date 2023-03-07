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
	s := store{
		m: make(map[int]*int),
	}
	v := []int{1, 2, 3}
	for k := range v {
		fmt.Printf("v[%d] = %p\n", k, &v[k])
	}

	s.put(v)
	for _, v := range s.m {
		fmt.Println(*v)
	}
}
