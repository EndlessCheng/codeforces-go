package main

/* 从前往后匹配+从后往前匹配

如果去掉第一种操作，由于交替字符串只有 $\texttt{1}$ 开头与 $\texttt{0}$ 开头两种，我们只需要检查 $s$ 与 $\texttt{1010}\cdots$ 和 $\texttt{0101}\cdots$ 中哪个不同字母最少，这样第二种操作次数最小。

加上第一种操作，相当于需要检查 $s$ 与 $\texttt{1010}\cdots$ 和 $\texttt{0101}\cdots$ 的循环同构串中哪个不同字母最少。

若 $s$ 长度为偶数，则循环同构串只有 $\texttt{1010}\cdots$ 和 $\texttt{0101}\cdots$ 两种情况。

若 $s$ 长度为奇数，则循环同构串中间会出现 $\texttt{11}$ 或 $\texttt{00}$。例如长度为 $7$ 的一种循环同构串为 $\texttt{1011010}\cdots$。

由于 $\texttt{11}$ 或 $\texttt{00}$ 至多会出现一次，我们可以枚举其出现位置，对该位置左边按照 $\texttt{1}$ 开头的交替字符串来匹配，右边按照 $\texttt{0}$ 结尾的交替字符串来匹配；或者左边按照 $\texttt{0}$ 开头的交替字符串来匹配，右边按照 $\texttt{1}$ 结尾的交替字符串来匹配。

进一步地，注意到无论 $s$ 长度为奇数还是偶数，我们总是可以将匹配目标分成以下几类：

- 以 $\texttt{1}$ 开头
- 以 $\texttt{0}$ 开头
- 以 $\texttt{1}$ 开头且以 $\texttt{0}$ 结尾
- 以 $\texttt{0}$ 开头且以 $\texttt{1}$ 结尾

因此可以使用同一枚举方式来求出最小不同字母个数。

时间复杂度：$O(n)$，$n$ 为 $s$ 的长度。
空间复杂度：$O(n)$。

*/

// github.com/EndlessCheng/codeforces-go
func minFlips(s string) int {
	n := len(s)
	ans := n
	// 枚举开头是 0 还是 1
	for head := byte('0'); head <= '1'; head++ {
		// 左边每个位置的不同字母个数
		leftDiff := make([]int, n)
		diff := 0
		for i := range s {
			if s[i] != head^byte(i&1) {
				diff++
			}
			leftDiff[i] = diff
		}

		// 右边每个位置的不同字母个数
		tail := head ^ 1
		diff = 0
		for i := n - 1; i >= 0; i-- {
			// 左边+右边即为整个字符串的不同字母个数，取最小值
			ans = min(ans, leftDiff[i]+diff)
			if s[i] != tail^byte((n-1-i)&1) {
				diff++
			}
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
