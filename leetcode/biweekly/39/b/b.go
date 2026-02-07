package main

import "strings"

// github.com/EndlessCheng/codeforces-go
func minimumDeletions1(s string) int {
	del := strings.Count(s, "a")
	ans := del
	for _, c := range s {
		// 'a' -> -1    'b' -> 1
		del += int((c-'a')*2 - 1)
		ans = min(ans, del)
	}
	return ans
}

func minimumDeletions(s string) int {
	f, cntB := 0, 0
	for _, c := range s {
		if c == 'b' { // f 值不变
			cntB++
		} else {
			f = min(f+1, cntB)
		}
	}
	return f
}
