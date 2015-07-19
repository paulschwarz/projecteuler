package string

import "strconv"

// IntLen returns the string length of the integer n.
func IntLen(n int) int {
	return len(strconv.Itoa(n))
}
