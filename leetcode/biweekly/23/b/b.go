package main

// github.com/EndlessCheng/codeforces-go
func canConstruct(s string, k int) (ans bool) {
	cnt := [26]int{}
	for _, b := range s {
		cnt[b-'a'] ^= 1
	}
	sum := 0
	for _, v := range cnt[:] {
		sum += v
	}
	return sum <= k && k <= len(s)
}
