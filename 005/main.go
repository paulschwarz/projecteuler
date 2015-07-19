package main

import (
	"fmt"
	"strconv"

	"github.com/paulsschwarz/projecteuler/math"
)

func main() {
	numDigits := "20"
	fmt.Scanln(&numDigits)
	fmt.Println(numDigits)
	digits, err := strconv.Atoi(numDigits)
	if err != nil {
		fmt.Println(fmt.Errorf("Got error %q", err))
	} else {
		fmt.Println(getResult(digits))
	}
}

func getResult(limit int) int {
	result, i := 1, 2
	for ; i <= limit; i++ {
		result = math.LCM(result, i)
	}

	return result
}
