package main

import (
	"strings"
)

// https://space.bilibili.com/206214
// 115. 不同的子序列
func numDistinct(s, t string) int {
	n, m := len(s), len(t)
	if n < m {
		return 0
	}

	f := make([]int, m+1)
	f[0] = 1
	for i, x := range s {
		for j := min(i, m-1); j >= max(m-n+i, 0); j-- {
			if byte(x) == t[j] {
				f[j+1] += f[j]
			}
		}
	}
	return f[m]
}

// 计算插入 T 产生的额外 LCT 子序列个数的最大值
func calcInsertT(s string) (res int) {
	cntT := strings.Count(s, "T") // s[i+1] 到 s[n-1] 的 'T' 的个数
	cntL := 0 // s[0] 到 s[i] 的 'L' 的个数
	for _, c := range s {
		if c == 'T' {
			cntT--
		}
		if c == 'L' {
			cntL++
		}
		res = max(res, cntL*cntT)
	}
	return
}

func numOfSubsequences(s string) int64 {
	extra := max(numDistinct(s, "CT"), numDistinct(s, "LC"), calcInsertT(s))
	return int64(numDistinct(s, "LCT") + extra)
}
