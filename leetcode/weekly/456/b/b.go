package main

// https://space.bilibili.com/206214
func lcp(s, t string) (cnt int) {
	n := min(len(s), len(t))
	for i := 0; i < n && s[i] == t[i]; i++ {
		cnt++
	}
	return
}

func longestCommonPrefix(words []string) []int {
	n := len(words)
	ans := make([]int, n)
	if n == 1 {
		return ans
	}

	mx1, mx2, mx3 := 0, 0, 0
	i1, i2 := -1, -1
	for i := range n - 1 {
		l := lcp(words[i], words[i+1])
		if l > mx1 {
			mx3 = mx2
			mx2, i2 = mx1, i1
			mx1, i1 = l, i
		} else if l > mx2 {
			mx3 = mx2
			mx2, i2 = l, i
		} else if l > mx3 {
			mx3 = l
		}
	}

	for i := range n {
		l := 0
		if 0 < i && i < n-1 {
			l = lcp(words[i-1], words[i+1])
		}
		if i != i1 && i != i1+1 { // 最大 LCP 没被破坏
			ans[i] = max(mx1, l)
		} else if i != i2 && i != i2+1 { // 次大 LCP 没被破坏
			ans[i] = max(mx2, l)
		} else { // 只能是第三大 LCP
			ans[i] = max(mx3, l)
		}
	}
	return ans
}

func longestCommonPrefix1(words []string) []int {
	n := len(words)
	ans := make([]int, n)
	if n == 1 {
		return ans
	}

	// 后缀 [i,n-1] 中的相邻 LCP 长度的最大值
	sufMax := make([]int, n)
	for i := n - 2; i > 0; i-- {
		sufMax[i] = max(sufMax[i+1], lcp(words[i], words[i+1]))
	}

	ans[0] = sufMax[1]
	preMax := 0 // 前缀 [0,i-1] 中的相邻 LCP 长度的最大值
	for i := 1; i < n-1; i++ {
		ans[i] = max(preMax, lcp(words[i-1], words[i+1]), sufMax[i+1])
		preMax = max(preMax, lcp(words[i-1], words[i])) // 为下一轮循环做准备
	}
	ans[n-1] = preMax
	return ans
}
