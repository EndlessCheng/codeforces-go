package main

import "math/bits"

/*
枚举回文子序列的中间字符，并枚举其左右字符（从 $\texttt{a}$ 到 $\texttt{z}$），若该字符在中间字符左右侧均存在，则找到一个回文子序列
我们可以在枚举中间字符的同时，计算字符串的前缀和与后缀和，这样就能判断某个字符在中间字符左右侧均存在
由于题目要求相同的子序列只计数一次，我们可以用一个 $26 \cdot 26$ 的布尔数组记录每个回文子序列是否存在
代码实现时可以将布尔数组的第二维用位运算压掉
*/

// github.com/EndlessCheng/codeforces-go
func countPalindromicSubsequence1(s string) (ans int) {
	n := len(s)
	// 统计每个后缀有哪些字母
	suf := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		suf[i] = suf[i+1] | 1<<(s[i]-'a')
	}

	pre := 0
	has := [26]int{}
	// 枚举回文子序列的中间字母 s[i]
	for i := 1; i < n-1; i++ {
		pre |= 1 << (s[i-1] - 'a') // 统计前缀有哪些字母
		// pre & suf[i+1] O(1) 计算在 s[i] 左右两侧都有的字母，添加到 has[s[i]-'a'] 中
		has[s[i]-'a'] |= pre & suf[i+1]
	}

	// 比如 has[1] = 1101 表示 s 中有回文子序列 aba、cbc 和 dbd
	for _, mask := range has {
		ans += bits.OnesCount(uint(mask))
	}
	return
}

func countPalindromicSubsequence(s string) (ans int) {
	n := len(s)
	sufCnt := [26]int{} // 统计后缀每个字母的个数
	suf := 0
	for _, ch := range s[1:] {
		ch -= 'a'
		sufCnt[ch]++
		suf |= 1 << ch
	}

	pre := 0
	has := [26]int{}
	for i := 1; i < n-1; i++ {
		pre |= 1 << (s[i-1] - 'a')
		ch := s[i] - 'a'
		sufCnt[ch]--
		if sufCnt[ch] == 0 { // 现在，后缀 [i+1,n-1] 不包含字母 ch
			suf ^= 1 << ch // 从 suf 中去掉 ch
		}
		has[ch] |= pre & suf
	}

	for _, mask := range has {
		ans += bits.OnesCount(uint(mask))
	}
	return
}
