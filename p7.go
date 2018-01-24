package main

func p7() (prime int) {
	idx := 0

	for {
		if IsPrime(prime) {
			idx++

			if idx == 10001 {
				return
			}
		}

		prime++
	}
}
