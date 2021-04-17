package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1509C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	min := func(a, b int64) int64 {
		if a < b {
			return a
		}
		return b
	}
	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	sort.Ints(a)
	dp := make([][]int64, n+1)
	for i := range dp {
		dp[i] = make([]int64, n+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	var f func(int, int) int64
	f = func(l, r int) (res int64) {
		if l == r {
			return
		}
		dv := &dp[l][r]
		if *dv >= 0 {
			return *dv
		}
		defer func() { *dv = res }()
		return int64(a[r]-a[l]) + min(f(l+1, r), f(l, r-1))
	}
	Fprint(out, f(0, n-1))
}

//func main() { CF1509C(os.Stdin, os.Stdout) }
