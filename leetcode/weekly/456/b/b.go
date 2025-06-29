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
