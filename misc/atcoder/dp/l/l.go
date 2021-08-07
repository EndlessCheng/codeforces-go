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
	var n int
	Fscan(in, &n)
	dp := make([]int, n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		dp[i] = a[i]
	}
	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			dp[j] = max(a[i]-dp[j], a[j]-dp[j-1])
		}
	}
	Fprint(out, dp[n-1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() { run(os.Stdin, os.Stdout) }
