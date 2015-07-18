package main

import (
	"fmt"
	"math"

	_math "github.com/paulsschwarz/projecteuler/math"
)

func main() {
	fmt.Println(largestPrimeFactor(600851475143))
}

func largestPrimeFactor(n int) (result int) {
	i := int(math.Ceil(math.Sqrt(float64(n))))
	for i > 1 {
		if n%i == 0 && _math.IsPrime(i) {
			return i
		}
		i--
	}
	return
}
