package main

func p17() (val int) {
	sum := 0

	for idx := 1; idx <= 1000; idx++ {
		sum += numberToWordCount(idx)
	}

	return sum
}

func numberToWordCount(num int) int {
	if num == 1000 {
		return 11
	}

	sum := 0

	if num >= 100 {
		/* Hundred = 7 */
		sum += digitNameCount(num/100) + 7
		num = num % 100

		/* Add the word "And" */
		if num > 0 {
			sum += 3
		} else {
			return sum
		}
	}

	if num > 19 {
		bob := num / 10

		if bob == 2 || bob == 3 || bob == 8 || bob == 9 {
			sum += 6
		} else if bob == 4 || bob == 5 || bob == 6 {
			sum += 5
		} else {
			sum += 7
		}

		num = num % 10
	}

	sum += digitNameCount(num)

	return sum
}

func digitNameCount(digit int) int {
	if digit == 0 {
		return 0
	}

	size := 0

	if digit == 1 || digit == 2 || digit == 6 || digit == 10 {
		size = 3
	} else if digit == 3 || digit == 7 || digit == 8 {
		size = 5
	} else if digit == 11 || digit == 12 {
		size = 6
	} else if digit == 15 || digit == 16 {
		size = 7
	} else if digit == 13 || digit == 14 || digit == 18 || digit == 19 {
		size = 8
	} else if digit == 17 {
		size = 9
	} else {
		size = 4
	}

	return size
}
