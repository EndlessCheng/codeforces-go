package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1783D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const mod = 998244353
	var n, bias int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		bias += a[i]
	}
	dp := make([][]int, n-1)
	for i := range dp {
		dp[i] = make([]int, bias*2+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	var f func(int, int) int
	f = func(i, s int) int {
		if i == n-1 {
			return 1
		}
		dv := &dp[i][s+bias]
		if *dv != -1 {
			return *dv
		}
		res := f(i+1, a[i+1]+s)
		if s != 0 {
			res = (res + f(i+1, a[i+1]-s)) % mod
		}
		*dv = res
		return res
	}
	Fprint(out, f(1, a[1]))
}

//func main() { CF1783D(os.Stdin, os.Stdout) }
