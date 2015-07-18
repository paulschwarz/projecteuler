package main

import (
	"fmt"

	"github.com/paulsschwarz/projecteuler/math"
)

func main() {
	fmt.Println(SumEvenFibonacciNumbers(4000000))
}

// SumEvenFibonacciNumbers sums only the even Fibonacci numbers whose values do
// not exceed the given limit.
func SumEvenFibonacciNumbers(limit int) (sum int) {
	fibonacci := math.Fibonacci()
	i := 0
	for {
		result := fibonacci(i)
		if result > limit {
			break
		} else if result%2 == 0 {
			sum += result
		}
		i++
	}
	return
}
