package main

func longestPrefix(s string) (ans string) {
	n := len(s)
	maxMatchLengths := make([]int, n)
	maxLength := 0
	for i := 1; i < n; i++ {
		c := s[i]
		for maxLength > 0 && s[maxLength] != c {
			maxLength = maxMatchLengths[maxLength-1]
		}
		if s[maxLength] == c {
			maxLength++
		}
		maxMatchLengths[i] = maxLength
	}
	return s[:maxMatchLengths[n-1]]
}
