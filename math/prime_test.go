package math

import "testing"

type primeTest struct {
	number int
	result bool
}

var primeTests = []primeTest{
	{0, false},
	{1, false},
	{2, true},
	{3, true},
	{4, false},
	{5, true},
	{6, false},
	{7, true},
	{8, false},
	{9, false},
	{10, false},
	{11, true},
	{12, false},
	{13, true},
	{14, false},
	{15, false},
}

func TestIsPrime(t *testing.T) {
	for _, test := range primeTests {
		actual := IsPrime(test.number)
		if actual != test.result {
			t.Error(
				"For", test.number,
				"expected", test.result,
				"got", actual,
			)
		}
	}
}
