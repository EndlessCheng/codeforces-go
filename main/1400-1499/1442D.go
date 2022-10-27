package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1442D(_r io.Reader, out io.Writer) {
	buf := make([]byte, 1<<12)
	_i := len(buf)
	rc := func() byte {
		if _i >= len(buf) {
			_r.Read(buf)
			_i = 0
		}
		b := buf[_i]
		_i++
		return b
	}
	r := func() (x int) {
		b := rc()
		for ; '0' > b; b = rc() {
		}
		for ; '0' <= b; b = rc() {
			x = x*10 + int(b&15)
		}
		return
	}
	max := func(a, b int64) int64 {
		if b > a {
			return b
		}
		return a
	}

	n, k := r(), r()
	a := make([][]int, n)
	tot := make([]int64, n)
	for i := range a {
		a[i] = make([]int, r())
		for j := range a[i] {
			a[i][j] = r()
			tot[i] += int64(a[i][j])
		}
	}
	ans := int64(0)
	dp := make([]int64, k+1)
	var f func([][]int, []int64)
	f = func(a [][]int, tot []int64) {
		if len(a) == 1 {
			s := int64(0)
			for i, v := range a[0] {
				if i >= k {
					break
				}
				s += int64(v)
				ans = max(ans, dp[k-(i+1)]+s)
			}
			return
		}
		tmp := append([]int64{}, dp...)
		m := len(a) / 2
		for i, r := range a[:m] {
			for j := k; j >= len(r); j-- {
				dp[j] = max(dp[j], dp[j-len(r)]+tot[i])
			}
		}
		f(a[m:], tot[m:])
		dp = tmp
		for i, r := range a[m:] {
			for j := k; j >= len(r); j-- {
				dp[j] = max(dp[j], dp[j-len(r)]+tot[m+i])
			}
		}
		f(a[:m], tot[:m])
	}
	f(a, tot)
	Fprintln(out, ans)
}

//func main() { CF1442D(os.Stdin, os.Stdout) }
