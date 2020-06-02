package main

import "strconv"

func hasAllCodes(s string, k int) (ans bool) {
	has := map[int64]bool{}
	for i := k; i <= len(s); i++ {
		v, _ := strconv.ParseInt(s[i-k:i], 2, 64)
		has[v] = true
	}
	return len(has) == 1<<k
}
