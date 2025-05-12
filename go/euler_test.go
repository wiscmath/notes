package main

import (
	"testing"
)

func BenchmarkEuler1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		euler1()
	}
}
