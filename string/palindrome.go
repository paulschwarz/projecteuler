package string

// IsPalindrome tests if a string is a palindrome.
func IsPalindrome(input string) bool {
	l := len(input)
	i := 0
	for i < l/2 {
		if input[i] != input[l-1-i] {
			return false
		}
		i++
	}
	return true
}
