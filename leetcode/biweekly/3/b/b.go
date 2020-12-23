package main

// github.com/EndlessCheng/codeforces-go
func numKLenSubstrNoRepeats(s string, k int) (ans int) {
	has := ['z' + 1]bool{}
	l := 0
	for i, b := range s {
		for has[b] {
			has[s[l]] = false
			l++
		}
		if i+1-l >= k {
			ans++
		}
		has[b] = true
	}
	return
}
