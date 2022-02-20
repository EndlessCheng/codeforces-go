package main

import (
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(in io.Reader, out io.Writer) {
	var n, maxW int
	Fscan(in, &n, &maxW)
	dp := make([]bool, maxW+1)
	dp[0] = true
	g := [2]int{}
	for ; n > 0; n-- {
		Fscan(in, &g[0], &g[1])
	next:
		for j := maxW; j >= 0; j-- {
			for _, w := range g {
				if w <= j && dp[j-w] {
					dp[j] = true
					continue next
				}
			}
			dp[j] = false
		}
	}
	if dp[maxW] {
		Fprint(out, "Yes")
	} else {
		Fprint(out, "No")
	}
}

func main() { run(os.Stdin, os.Stdout) }
