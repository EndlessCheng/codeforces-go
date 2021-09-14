package main

import "math/bits"

// todo O(log n) ?

// github.com/EndlessCheng/codeforces-go
func concatenatedBinary(n int) (ans int) {
	for i := 1; i <= n; i++ {
		ans = (ans<<bits.Len(uint(i)) | i) % (1e9 + 7)
	}
	return
}
