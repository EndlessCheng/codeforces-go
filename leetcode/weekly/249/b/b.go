package main

import "math/bits"

/*
枚举回文子序列的中间字符，并枚举其左右字符（从 $\texttt{a}$ 到 $\texttt{z}$），若该字符在中间字符左右侧均存在，则找到一个回文子序列
我们可以在枚举中间字符的同时，计算字符串的前缀和与后缀和，这样就能判断某个字符在中间字符左右侧均存在
由于题目要求相同的子序列只计数一次，我们可以用一个 $26 \cdot 26$ 的布尔数组记录每个回文子序列是否存在
代码实现时可以将布尔数组的第二维用位运算压掉
*/

// github.com/EndlessCheng/codeforces-go
func countPalindromicSubsequence(s string) (ans int) {
	var pre, suf, has [26]int
	for _, b := range s[1:] {
		suf[b-'a']++
	}
	for i := 1; i < len(s)-1; i++ { // 枚举回文子序列的中间字符
		pre[s[i-1]-'a']++
		suf[s[i]-'a']--
		for j := 0; j < 26; j++ { // 枚举中间字符的左右字符
			if pre[j] > 0 && suf[j] > 0 { // 找到了一个回文子序列
				has[s[i]-'a'] |= 1 << j
			}
		}
	}
	for _, mask := range has {
		ans += bits.OnesCount(uint(mask)) // 累加二进制中 1 的个数即为答案
	}
	return
}
