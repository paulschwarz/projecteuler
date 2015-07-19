package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/paulsschwarz/projecteuler/string"
)

func main() {
	numDigits := "3"
	fmt.Scanln(&numDigits)
	fmt.Println(numDigits)
	digits, err := strconv.Atoi(numDigits)
	if err != nil {
		fmt.Println(fmt.Errorf("Got error %q", err))
	} else {
		fmt.Println(largestPalindrome(digits))
	}
}

func largestPalindrome(digits int) (palindrome int, a int, b int) {
	max := int(math.Pow(10, float64(digits)) - 1)
	palindrome = max * max
	for palindrome > 0 {
		if string.IsPalindrome(strconv.Itoa(palindrome)) {
			a := max
			for a > 0 {
				b = palindrome / a
				if palindrome%a == 0 && string.IntLen(b) == digits {
					return palindrome, a, b
				}
				a--
			}
		}
		palindrome--
	}
	return
}
