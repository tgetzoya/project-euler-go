package main

func p25() int {
	idx := 0
	sum := 0

	for sum < 1000 {
		idx++
		num := Fibonacci(idx)
		sum = len(num.String())
	}

	return idx
}
