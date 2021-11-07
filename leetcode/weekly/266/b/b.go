package main

import (
	"strings"
)

/* 单独计算每个元音的出现次数

遍历 $\textit{word}$，若 $\textit{word}[i]$ 是元音，我们考察它能出现在多少个子字符串中。

设 $\textit{word}$ 的长度为 $n$。子字符串 $\textit{word}[l..r]$ 若要包含 $\textit{word}[i]$，则必须满足

- $0\le l\le i$
- $i\le r\le n-1$

这样的 $l$ 有 $i+1$ 个，$r$ 有 $n-i$ 个，因此有 $(i+1)(n-i)$ 个子字符串，所以 $\textit{word}[i]$ 在所有子字符串中一共出现了 $(i+1)(n-i)$ 次。

累加所有出现次数即为答案。
*/

// github.com/EndlessCheng/codeforces-go
func countVowels(word string) (ans int64) {
	for i, ch := range word {
		if strings.ContainsRune("aeiou", ch) {
			ans += int64(i+1) * int64(len(word)-i)
		}
	}
	return
}
