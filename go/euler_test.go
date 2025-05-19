package main

import (
	"fmt"
	"testing"
)

func TestEuler(t *testing.T) {
	fmt.Println(euler1())
	fmt.Println(euler2(1000))
}

func BenchmarkEuler1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		euler1()
	}
}
