package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
const mod64 = 998244353

func pow64(x, n int) (res int) {
	res = 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod64
		}
		x = x * x % mod64
	}
	return
}

type comb64 struct{ _f, _invF []int }

func newComb64(mx int) *comb64 {
	c := &comb64{[]int{1}, []int{1}}
	c._grow(mx)
	return c
}

func (c *comb64) _grow(mx int) {
	n := len(c._f)
	c._f = slices.Grow(c._f, mx+1)[:mx+1]
	for i := n; i <= mx; i++ {
		c._f[i] = c._f[i-1] * i % mod64
	}
	c._invF = slices.Grow(c._invF, mx+1)[:mx+1]
	c._invF[mx] = pow64(c._f[mx], mod64-2)
	for i := mx; i > n; i-- {
		c._invF[i-1] = c._invF[i] * i % mod64
	}
}

func (c *comb64) f(n int) int {
	if n >= len(c._f) {
		c._grow(n * 2)
	}
	return c._f[n]
}

func (c *comb64) invF(n int) int {
	if n >= len(c._f) {
		c._grow(n * 2)
	}
	return c._invF[n]
}

func (c *comb64) c(n, k int) int {
	if k < 0 || k > n {
		return 0
	}
	return c.f(n) * c.invF(k) % mod64 * c.invF(n-k) % mod64
}

func cf1264D2(in io.Reader, out io.Writer) {
	cm := newComb64(0)
	var s string
	var pre, preQ, suf, sufQ, ans int
	Fscan(in, &s)

	for _, b := range s {
		if b == ')' {
			suf++
		} else if b == '?' {
			sufQ++
		}
	}

	for _, b := range s[:len(s)-1] {
		if b == ')' {
			suf--
		} else if b == '?' {
			sufQ--
		}
		if b == '(' {
			pre++
		} else if b == '?' {
			preQ++
		}
		ans = (ans + preQ*cm.c(preQ+sufQ-1, pre+preQ-suf) + pre*cm.c(preQ+sufQ, pre+preQ-suf)) % mod64
	}
	Fprint(out, ans)
}

//func main() { cf1264D2(bufio.NewReader(os.Stdin), os.Stdout) }
