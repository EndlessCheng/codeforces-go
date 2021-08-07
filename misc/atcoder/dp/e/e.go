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
	var n, maxW, totValue int
	Fscan(in, &n, &maxW)
	values := make([]int, n)
	weights := make([]int, n)
	for i := range values {
		Fscan(in, &weights[i], &values[i])
		totValue += values[i]
	}

	dp := make([]int, totValue+1)
	for i := range dp {
		dp[i] = 1e18
	}
	dp[0] = 0
	totValue = 0
	for i, v := range values {
		w := weights[i]
		totValue += v
		for j := totValue; j >= v; j-- {
			dp[j] = min(dp[j], dp[j-v]+w)
		}
	}
	for i := totValue; ; i-- {
		if dp[i] <= maxW {
			Fprint(out, i)
			return
		}
	}
}

func main() { run(os.Stdin, os.Stdout) }

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
