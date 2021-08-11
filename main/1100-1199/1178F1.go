package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1178F1(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const mod = 998244353

	var n int
	Fscan(in, &n, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	pMin := make([][]int, n)
	for i := range pMin {
		pMin[i] = make([]int, n)
		pMin[i][i] = i
		for j, p := i+1, i; j < n; j++ {
			if a[j] < a[p] {
				p = j
			}
			pMin[i][j] = p
		}
	}

	dp := make([][]int64, n)
	for i := range dp {
		dp[i] = make([]int64, n)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	var f func(int, int) int64
	f = func(l, r int) (res int64) {
		if l >= r {
			return 1
		}
		dv := &dp[l][r]
		if *dv != -1 {
			return *dv
		}
		defer func() { *dv = res }()
		p := pMin[l][r]
		var sl, sr int64
		for i := l; i <= p; i++ {
			sl = (sl + f(l, i-1)*f(i, p-1)) % mod
		}
		for i := p; i <= r; i++ {
			sr = (sr + f(p+1, i)*f(i+1, r)) % mod
		}
		return sl * sr % mod
	}
	Fprint(out, f(0, n-1))
}

//func main() { CF1178F1(os.Stdin, os.Stdout) }
