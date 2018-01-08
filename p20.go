package main

import (
	"math/big"
)

func p20() (val int) {
	sum := 0
	big10 := big.NewInt(10)
	rem := big.NewInt(0)

	fact := bigFactorial(10000)

	for fact.Cmp(big.NewInt(1)) > 0 {
		rem.Mod(fact, big10)
		sum += int(rem.Int64())
		fact.Div(fact, big10)
	}

	return sum
}
