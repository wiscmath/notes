package main

import (
	"math"
	"math/big"
)

func euler1() float64 {
	return math.Exp(1.0)
}

func euler2(n int) float64 {
	return math.Pow(1+1/float64(n), float64(n))
}

func euler3(iterations uint64) float64 {
	var n uint64
	var e float64

	for {
		term := float64(1) / float64(factorial(n))
		e += term
		if n > iterations {
			break
		}
		n += 1
	}
	return e
}

func factorial(n uint64) (result uint64) {

	if n == 0 || n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func factorial2(n big.Int) (result *big.Int) {
	if n.Cmp(big.NewInt(0)) == 0 || n.Cmp(big.NewInt(1)) == 0 {
		return big.NewInt(1)
	}
	sub := n.Sub(&n, big.NewInt(1))
	return n.Mul(&n, factorial2(*sub))
}
