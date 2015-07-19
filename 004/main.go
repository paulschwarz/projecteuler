package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/paulsschwarz/projecteuler/string"
)

func main() {
	fmt.Println(largestPalindrome(3))
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
