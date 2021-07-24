package main

// 模拟

// github.com/EndlessCheng/codeforces-go
func areOccurrencesEqual(s string) bool {
	cnt := [26]int{}
	for _, b := range s {
		cnt[b-'a']++
	}
	target := 0
	for _, c := range cnt {
		if c == 0 {
			continue
		}
		if target == 0 {
			target = c
		} else if c != target {
			return false
		}
	}
	return true
}
