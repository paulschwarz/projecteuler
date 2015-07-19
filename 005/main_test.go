package main

import "testing"

func TestGetResult(t *testing.T) {

	type testPair struct {
		limit  int
		result int
	}

	var testPairs = []testPair{
		{5, 60},
		{10, 2520},
	}

	for _, test := range testPairs {
		actual := getResult(test.limit)
		if actual != test.result {
			t.Error(
				"For", test.limit,
				"expected", test.result,
				"got", actual,
			)
		}
	}
}
