package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1481E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	type pair struct{ l, r int }

	var n int
	Fscan(in, &n)
	a := make([]int, n)
	pos := make([]pair, n+1)
	for i := range pos {
		pos[i].l = -1
	}
	for i := range a {
		Fscan(in, &a[i])
		v := a[i]
		if pos[v].l < 0 {
			pos[v].l = i
		}
		pos[v].r = i
	}
	dp := make([]int, n+1)
	c := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		v := a[i]
		c[v]++
		dp[i] = dp[i+1]
		if p := pos[v]; i > p.l {
			dp[i] = max(dp[i], c[v])
		} else {
			dp[i] = max(dp[i], dp[p.r+1]+c[v])
		}
	}
	Fprint(out, n-dp[0])
}

//func main() { CF1481E(os.Stdin, os.Stdout) }
