package main

import (
	"cmp"
	"slices"
)

// https://space.bilibili.com/206214
// 计算 s 和 t 的最长公共前缀（LCP）长度
func calcLCP(s, t string) int {
	n := min(len(s), len(t))
	for i := range n {
		if s[i] != t[i] {
			return i
		}
	}
	return n
}

func longestCommonPrefix(words []string, k int) []int {
	n := len(words)
	if k >= n { // 移除一个字符串后，剩余字符串少于 k 个
		return make([]int, n)
	}

	idx := make([]int, n)
	for i := range idx {
		idx[i] = i
	}
	slices.SortFunc(idx, func(i, j int) int { return cmp.Compare(words[i], words[j]) })

	// 计算最大 LCP 长度和次大 LCP 长度，同时记录最大 LCP 来自哪里
	mx, mx2, mxI := -1, -1, -1
	for i := range n - k + 1 {
		// 排序后，[i, i+k-1] 的 LCP 等于两端点的 LCP
		lcp := calcLCP(words[idx[i]], words[idx[i+k-1]])
		if lcp > mx {
			mx, mx2, mxI = lcp, mx, i
		} else if lcp > mx2 {
			mx2 = lcp
		}
	}

	ans := make([]int, n)
	for i := range ans {
		ans[i] = mx // 先初始化成最大 LCP 长度
	}
	// 移除下标在 [mxI, mxI+k-1] 中的字符串，会导致最大 LCP 变成次大 LCP
	for _, i := range idx[mxI : mxI+k] {
		ans[i] = mx2 // 改成次大 LCP 长度
	}
	return ans
}
