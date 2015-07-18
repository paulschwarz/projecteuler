package math

import "testing"

type fibTest struct {
	index  int
	result int
}

var fibTests = []fibTest{
	{0, 0},
	{1, 1},
	{2, 1},
	{3, 2},
	{4, 3},
	{5, 5},
	{6, 8},
	{7, 13},
	{8, 21},
	{9, 34},
	{10, 55},
}

func TestFibonacci(t *testing.T) {
	fibonacci := Fibonacci()
	for _, test := range fibTests {
		actual := fibonacci(test.index)
		if actual != test.result {
			t.Error(
				"For", test.index,
				"expected", test.result,
				"got", actual,
			)
		}
	}
}
