package main

func p4() (largestPalindome int) {

	for idx := 999; idx >= 1; idx-- {
		for jdx := 999; jdx >= 1; jdx-- {
			num := idx * jdx

			if isNumberPalindrome(num) && largestPalindome < num {
				largestPalindome = num
			}
		}
	}

	return
}
