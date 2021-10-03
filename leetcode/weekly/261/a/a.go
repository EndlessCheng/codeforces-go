package main

/* 贪心

遍历 $s$，遇到 $\texttt{X}$ 就将其与后面两个字符改为 $\texttt{O}$，操作次数加一，然后跳过后面两个字符。
 */

// github.com/EndlessCheng/codeforces-go
func minimumMoves(s string) (ans int) {
	for i := 0; i < len(s); i++ {
		if s[i] == 'X' {
			ans++
			i += 2 // 跳过两个字符
		}
	}
	return
}
