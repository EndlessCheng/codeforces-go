package main

import (
	. "fmt"
	"io"
	"os"
	"slices"
)

// https://github.com/EndlessCheng
const mod = 1_000_000_007

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

func cf756D(in io.Reader, out io.Writer) {
	var n, ans int
	var s string
	Fscan(in, &n, &s)
	f := [26][]int{}
	for i := range f {
		f[i] = make([]int, n+1)
	}
	sumF := make([]int, n+1)
	for i, b := range s {
		b -= 'a'
		for sz := i + 1; sz > 0; sz-- {
			old := f[b][sz]
			if sz > 1 {
				f[b][sz] = (sumF[sz-1] - f[b][sz-1]) % mod
			} else {
				f[b][sz] = 1
			}
			sumF[sz] = (sumF[sz] + f[b][sz] - old) % mod
		}
	}

	for sz := 1; sz <= n; sz++ {
		ans = (ans + sumF[sz]*cm.c(n-1, sz-1)) % mod
	}
	Fprint(out, (ans+mod)%mod)
}

func main() { cf756D(os.Stdin, os.Stdout) }
