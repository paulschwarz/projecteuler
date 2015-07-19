package math

// Sum returns the sum of an array of ints.
func Sum(nums []int) (total int) {
	for _, num := range nums {
		total += num
	}
	return
}

// Multiples finds the multiples of 3 or 5 in the range 0 to limit.
func Multiples(limit int) (multiples []int) {
	i := 1
	for i < limit {
		if i%15 == 0 || i%5 == 0 || i%3 == 0 {
			multiples = append(multiples, i)
		}
		i++
	}
	return
}

// GCD finds the greatest common divisor of two integers.
func GCD(a, b int) int {
	c := a % b
	if c == 0 {
		return b
	}
	return GCD(b, c)
}

// LCM finds the lowest common multiple of two integers.
func LCM(a, b int) int {
	return a * b / GCD(a, b)
}
