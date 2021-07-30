package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF981E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, c int
	Fscan(in, &n, &m)
	a := make([]struct{ l, r, v int }, n)
	for i := range a {
		Fscan(in, &a[i].l, &a[i].r, &a[i].v)
	}
	sort.Slice(a, func(i, j int) bool { return a[i].r < a[j].r })
	dp := make([]int, n+1)
	for _, p := range a {
		v := p.v
		for j := n; j > v; j-- {
			if dp[j-v] >= p.l && dp[j-v] > dp[j] {
				dp[j] = dp[j-v]
			}
		}
		dp[v] = p.r
	}
	for _, v := range dp {
		if v > 0 {
			c++
		}
	}
	Fprintln(out, c)
	for i, v := range dp {
		if v > 0 {
			Fprint(out, i, " ")
		}
	}
}

//func main() { CF981E(os.Stdin, os.Stdout) }
