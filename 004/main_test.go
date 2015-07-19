package main

import "testing"

func TestLargestPalindrome(t *testing.T) {
	expected := 9009
	expectedA := 99
	expectedB := 91
	actual, actualA, actualB := largestPalindrome(2)
	if actual != expected || actualA != expectedA || actualB != expectedB {
		t.Error(
			"For", 2,
			"expected", expected, "using", expectedA, "and", expectedB,
			"got", actual, "using", actualA, "and", actualB,
		)
	}
}
