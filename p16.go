package main

import (
	"math/big"
	"strconv"
)

func p16() (val int) {
	largeNum := new(big.Int).Exp(big.NewInt(2), big.NewInt(1000), nil).String()
	sum := 0

	for _, c := range largeNum {
		ch, _ := strconv.Atoi(string(c))
		sum += ch
	}

	return sum
}
