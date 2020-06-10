package main

import (
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func CF510D(in io.Reader, out io.Writer) {
	type pair struct{ v, c int }
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}

	var n int
	Fscan(in, &n)
	a := make([]pair, n)
	for i := range a {
		Fscan(in, &a[i].v)
	}
	for i := range a {
		Fscan(in, &a[i].c)
	}

	dp := map[int]int{0: 0}
	for _, p := range a {
		for g, c := range dp {
			if g = gcd(g, p.v); dp[g] == 0 || c+p.c < dp[g] {
				dp[g] = c + p.c
			}
		}
	}
	if dp[1] > 0 {
		Fprint(out, dp[1])
	} else {
		Fprint(out, -1)
	}
}

func main() { CF510D(os.Stdin, os.Stdout) }
