package main

import (
	"bufio"
	. "fmt"
	"io"
	)

// github.com/EndlessCheng/codeforces-go
func Sol1096D(reader io.Reader, writer io.Writer) {
	min := func(a, b int64) int64 {
		if a <= b {
			return a
		}
		return b
	}
	mins := func(vals ...int64) int64 {
		ans := vals[0]
		for _, val := range vals[1:] {
			if val < ans {
				ans = val
			}
		}
		return ans
	}

	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var n int
	var s string
	Fscan(in, &n, &s)
	costs := make([]int, n)
	for i := range costs {
		Fscan(in, &costs[i])
	}

	dp := [2][4]int64{}
	for i := 1; i <= n; i++ {
		cur, prev := i&1, (i-1)&1
		for j := range dp[cur] {
			dp[cur][j] = dp[prev][j]
		}
		cost := int64(costs[i-1])
		switch s[i-1] {
		case 'h':
			dp[cur][0] += cost
			dp[cur][1] = min(dp[cur][1], dp[prev][0])
		case 'a':
			dp[cur][1] += cost
			dp[cur][2] = min(dp[cur][2], dp[prev][1])
		case 'r':
			dp[cur][2] += cost
			dp[cur][3] = min(dp[cur][3], dp[prev][2])
		case 'd':
			dp[cur][3] += cost
		}
	}
	Fprint(out, mins(dp[n&1][:]...))
}

//func main() {
//	Sol1096D(os.Stdin, os.Stdout)
//}
