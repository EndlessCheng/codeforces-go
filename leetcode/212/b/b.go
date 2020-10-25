package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func checkArithmeticSubarrays(a []int, ls, rs []int) (ans []bool) {
o:
	for i, l := range ls {
		b := append([]int(nil), a[l:rs[i]+1]...)
		sort.Ints(b)
		for i := 2; i < len(b); i++ {
			if b[i]-b[i-1] != b[1]-b[0] {
				ans = append(ans, false)
				continue o
			}
		}
		ans = append(ans, true)
	}
	return
}
