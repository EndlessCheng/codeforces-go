package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
const mod4071 = 1_000_000_007

func pow4071(x int, n int) (res int) {
	x %= mod4071
	res = 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod4071
		}
		x = x * x % mod4071
	}
	return
}

type comb4071 struct{ _f, _invF []int }

func newComb4071(mx int) *comb4071 {
	c := &comb4071{[]int{1}, []int{1}}
	c._grow(mx)
	return c
}

func (c *comb4071) _grow(mx int) {
	n := len(c._f)
	c._f = append(make([]int, 0, mx+1), c._f...)[:mx+1]
	for i := n; i <= mx; i++ {
		c._f[i] = c._f[i-1] * i % mod4071
	}
	c._invF = append(make([]int, 0, mx+1), c._invF...)[:mx+1]
	c._invF[mx] = pow4071(c._f[mx], mod4071-2)
	for i := mx; i > n; i-- {
		c._invF[i-1] = c._invF[i] * i % mod4071
	}
}

func (c *comb4071) f(n int) int {
	if n >= len(c._f) {
		c._grow(n * 2)
	}
	return c._f[n]
}

func (c *comb4071) invF(n int) int {
	if n >= len(c._f) {
		c._grow(n * 2)
	}
	return c._invF[n]
}

func (c *comb4071) c(n, k int) int {
	if k < 0 || k > n {
		return 0
	}
	return c.f(n) * c.invF(k) % mod4071 * c.invF(n-k) % mod4071
}

func p4071(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	
	d := [1e6 + 1]int{1, 0}
	for i := 2; i <= 1e6; i++ {
		d[i] = (i - 1) * (d[i-1] + d[i-2]) % mod4071
	}
	cm := newComb4071(0)

	var T, n, m int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		Fprintln(out, cm.c(n, m)*d[n-m]%mod4071)
	}
}

//func main() { p4071(os.Stdin, os.Stdout) }
