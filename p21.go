package main

func p21() (val int) {
	sum := 0

	checked := make(map[int]bool)

	for da := 0; da < 10000; da++ {
		if checked[da] {
			continue
		}

		db := sumOfDivsors(da)
		dc := sumOfDivsors(db)

		if da == dc && da != db {
			sum += da + db

			checked[da] = true
			checked[db] = true
		}
	}

	return sum
}

func sumOfDivsors(num int) int {
	sum := 0

	facts := Factors(num, false)
	facts = append(facts, 1)

	for _, val := range facts {
		sum += val
	}

	return sum
}
