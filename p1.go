package main

func p1() (sum int) {

	for idx := 1; idx < 1000; idx++ {
		if idx % 5 == 0 {
			sum += idx
		} else if idx % 3 == 0 {
			sum += idx
		}
	}

	return
}
