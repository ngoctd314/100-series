package main

import "fmt"

func main() {
	v := []int{1, 2, 3}
	for i := 0; i < len(v); i++ {
		v = append(v, i)
	}
	fmt.Println(v)
}

func fn() int {
	fmt.Println("RUN")
	return 5
}
