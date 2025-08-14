package main

// github.com/EndlessCheng/codeforces-go
func checkPowersOfThree(n int) bool {
	for ; n > 0; n /= 3 {
		if n%3 == 2 {
			return false
		}
	}
	return true
}
