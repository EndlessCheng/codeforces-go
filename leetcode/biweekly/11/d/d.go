package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func maximizeSweetness(a []int, k int) (ans int) {
	return sort.Search(1e9+1, func(low int) bool {
		cnt := 0
		for i, n := 0, len(a); i < n; {
			s := 0
			for ; i < n && s < low; i++ {
				s += a[i]
			}
			// 循环结束后先判断是否达到要求
			if s < low {
				break
			}
			if cnt++; cnt > k {
				return false
			}
		}
		return true
	}) - 1
}
