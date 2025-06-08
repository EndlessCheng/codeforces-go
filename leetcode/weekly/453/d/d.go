package main

import "math"

// https://space.bilibili.com/206214
func minOperations(s, t string) int {
	n := len(s)
	// 预处理 revOp
	revOp := make([][]int, n)
	for i := range revOp {
		revOp[i] = make([]int, n)
	}
	// i 为偶数表示奇回文串，i 为奇数表示偶回文串
	for i := range 2*n - 1 {
		cnt := [26][26]int{}
		op := 1
		// 从闭区间 [l,r] 开始向左右扩展
		l, r := i/2, (i+1)/2
		for l >= 0 && r < n {
			x, y := s[l]-'a', t[r]-'a'
			if x != y {
				if cnt[y][x] > 0 {
					cnt[y][x]--
				} else {
					cnt[x][y]++
					op++
				}
			}
			revOp[l][r] = op
			l--
			r++
		}
	}

	f := make([]int, n+1)
	for i := range n {
		res := math.MaxInt
		cnt := [26][26]int{}
		op := 0
		for j := i; j >= 0; j-- {
			// 不反转
			x, y := s[j]-'a', t[j]-'a'
			if x != y {
				if cnt[y][x] > 0 {
					cnt[y][x]--
				} else {
					cnt[x][y]++
					op++
				}
			}
			res = min(res, f[j]+min(op, revOp[j][i]))
		}
		f[i+1] = res
	}
	return f[n]
}
