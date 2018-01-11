package main

func p23() (val int) {
	sum := 0

	isAbundant := make([]bool, 28123)

	for idx := 12; idx < 28123; idx++ {
		divisors := getProperDivisors(idx)
		sumOfDivisors := 0

		for _, div := range divisors {
			sumOfDivisors += div
		}

		if sumOfDivisors > idx {
			isAbundant[idx] = true
		}
	}

	abunbantNumbers := make([]int, 0)

	for idx, val := range isAbundant {

		if val {
			abunbantNumbers = append(abunbantNumbers, idx)
		}
	}

	for idx := 28123; idx > 0; idx-- {
		notFromSumAbundant := true
		for _, an := range abunbantNumbers {
			pos := idx - an
			if pos > 0 && pos < len(isAbundant) && isAbundant[pos] {
				notFromSumAbundant = false
				break
			}
		}

		if notFromSumAbundant {
			sum += idx
		}
	}

	return sum
}
