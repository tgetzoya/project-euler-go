package main

import (
	"bufio"
	"math"
	"math/big"
	"os"
	"sort"
)

func nextFibonacciSequence(num1 int, num2 int) (int, int) {
	hold := num1
	num1 = num2
	num2 = num1 + hold

	return num1, num2
}

func getSortedFactors(num int, includeSelf bool) (factors []int) {
	factors = getFactors(num, includeSelf)
	sort.Ints(factors)
	return
}

func getProperDivisors(num int) []int {
	divisors := getFactors(num, false)
	divisors = append(divisors, 1)

	return divisors
}

func getFactors(num int, includeSelf bool) []int {
	factorMap := make(map[int]bool, 0)

	if includeSelf {
		factorMap[1] = true
		factorMap[num] = true
	}

	maxValue := num / 2

	for idx := 2; idx < maxValue; idx++ {
		if num%idx == 0 {
			maxValue = num / idx
			factorMap[idx] = true
			factorMap[maxValue] = true
		}
	}

	factors := make([]int, 0)

	for k := range factorMap {
		factors = append(factors, k)
	}

	return factors
}

func getFactorCount(num int, includeSelf bool) (count int) {

	if includeSelf {
		count += 2
	}

	for idx := 2; idx < int(math.Sqrt(float64(num))); idx++ {
		if num%idx == 0 {
			count += 2
		}
	}

	return
}

func isPrime(num int) bool {

	if num < 2 {
		return false
	}

	if num == 2 {
		return true
	}

	if num%2 == 0 {
		return false
	}

	for idx := 3; idx*idx <= num; idx++ {
		if num%idx == 0 {
			return false
		}
	}

	return true
}

func isNumberPalindrome(num int) bool {
	digits := make([]int, 0)

	for num > 10 {
		digits = append(digits, num%10)
		num /= 10
	}

	digits = append(digits, num)
	length := len(digits) - 1

	for idx := 0; idx <= (length / 2); idx++ {
		if digits[idx] != digits[length-idx] {
			return false
		}
	}

	return true
}

func primesInList(list []int) (primes []int) {

	for _, val := range list {
		if isPrime(val) {
			primes = append(primes, val)
		}
	}

	return
}

func triangleNumber(position int) (triangleNumber int) {
	for idx := 1; idx <= position; idx++ {
		triangleNumber += idx
	}

	return
}

func factorial(num int64) (result int64) {

	if num < 0 {
		return -1
	}

	if num == 0 {
		return 1
	}

	result = num

	for idx := num - 1; idx > 1; idx-- {
		result *= idx
	}

	return
}

func bigFactorial(num int64) *big.Int {
	if num < 0 {
		return big.NewInt(-1)
	}

	result := big.NewInt(1)

	if num > 0 {
		for idx := num - 1; idx > 1; idx-- {
			result = result.Mul(result, big.NewInt(idx))
		}
	}

	return result
}

func readFile(fileName string) (list []string) {
	inFile, err := os.Open(fileName)
	check(err)
	scanner := bufio.NewScanner(inFile)

	for scanner.Scan() {
		list = append(list, scanner.Text())
	}

	return list
}
