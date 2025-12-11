package main

import "slices"

// github.com/EndlessCheng/codeforces-go
func minDeletionSize(strs []string) int {
	// 对于每一行，j 列的字母都 <= i 列的字母？
	lessEq := func(j, i int) bool {
		for _, s := range strs {
			if s[j] > s[i] {
				return false
			}
		}
		return true
	}

	m := len(strs[0])
	f := make([]int, m)
	for i := range m {
		for j := range i {
			// 如果 f[j] <= f[i]，就不用跑 O(n) 的 lessEq 了
			if f[j] > f[i] && lessEq(j, i) {
				f[i] = f[j]
			}
		}
		f[i]++
	}
	return m - slices.Max(f)
}
