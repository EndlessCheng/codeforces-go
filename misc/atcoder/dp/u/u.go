package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	a := make([][]int, n)
	for i := range a {
		a[i] = make([]int, n)
		for j := range a[i] {
			Fscan(in, &a[i][j])
		}
	}
	m := 1 << n
	sum := make([]int, m)
	for i := range sum {
		for s := uint(i); s > 0; s &= s - 1 {
			p := bits.TrailingZeros(s)
			for t := s & (s - 1); t > 0; t &= t - 1 {
				sum[i] += a[p][bits.TrailingZeros(t)]
			}
		}
	}
	dp := make([]int, m)
	for s, dv := range dp {
		t := m - 1 ^ s
		for sub := t; sub > 0; sub = (sub - 1) & t {
			dp[s|sub] = max(dp[s|sub], dv+sum[sub])
		}
	}
	Fprint(out, dp[m-1])
}

func main() { run(os.Stdin, os.Stdout) }

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
