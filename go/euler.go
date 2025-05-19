package main

import "math"

func euler1() float64 {
	return math.Exp(1.0)
}

func euler2(n int) float64 {
	return math.Pow(1+1/float64(n), float64(n))
}
