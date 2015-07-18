package math

import "testing"

type sumTest struct {
	inputs []int
	result int
}

var sumTests = []sumTest{
	{[]int{}, 0},
	{[]int{1}, 1},
	{[]int{1, 2, 3, 4}, 10},
	{[]int{1, 2, 3, 4, -5}, 5},
}

func TestSum(t *testing.T) {
	for _, test := range sumTests {
		actual := Sum(test.inputs)
		if actual != test.result {
			t.Error(
				"For", test.inputs,
				"expected", test.result,
				"got", actual,
			)
		}
	}
}

type multiplesTest struct {
	limit   int
	results []int
}

var multiplesTests = []multiplesTest{
	{6, []int{3, 5}},
	{10, []int{3, 5, 6, 9}},
	{12, []int{3, 5, 6, 9, 10}},
}

func TestMultiples(t *testing.T) {
	for _, test := range multiplesTests {
		actual := Multiples(test.limit)
		if !intArraysMatch(actual, test.results) {
			t.Error(
				"For", test.limit,
				"expected", test.results,
				"got", actual,
			)
		}
	}
}

func intArraysMatch(as []int, bs []int) bool {
	if len(as) != len(bs) {
		return false
	}

	for i, a := range as {
		if bs[i] != a {
			return false
		}
	}

	return true
}
