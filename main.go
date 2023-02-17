package main

import (
	"fmt"
	_ "net/http/pprof"
	"runtime"
)

func main() {

}

func printAlloc() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d KB\n", m.Alloc/1024)
}

func mapWithoutInitialize() {
	m := make(map[int]int)
	for i := 0; i < 1e6; i++ {
		m[i] = i
	}
}

func mapWithInitialize() {
	m := make(map[int]int, 1e6)
	for i := 0; i < 1e6; i++ {
		m[i] = i
	}
}
