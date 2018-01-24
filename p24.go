package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/gitchander/permutation"
)

func p24() string {
	digits := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	permutations := permutation.New(permutation.IntSlice(digits))

	list := make([]string, 0)

	for permutations.Scan() {
		list = append(list, strings.Trim(strings.Join(strings.Split(fmt.Sprint(digits), " "), ""), "[]"))
	}

	sort.Strings(list)

	return list[999999]
}
