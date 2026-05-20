package main

import (
	. "fmt"
	"io"
	"os"
	"slices"
)

// https://github.com/EndlessCheng
const mod = 998244353

func pow(x, n int) (res int) {
	res = 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return
}

type comb struct{ _f, _invF []int }

func newComb(mx int) *comb {
	c := &comb{[]int{1}, []int{1}}
	c._grow(mx)
	return c
}

func (c *comb) _grow(mx int) {
	n := len(c._f)
	c._f = slices.Grow(c._f, mx+1)[:mx+1]
	for i := n; i <= mx; i++ {
		c._f[i] = c._f[i-1] * i % mod
	}
	c._invF = slices.Grow(c._invF, mx+1)[:mx+1]
	c._invF[mx] = pow(c._f[mx], mod-2)
	for i := mx; i > n; i-- {
		c._invF[i-1] = c._invF[i] * i % mod
	}
}

func (c *comb) f(n int) int {
	if n >= len(c._f) {
		c._grow(n * 2)
	}
	return c._f[n]
}

func (c *comb) invF(n int) int {
	if n >= len(c._f) {
		c._grow(n * 2)
	}
	return c._invF[n]
}

func (c *comb) c(n, k int) int {
	if k < 0 || k > n {
		return 0
	}
	return c.f(n) * c.invF(k) % mod * c.invF(n-k) % mod
}

var cm = newComb(0)

func run(in io.Reader, out io.Writer) {
	var mx, c int
	Fscan(in, &mx)
	f := make([]int, mx+1)
	f[0] = 1
	for range 26 {
		Fscan(in, &c)
		for j := mx; j > 0; j-- {
			for k := 1; k <= min(j, c); k++ {
				f[j] = (f[j] + f[j-k]*cm.c(j, k)) % mod
			}
		}
	}

	ans := 0
	for _, v := range f[1:] {
		ans += v
	}
	Fprint(out, ans%mod)
}

func main() { run(os.Stdin, os.Stdout) }
