package main

import (
	"fmt"
	"sort"
)

func p24() (val int) {
	sum := 0

	chrs := []string{"0", "1", "2"}

	permutations := make(map[string]bool, 0)

	for idx := range chrs {
		perms := copyStringArray(chrs)

		for jdx := range perms {

		}
	}

	keys := make([]string, 0)

	for idx := range permutations {
		keys = append(keys, idx)
	}

	sort.Strings(keys)

	fmt.Println(keys)

	return sum
}

func copyStringArray(src []string) []string {
	dest := make([]string, 0)

	for _, val := range src {
		dest = append(dest, val)
	}

	return dest
}
