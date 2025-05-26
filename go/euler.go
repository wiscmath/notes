package main

import "math"

func euler1() float64 {
	return math.Exp(1.0)
}

func euler2(n int) float64 {
	return math.Pow(1+1/float64(n), float64(n))
}

func euler3(precision uint64) float64 {
	var n uint64
	var e float64

	for {
		term := 1 / float64(factorial(n))
		e += term
		if n > precision {
			break
		}
		n += 1
	}
	return e
}

func factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * factorial(n-1)
		return result
	}
	return 1
}
