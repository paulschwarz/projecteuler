package string

import "testing"
import "math"

type testPairIntLen struct {
	input  int
	result int
}

var testPairIntLens = []testPairIntLen{
	{0, 1},
	{1, 1},
	{2, 1},
	{11, 2},
	{111, 3},
	{int(math.Pow(10, 3)), 4},
}

func TestIntLen(t *testing.T) {
	for _, test := range testPairIntLens {
		actual := IntLen(test.input)
		if actual != test.result {
			t.Error(
				"For", test.input,
				"expected", test.result,
				"got", actual,
			)
		}
	}
}
