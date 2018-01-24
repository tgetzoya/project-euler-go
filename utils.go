package main

import (
	"bufio"
	"math"
	"math/big"
	"os"
	"sort"
)

//Fibonacci using big.Int
func Fibonacci(term int) *big.Int {
	if term < 1 {
		panic("Position must be 1 or greater")
	}

	if term < 3 {
		return big.NewInt(1)
	}

	if term == 3 {
		return big.NewInt(2)
	}

	n := big.NewInt(1)
	m := big.NewInt(2)
	hold := big.NewInt(0)

	for idx := 3; idx < term; idx++ {
		hold.Set(m)
		m = m.Add(m, n)
		n.Set(hold)
	}

	return m
}

//NextFibonacciSequence using int
func NextFibonacciSequence(num1 int, num2 int) (int, int) {
	hold := num1
	num1 = num2
	num2 = num1 + hold

	return num1, num2
}

//SortedFactors gets the factors of a number and sorts them
func SortedFactors(num int, includeSelf bool) (factors []int) {
	factors = Factors(num, includeSelf)
	sort.Ints(factors)
	return
}

//ProperDivisors is the same as SortedFactors, but includes the 1
func ProperDivisors(num int) []int {
	divisors := Factors(num, false)
	divisors = append(divisors, 1)

	return divisors
}

//Factors gets the factors for an integer
func Factors(num int, includeSelf bool) []int {
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

//FactorCount returns the number of factors for an integer
func FactorCount(num int, includeSelf bool) (count int) {

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

//IsPrime check the integer to see if it is prime
func IsPrime(num int) bool {

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

//IsNumberPalindrome checks to see if the number is a palindrome
func IsNumberPalindrome(num int) bool {
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

//PrimesInList checks an array of ints to see if there are any prime numbers
func PrimesInList(list []int) (primes []int) {

	for _, val := range list {
		if IsPrime(val) {
			primes = append(primes, val)
		}
	}

	return
}

//TriangleNumber returns a number at the position in a triangle
func TriangleNumber(position int) (triangleNumber int) {
	for idx := 1; idx <= position; idx++ {
		triangleNumber += idx
	}

	return
}

//Factorial does something....
func Factorial(num int64) (result int64) {

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

//BigFactorial is the same as above, but has support for extremely large numbers
func BigFactorial(num int64) *big.Int {
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

//ReadFile reads a file and returns the contents as a string slice
func ReadFile(fileName string) (list []string) {
	inFile, err := os.Open(fileName)
	check(err)
	scanner := bufio.NewScanner(inFile)

	for scanner.Scan() {
		list = append(list, scanner.Text())
	}

	return list
}
