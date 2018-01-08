package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func p18(fileName string) (val int) {
	tree := makeTree(fileName)

	for idx := len(tree) - 2; idx >= 0; idx-- {
		for jdx := 0; jdx < len(tree[idx]); jdx++ {
			if tree[idx+1][jdx] > tree[idx+1][jdx+1] {
				tree[idx][jdx] += tree[idx+1][jdx]
			} else {
				tree[idx][jdx] += tree[idx+1][jdx+1]
			}
		}
	}

	return tree[0][0]
}

func makeTree(fileName string) [][]int {
	inFile, err := os.Open(fileName)
	check(err)
	scanner := bufio.NewScanner(inFile)

	tree := make([][]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		splits := strings.Split(line, " ")

		row := make([]int, 0)

		for _, split := range splits {
			asNumber, err := strconv.Atoi(split)
			check(err)
			row = append(row, asNumber)
		}

		tree = append(tree, row)
	}

	return tree
}
