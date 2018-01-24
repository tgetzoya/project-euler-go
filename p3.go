package main

func p3() (maxPrimeFactor int) {

	factors := Factors(600851475143, false)

	for _, factor := range factors {
		if IsPrime(factor) && maxPrimeFactor < factor {
			maxPrimeFactor = factor
		}
	}

	return
}
