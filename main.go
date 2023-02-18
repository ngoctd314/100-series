package main

import (
	"fmt"
	"log"
	_ "net/http/pprof"
	"runtime"
)

func main() {
	a := [3]int{0, 1, 2}
	for _, v := range &a {
		a[2] = 10
		log.Println(v)
	}
	fmt.Println(a)
}

func printAlloc() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d KB\n", m.Alloc/1024)
}
