package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF41D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, K int
	Fscan(in, &n, &m, &K)
	K++
	a := make([][]byte, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	dp := make([][][]int, n)
	for i := range dp {
		dp[i] = make([][]int, m)
		for j := range dp[i] {
			dp[i][j] = make([]int, K)
			for k := range dp[i][j] {
				dp[i][j][k] = -1
			}
		}
	}
	type pair struct{ j, k int }
	to := make([][100][11]pair, n)
	var f func(i, j, k int) int
	f = func(i, j, k int) (res int) {
		v := int(a[i][j] & 15)
		kk := (k + v) % K
		if i == 0 {
			if kk > 0 {
				return -1e9
			}
			return v
		}
		dv := &dp[i][j][k]
		if *dv != -1 {
			return *dv
		}
		defer func() { *dv = res }()
		res = -1e9
		if j > 0 {
			if r := f(i-1, j-1, kk); r > res {
				res = r
				to[i][j][k] = pair{-1, kk}
			}
		}
		if j < m-1 {
			if r := f(i-1, j+1, kk); r > res {
				res = r
				to[i][j][k] = pair{1, kk}
			}
		}
		return res + v
	}
	ans, mxJ, mxTo := -1, 0, [][100][11]pair{}
	for j := 0; j < m; j++ {
		res := f(n-1, j, 0)
		if res > ans {
			ans, mxJ, mxTo = res, j, append([][100][11]pair(nil), to...)
		}
	}
	if ans < 0 {
		Fprint(out, -1)
		return
	}
	Fprintln(out, ans)
	Fprintln(out, mxJ+1)
	path := make([]byte, 0, n-1)
	for i, j, k := n-1, mxJ, 0; i > 0; i-- {
		t := mxTo[i][j][k]
		if t.j < 0 {
			path = append(path, 'L')
			j--
		} else {
			path = append(path, 'R')
			j++
		}
		k = t.k
	}
	Fprintf(out, "%s", path)
}

//func main() { CF41D(os.Stdin, os.Stdout) }
