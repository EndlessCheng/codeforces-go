package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1398D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	var r, g, b int
	Fscan(in, &r, &g, &b)
	x := make(sort.IntSlice, r)
	for i := range x {
		Fscan(in, &x[i])
	}
	sort.Sort(sort.Reverse(x))
	y := make(sort.IntSlice, g)
	for i := range y {
		Fscan(in, &y[i])
	}
	sort.Sort(sort.Reverse(y))
	z := make(sort.IntSlice, b)
	for i := range z {
		Fscan(in, &z[i])
	}
	sort.Sort(sort.Reverse(z))
	dp := make([][][]int, r+1)
	for i := range dp {
		dp[i] = make([][]int, g+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, b+1)
			for k := range dp[i][j] {
				dp[i][j][k] = -1
			}
		}
	}
	var f func(i, j, k int) int
	f = func(i, j, k int) (res int) {
		dv := &dp[i][j][k]
		if *dv != -1 {
			return *dv
		}
		defer func() { *dv = res }()
		if i < r && j < g {
			res = max(res, x[i]*y[j]+f(i+1, j+1, k))
		}
		if i < r && k < b {
			res = max(res, x[i]*z[k]+f(i+1, j, k+1))
		}
		if j < g && k < b {
			res = max(res, y[j]*z[k]+f(i, j+1, k+1))
		}
		return
	}
	Fprint(out, f(0, 0, 0))
}

//func main() { CF1398D(os.Stdin, os.Stdout) }
