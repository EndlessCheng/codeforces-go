package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func Sol225C(reader io.Reader, writer io.Writer) {
	min := func(a, b int) int {
		if a <= b {
			return a
		}
		return b
	}
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var n, m, x, y int
	Fscan(in, &n, &m, &x, &y)
	cnt := make([]int, m)
	var s string
	for i := 0; i < n; i++ {
		Fscan(in, &s)
		for j, c := range s {
			if c == '.' {
				cnt[j]++
			}
		}
	}

	sum := make([][2]int, m+1)
	for i := 1; i <= m; i++ {
		sum[i][0] = sum[i-1][0] + cnt[i-1]
		sum[i][1] = sum[i-1][1] + n - cnt[i-1]
	}
	dp := make([][2]int, m+1)
	for i := range dp {
		dp[i][0] = 1e8
		dp[i][1] = 1e8
	}
	dp[0][0] = 0
	dp[0][1] = 0
	for i := 1; i <= m; i++ {
		for j := x; j <= y; j++ {
			if i-j < 0 {
				break
			}
			dp[i][0] = min(dp[i][0], dp[i-j][1]+sum[i][0]-sum[i-j][0])
			dp[i][1] = min(dp[i][1], dp[i-j][0]+sum[i][1]-sum[i-j][1])
		}
	}
	Fprint(out, min(dp[m][0], dp[m][1]))
}

func main() {
	Sol225C(os.Stdin, os.Stdout)
}
