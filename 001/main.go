package main

import (
	"fmt"

	"github.com/paulsschwarz/projecteuler/math"
)

func main() {
	multiples := math.Multiples(1000)
	fmt.Println(math.Sum(multiples))
}
