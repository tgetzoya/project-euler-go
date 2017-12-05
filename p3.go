package main

func p3() (maxPrimeFactor int) {

	factors := getFactors(600851475143)

	for _,factor := range factors {
		if isPrime(factor) && maxPrimeFactor < factor {
			maxPrimeFactor = factor
		}
	}

	return
}
