package main

import (
	"fmt"
	"math/big"
)

func p26() int {
	sum := 0

	bob := big.NewFloat(1.0 / 983)
	nancy := bob.Text('f', 1000)

	fmt.Println(nancy)

	return sum
}
