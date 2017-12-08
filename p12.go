package main

func p12() int {

	idx := 1

	for {
		triangleNumber := triangleNumber(idx)

		factorCount := getFactorCount(triangleNumber, true)

		if factorCount > 500 {
			return triangleNumber
		}

		idx++
	}
}
