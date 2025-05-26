package main

import (
	"fmt"
	"testing"
)

func TestEuler(t *testing.T) {
	fmt.Println(euler1())
	fmt.Println(euler2(1000))
	fmt.Println(euler3(1000))
}

func BenchmarkEuler1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		euler1()
	}
}

func BenchmarkEuler2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		euler2(1000000)
	}
}

func BenchmarkEuler3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		euler3(1000)
	}
}
