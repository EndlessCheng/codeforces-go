package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, maxW, w, v int
	Fscan(in, &n, &maxW)
	dp := make([]int, maxW+1)
	for ; n > 0; n-- {
		Fscan(in, &w, &v)
		for j := maxW; j >= w; j-- {
			dp[j] = max(dp[j], dp[j-w]+v)
		}
	}
	Fprint(out, dp[maxW])
}

func main() { run(os.Stdin, os.Stdout) }

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
