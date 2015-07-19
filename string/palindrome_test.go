package string

import "testing"

type testPairPalindrome struct {
	input  string
	result bool
}

var testPairPalindromes = []testPairPalindrome{
	{"a", true},
	{"bb", true},
	{"abc", false},
	{"aba", true},
	{"abba", true},
	{"1", true},
	{"11", true},
	{"111", true},
	{"101", true},
	{"10", false},
	{"100", false},
}

func TestIsPalindrome(t *testing.T) {
	for _, test := range testPairPalindromes {
		actual := IsPalindrome(test.input)
		if actual != test.result {
			t.Error(
				"For", test.input,
				"expected", test.result,
				"got", actual,
			)
		}
	}
}
