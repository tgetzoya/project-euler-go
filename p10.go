package main

func p10() (sum int) {

	for idx := 0; idx < 2000000; idx++ {
		if IsPrime(idx) {
			sum += idx
		}
	}

	return
}
