package math

// Fibonacci returns a function which computes the nth number in the Fibonacci
// sequence. It makes efficient of use memoization.
func Fibonacci() func(n int) int {
	memo := map[int]int{
		0: 0,
	}
	var fib func(n int) int
	fib = func(n int) int {
		if result, ok := memo[n]; !ok {
			if n > 1 {
				result = fib(n-1) + fib(n-2)
			} else {
				result = n
			}
			memo[n] = result
		}
		return memo[n]
	}
	return fib
}
