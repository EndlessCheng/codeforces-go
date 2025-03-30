package main

import "slices"

// https://space.bilibili.com/206214
func calc(s, t string) int {
	n, m := len(s), len(t)
	mx := make([]int, n+1)
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, m+1)
	}
	for i, x := range s {
		for j, y := range t {
			if x == y {
				f[i+1][j] = f[i][j+1] + 1
			}
		}
		mx[i+1] = slices.Max(f[i+1])
	}
	ans := slices.Max(mx) * 2 // |x| = |y| 的情况

	// 计算 |x| > |y| 的情况，中心扩展法
	for i := range 2*n - 1 {
		l, r := i/2, (i+1)/2
		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
		}
		if l+1 <= r-1 { // s[l+1] 到 s[r-1] 是非空回文串
			ans = max(ans, r-l-1+mx[l+1]*2)
		}
	}
	return ans
}

func longestPalindrome(s, t string) int {
	return max(calc(s, t), calc(reverse(t), reverse(s)))
}

func reverse(s string) string {
	t := []byte(s)
	slices.Reverse(t)
	return string(t)
}
