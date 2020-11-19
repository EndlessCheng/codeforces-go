package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	b := make([]int, n)
	for i := range b {
		Fscan(in, &b[i])
	}
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i, v := range a {
		val := 0
		for j, w := range b {
			if v == w {
				dp[i+1][j] = val + 1
			} else {
				dp[i+1][j] = dp[i][j]
			}
			if w < v {
				val = max(val, dp[i][j])
			}
		}
	}
	ans := 0
	for _, v := range dp[n] {
		ans = max(ans, v)
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
