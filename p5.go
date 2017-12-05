package main

func p5() int {
	num := 11

	for {
		if isEvenlyDivisible(num,1,20) {
			return num
		}

		num++
	}
}

func isEvenlyDivisible(num, min, max int) bool {

	for idx := min; idx <= max; idx++ {
		if num % idx != 0 {
			return false
		}
	}

	return true
}
