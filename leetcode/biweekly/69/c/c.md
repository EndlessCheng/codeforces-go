构造 + 贪心 + 分类讨论

先将每个字符串的出现次数统计出来，记在 $\textit{cnt}$ 中。

然后分字符串的两个字母相同和不同两种情况讨论。

设初始回文串为空。对于两个字母不同的情况，用 $\texttt{A}$ 和 $\texttt{B}$ 表示两个不同的小写字母，我们可以选择将 $\min(\textit{cnt}[\texttt{AB}],\textit{cnt}[\texttt{BA}])$ 个 $\texttt{AB}$ 和 $\texttt{BA}$ 字符串，对称添加到当前回文串左右两侧。

对于两个字母相同的情况，与上述类似，我们可以选择偶数个 $\texttt{AA}$ 对称添加到当前回文串左右两侧。如果某个 $\texttt{AA}$ 出现了奇数次，我们还可以将其添加到当前回文串的正中。

代码实现时，我们可以用一个 $26 \times 26$ 的二维数组来统计各个字符串的出现次数，从而优化代码运行时间。

```go
func longestPalindrome(words []string) (ans int) {
	cnt := [26][26]int{}
	for _, s := range words {
		cnt[s[0]-'a'][s[1]-'a']++
	}
	odd := 0 // 是否有一个出现奇数次的 AA
	for i := 0; i < 26; i++ {
		c := cnt[i][i] // 相同字符
		ans += c &^ 1  // c &^ 1 等价于 c - c % 2
		odd |= c & 1
		for j := i + 1; j < 26; j++ {
			ans += min(cnt[i][j], cnt[j][i]) * 2 // AB BA 取出现次数最小值
		}
	}
	return (ans + odd) * 2 // 上面的循环统计的是字符串个数，最后再乘 2
}

func min(a, b int) int { if a > b { return b }; return a }
```
