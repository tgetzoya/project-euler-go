package main

func p12() int {

	idx := 1

	for {
		triangleNumber := TriangleNumber(idx)

		factorCount := FactorCount(triangleNumber, true)

		if factorCount > 500 {
			return triangleNumber
		}

		idx++
	}
}
