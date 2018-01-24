package main

func p2() (sum int) {

	i := 1
	j := 2

	sum += j

	for j < 4000000 {
		i, j = NextFibonacciSequence(i, j)

		if j%2 == 0 {
			sum += j
		}
	}

	return
}
