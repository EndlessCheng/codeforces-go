package main

/* 预处理 + 动态规划

定义 $f[i][j]$ 表示 $\textit{num}$ 的前 $j$ 个字符划分出的最后一个整数的起始位置为 $i$ 时的方案数。我们所求的答案即为 $\sum\limits_{i=0}^{n-1} f[i][n-1]$。

定义 $\textit{lcp}(i,j)$ 表示后缀 $\textit{num}[i:]$ 和后缀 $\textit{num}[j:]$ 的最长公共前缀的长度。

我们可以通过比较倒数第二个划分出的整数和最后一个划分出的整数的大小，来计算状态转移

对于倒数第二个划分出的整数，记其起始位置为 $k$，结束位置为 $i-1$，有：

- 若其长度小于最后一个划分出的整数，则可以将其方案数全部加到 $f[i][j]$ 上，即 $f[i][j] += \sum f[k][i-1]$，这里 $i-k<j-i+1$；
- 若其长度等于最后一个划分出的整数，则需要比较两个整数的大小，这可以通过比较两个整数（子串）最长公共前缀的下一个字符得出；
- 其其长度大于最后一个划分出的整数，由于不满足题目要求，无法转移。

*/

// github.com/EndlessCheng/codeforces-go
const mod int = 1e9 + 7

func numberOfCombinations(s string) (ans int) {
	if s[0] == '0' {
		return
	}

	n := len(s)
	// 计算 lcp
	lcp := make([][]int, n+1)
	for i := range lcp {
		lcp[i] = make([]int, n+1)
	}
	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if s[i] == s[j] {
				lcp[i][j] = lcp[i+1][j+1] + 1
			}
		}
	}
	// 返回 s[l1:l2] <= s[l2:r2]
	lessEq := func(l1, l2, r2 int) bool {
		l := lcp[l1][l2]
		return l >= r2-l2 || s[l1+l] < s[l2+l]
	}

	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, n)
	}
	for j := 0; j < n; j++ {
		f[0][j] = 1
	}
	for i := 1; i < n; i++ {
		if s[i] == '0' {
			continue
		}
		// k 和 j 同时向左向右扩展
		for j, k, sum := i, i-1, 0; j < n; j++ {
			f[i][j] = sum // 对应上面所说的长度小于最后一个划分出的整数
			if k < 0 {
				continue
			}
			if s[k] > '0' && lessEq(k, i, j+1) {
				f[i][j] = (f[i][j] + f[k][i-1]) % mod // 对应上面所说的长度等于最后一个划分出的整数
			}
			sum = (sum + f[k][i-1]) % mod
			k--
		}
	}
	for _, row := range f {
		ans = (ans + row[n-1]) % mod
	}
	return
}
