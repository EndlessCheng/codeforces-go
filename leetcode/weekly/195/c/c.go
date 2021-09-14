package main

import "sort"

const mod int = 1e9 + 7

var pow2 = [1e5 + 5]int{1}

func init() {
	for i := 1; i <= 1e5+4; i++ {
		pow2[i] = pow2[i-1] << 1 % mod
	}
}

func numSubseq(a []int, target int) (ans int) {
	sort.Ints(a)
	for r, v := range a {
		l := sort.SearchInts(a[:r], target-v+1)
		if a[l]+a[r] > target {
			l--
		}
		if l == r {
			ans = (ans + pow2[l]) % mod
		} else if l >= 0 {
			ans = (ans + pow2[r-l-1]*(pow2[l+1]-1)) % mod
		}
	}
	return
}
