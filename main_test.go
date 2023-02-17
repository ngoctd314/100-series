package main

import "testing"

func BenchmarkMapWithout(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mapWithoutInitialize()
	}
}

func BenchmarkMapWith(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mapWithInitialize()
	}
}
