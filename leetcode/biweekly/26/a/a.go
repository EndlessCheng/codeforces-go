package main

// github.com/EndlessCheng/codeforces-go
func maxPower(s string) (ans int) {
	for i, n := 0, len(s); i < n; {
		st := i
		for ; i < n && s[i] == s[st]; i++ {
		}
		if i-st > ans {
			ans = i - st
		}
	}
	return
}
