package main

import "math"

// https://space.bilibili.com/206214
func minOperations(s, t string) int {
	var cnt [26][26]int
	var op int
	update := func(x, y byte) {
		if x == y {
			return
		}
		x -= 'a'
		y -= 'a'
		if cnt[y][x] > 0 {
			cnt[y][x]--
		} else {
			cnt[x][y]++
			op++
		}
	}

	n := len(s)
	// 预处理 revOp
	revOp := make([][]int, n)
	for i := range revOp {
		revOp[i] = make([]int, n)
	}
	// 中心扩展法
	// i 为偶数表示奇长度子串，i 为奇数表示偶长度子串
	for i := range 2*n - 1 {
		cnt = [26][26]int{}
		op = 1
		// 从闭区间 [l,r] 开始向左右扩展
		l, r := i/2, (i+1)/2
		for l >= 0 && r < n {
			update(s[l], t[r])
			if l != r {
				update(s[r], t[l])
			}
			revOp[l][r] = op
			l--
			r++
		}
	}

	f := make([]int, n+1)
	for i := range n {
		res := math.MaxInt
		cnt = [26][26]int{}
		op = 0 // 不反转时的最小操作次数
		for j := i; j >= 0; j-- {
			update(s[j], t[j])
			res = min(res, f[j]+min(op, revOp[j][i]))
		}
		f[i+1] = res
	}
	return f[n]
}
