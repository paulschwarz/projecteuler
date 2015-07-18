package main

import "testing"

type testPair struct {
	limit  int
	result int
}

var testPairs = []testPair{
	{3, 2},
	{8, 10},
	{100, 44},
}

func TestSumEvenFibonacciNumbers(t *testing.T) {
	for _, test := range testPairs {
		actual := SumEvenFibonacciNumbers(test.limit)
		if actual != test.result {
			t.Error(
				"For", test.limit,
				"expected", test.result,
				"got", actual,
			)
		}
	}
}
