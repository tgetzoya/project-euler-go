package main

import (
	"math"
)

func p9() int {
	for idx := 999; idx > 0 ; idx-- {
		for jdx := idx - 1; jdx > 0; jdx-- {
			c := math.Sqrt(float64(idx * idx)+float64(jdx * jdx))

			if (c != math.Trunc(c)) {
				continue
			}

			kdx := int(c)

			if (idx + jdx + kdx == 1000) {
				return idx * jdx * kdx
			}
		}
	}

	return 1001
}