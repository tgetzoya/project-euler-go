package main

import (
	"sort"
	"strings"
)

func p22() (val int) {
	sum := 0

	names := readNamesFile()

	for idx, name := range names {
		lineVal := 0

		for _, chr := range name {
			lineVal += (int(chr) - 64)
		}

		sum += (lineVal * (idx + 1))
	}

	return sum
}

func readNamesFile() (names []string) {
	lines := readFile("/home/tgetzoyan/names.txt")

	for _, line := range lines {
		words := strings.Split(line, ",")

		for _, word := range words {
			names = append(names, strings.Replace(word, "\"", "", -1))
		}
	}

	sort.Strings(names)

	return names
}
