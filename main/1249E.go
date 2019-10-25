package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func Sol1249E(reader io.Reader, writer io.Writer) {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var n, c int
	Fscan(in, &n, &c)
	a := make([]int, n)
	for i := 1; i < n; i++ {
		Fscan(in, &a[i])
	}
	b := make([]int, n)
	for i := 1; i < n; i++ {
		Fscan(in, &b[i])
	}

	dp := make([][2]int, n)
	dp[1][0] = a[1]
	dp[1][1] = c + b[1]
	for i := 2; i < n; i++ {
		dp[i][0] = min(dp[i-1][0], dp[i-1][1]) + a[i]
		dp[i][1] = min(dp[i-1][0]+c, dp[i-1][1]) + b[i]
	}
	for _, d := range dp {
		Fprint(out, min(d[0], d[1]), " ")
	}
}

func main() {
	Sol1249E(os.Stdin, os.Stdout)
}
