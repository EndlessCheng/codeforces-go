package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, v, s int
	Fscan(in, &n, &s)
	dp := make([]int, s+1)
	for i := range dp {
		dp[i] = -1e9
	}
	ans := n + 1
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		// dp[x] 表示和为 x 的时候，左边界的最大值
		for j := s; j >= v; j-- {
			dp[j] = max(dp[j], dp[j-v])
		}
		dp[v] = i
		ans = min(ans, i-dp[s]+1)
	}
	if ans > n {
		ans = -1
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
