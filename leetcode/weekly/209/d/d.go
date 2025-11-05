package main

import "math/bits"

// github.com/EndlessCheng/codeforces-go
func minimumOneBitOperations1(n int) int {
	if n == 0 {
		return 0
	}
	k := bits.Len(uint(n))
	return 1<<k - 1 - minimumOneBitOperations(n-1<<(k-1))
}

func minimumOneBitOperations(n int) (ans int) {
	for n > 0 {
		lb := n & -n
		ans = lb<<1 - 1 - ans
		n ^= lb
	}
	return
}

func minimumOneBitOperations2(n int) (ans int) {
	for ; n > 0; n >>= 1 {
		ans ^= n
	}
	return
}
