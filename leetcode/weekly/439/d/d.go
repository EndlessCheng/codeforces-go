package main

import (
	"bytes"
)

// https://space.bilibili.com/206214
func calcZ(s string) []int {
	n := len(s)
	z := make([]int, n)
	boxL, boxR := 0, 0 // z-box 左右边界（闭区间）
	for i := 1; i < n; i++ {
		if i <= boxR {
			z[i] = min(z[i-boxL], boxR-i+1)
		}
		for i+z[i] < n && s[z[i]] == s[i+z[i]] {
			boxL, boxR = i, i+z[i]
			z[i]++
		}
	}
	z[0] = n
	return z
}

func generateString(s, t string) string {
	z := calcZ(t)
	n, m := len(s), len(t)
	ans := bytes.Repeat([]byte{'?'}, n+m-1)
	pre := -m
	for i, b := range s {
		if b != 'T' {
			continue
		}
		size := max(pre+m-i, 0)
		// t 的长为 size 的前后缀必须相同
		if size > 0 && z[m-size] < size {
			return ""
		}
		// size 后的内容都是 '?'，填入 t
		copy(ans[i+size:], t[size:])
		pre = i
	}

	// 计算 <= i 的最近待定位置
	preQ := make([]int, len(ans))
	pre = -1
	for i, c := range ans {
		if c == '?' {
			ans[i] = 'a' // 待定位置的初始值为 a
			pre = i
		}
		preQ[i] = pre
	}

	// 找 ans 中的等于 t 的位置，可以用 KMP 或者 Z 函数
	z = calcZ(t + string(ans))
	for i := 0; i < n; i++ {
		if s[i] != 'F' {
			continue
		}
		// 子串必须不等于 t 
		if z[m+i] < m {
			continue
		}
		// 找最后一个待定位置
		j := preQ[i+m-1]
		if j < i { // 没有
			return ""
		}
		ans[j] = 'b'
		i = j // 直接跳到 j
	}

	return string(ans)
}
