package main

func p7() (prime int) {
	idx := 0

	for {
		if isPrime(prime) {
			idx++

			if idx == 10001 {
				return
			}
		}
		
		prime++
	}
}