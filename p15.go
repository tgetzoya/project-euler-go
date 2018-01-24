package main

func p15() (val int64) {
	return binomialCoefficient(int64(40), int64(20))
}

func binomialCoefficient(n, k int64) int64 {
	if n < k {
		hold := n
		n = k
		k = hold
	}

	return Factorial(n) / (Factorial(k) * Factorial(n-k))
}
