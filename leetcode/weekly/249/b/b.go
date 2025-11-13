package main

import (
	"math/bits"
	"strings"
)

/*
枚举回文子序列的中间字符，并枚举其左右字符（从 $\texttt{a}$ 到 $\texttt{z}$），若该字符在中间字符左右侧均存在，则找到一个回文子序列
我们可以在枚举中间字符的同时，计算字符串的前缀和与后缀和，这样就能判断某个字符在中间字符左右侧均存在
由于题目要求相同的子序列只计数一次，我们可以用一个 $26 \cdot 26$ 的布尔数组记录每个回文子序列是否存在
代码实现时可以将布尔数组的第二维用位运算压掉
*/

// github.com/EndlessCheng/codeforces-go
func countPalindromicSubsequence1(s string) (ans int) {
	for alpha := byte('a'); alpha <= 'z'; alpha++ { // 枚举两侧字母
		i := strings.IndexByte(s, alpha)
		if i < 0 {
			continue
		}
		j := strings.LastIndexByte(s, alpha)
		if i+1 >= j { // 长度不足 3
			continue
		}

		has := [26]bool{}
		for _, mid := range s[i+1 : j] { // 枚举中间字母
			if !has[mid-'a'] {
				has[mid-'a'] = true
				ans++
			}
		}
	}
	return
}

func countPalindromicSubsequence2(s string) (ans int) {
	// 统计 s[1:] 每个字母的个数
	sufCnt := [26]int{}
	for _, ch := range s[1:] {
		sufCnt[ch-'a']++
	}

	preHas := [26]bool{}
	has := [26][26]bool{}
	for i := 1; i < len(s)-1; i++ { // 枚举中间字母
		mid := s[i] - 'a'
		sufCnt[mid]--             // 撤销 mid 的计数，suf_cnt 剩下的就是后缀 [i+1,n-1] 每个字母的个数
		preHas[s[i-1]-'a'] = true // 记录前缀 [0,i-1] 有哪些字母
		for alpha := range 26 {
			// 判断 mid 的左右两侧是否都有字母 alpha
			if preHas[alpha] && sufCnt[alpha] > 0 && !has[mid][alpha] {
				has[mid][alpha] = true
				ans++
			}
		}
	}
	return
}

func countPalindromicSubsequence(s string) (ans int) {
	// 统计 [1,n-1] 每个字母的个数
	sufCnt := [26]int{}
	suf := 0
	for _, ch := range s[1:] {
		ch -= 'a'
		sufCnt[ch]++
		suf |= 1 << ch // 把 ch 记录到二进制数 suf 中，表示后缀有 ch
	}

	pre := 0
	has := [26]int{}                // has[mid] = 由 alpha 组成的二进制数
	for i := 1; i < len(s)-1; i++ { // 枚举中间字母 mid
		mid := s[i] - 'a'
		sufCnt[mid]--         // 撤销 mid 的计数，sufCnt 剩下的就是后缀 [i+1,n-1] 每个字母的个数
		if sufCnt[mid] == 0 { // 后缀 [i+1,n-1] 不包含 mid
			suf ^= 1 << mid // 从 suf 中去掉 mid
		}
		pre |= 1 << (s[i-1] - 'a') // 把 s[i-1] 记录到二进制数 pre 中，表示前缀有 s[i-1]
		has[mid] |= pre & suf      // 计算 pre 和 suf 的交集，|= 表示把交集中的字母加到 has[mid] 中
	}

	for _, mask := range has {
		ans += bits.OnesCount(uint(mask)) // mask 中的每个 1 对应着一个 alpha
	}
	return
}

func countPalindromicSubsequenceAnother(s string) (ans int) {
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
