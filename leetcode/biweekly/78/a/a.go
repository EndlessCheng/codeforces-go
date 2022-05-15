package main

import "strconv"

// github.com/EndlessCheng/codeforces-go
func divisorSubstrings(num, k int) (ans int) {
	s := strconv.Itoa(num)
	for i := k; i <= len(s); i++ {
		v, _ := strconv.Atoi(s[i-k : i])
		if v > 0 && num%v == 0 {
			ans++
		}
	}
	return
}
