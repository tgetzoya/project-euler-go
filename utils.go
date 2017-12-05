package main

import (
	"sort"
)

func nextFibonacciSequence(num1 int, num2 int) (int,int) {
	hold := num1
	num1 = num2
	num2 = num1 + hold

	return num1, num2
}

func getFactors(num int) []int {
	factors := make([]int, 0)

	maxValue := num / 2

	for idx := 2; idx < maxValue; idx++ {
		if num % idx == 0 {
			maxValue = num / idx
			factors = append(factors,idx)
			factors = append(factors,maxValue)
		}
	}

	sort.Ints(factors)

	return factors
}

func isPrime(num int) bool {

	if num < 2 {
		return false
	}

	if num == 2 {
		return true
	}

	if num % 2 == 0 {
		return false
	}

	for idx := 3; idx * idx <= num; idx++ {
		if num % idx == 0 {
			return false
		}
	}

	return true
}

func isNumberPalindrome(num int) bool {
	digits := make([]int,0)

	for num > 10 {
		digits = append(digits,num % 10)
		num /= 10
	}

	digits = append(digits,num)
	length := len(digits) - 1

	for idx := 0; idx <= (length / 2); idx++ {
		if digits[idx] != digits[length - idx] {
			return false
		}
	}

	return true
}

func primesInList(list []int) (primes []int) {
	
	for _,val := range list {
		if isPrime(val) {
			primes = append(primes,val)
		}
	}

	return
}