package main

import "testing"

type testPair struct {
	limit  int
	result int
}

var testPairs = []testPair{
	{2, 2},
	{4, 2},
	{6, 3},
	{13195, 29},
}

func TestLargestPrimeFactor(t *testing.T) {
	for _, test := range testPairs {
		actual := largestPrimeFactor(test.limit)
		if actual != test.result {
			t.Error(
				"For", test.limit,
				"expected", test.result,
				"got", actual,
			)
		}
	}
}
