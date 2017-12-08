package main

func p14() (val int) {
	longest := 0

	for idx := 1; idx < 1000000; idx++ {
		count := 1
		step := idx

		for step != 1 {
			if step%2 == 0 {
				step = step / 2
			} else {
				step = 3*step + 1
			}

			count++
		}

		if count > longest {
			longest = count
			val = idx
		}
	}

	return
}
