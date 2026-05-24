package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
const mod62 = 998244353

func pow62(x, n int) (res int) {
	res = 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod62
		}
		x = x * x % mod62
	}
	return
}

type comb62 struct{ _f, _invF []int }

func newComb62(mx int) *comb62 {
	c := &comb62{[]int{1}, []int{1}}
	c._grow(mx)
	return c
}

func (c *comb62) _grow(mx int) {
	n := len(c._f)
	c._f = slices.Grow(c._f, mx+1)[:mx+1]
	for i := n; i <= mx; i++ {
		c._f[i] = c._f[i-1] * i % mod62
	}
	c._invF = slices.Grow(c._invF, mx+1)[:mx+1]
	c._invF[mx] = pow62(c._f[mx], mod62-2)
	for i := mx; i > n; i-- {
		c._invF[i-1] = c._invF[i] * i % mod62
	}
}

func (c *comb62) f(n int) int {
	if n >= len(c._f) {
		c._grow(n * 2)
	}
	return c._f[n]
}

func (c *comb62) invF(n int) int {
	if n >= len(c._f) {
		c._grow(n * 2)
	}
	return c._invF[n]
}

func (c *comb62) c(n, k int) int {
	if k < 0 || k > n {
		return 0
	}
	return c.f(n) * c.invF(k) % mod62 * c.invF(n-k) % mod62
}

func cf1762E(in io.Reader, out io.Writer) {
	cm := newComb62(0)
	var n int
	Fscan(in, &n)
	ans := 0
	for i := 1; i < n; i++ {
		res := pow62(i, i-1) * pow62(n-i, n-i-1) % mod62 * cm.c(n-2, i-1) % mod62
		ans += (1 - i%2*2) * res
	}
	Fprint(out, (ans%mod62+mod62)%mod62)
}

//func main() { cf1762E(bufio.NewReader(os.Stdin), os.Stdout) }
