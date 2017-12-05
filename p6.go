package main

func p6() int {
	return squareOfSumes(100) - sumOfSquares(100)
}

func sumOfSquares(num int) (sum int) {
	for idx := 1; idx <= num; idx++ {
		sum += idx * idx;
	}

	return
}

func squareOfSumes(num int) (sum int) {
	for idx := 1; idx <= num; idx++ {
		sum += idx;
	}

	return sum * sum
}