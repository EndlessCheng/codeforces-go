package main

// github.com/EndlessCheng/codeforces-go
func minSwaps(s string) int {
	c := 0
	for _, b := range s {
		if b == '[' || c == 0 {
			c++
		} else {
			c--
		}
	}
	return c / 2
}
