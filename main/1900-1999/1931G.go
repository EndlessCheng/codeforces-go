package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
const mod31 = 998244353

func pow31(x, n int) (res int) {
	x %= mod31
	res = 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod31
		}
		x = x * x % mod31
	}
	return
}

type comb31 struct{ _f, _invF []int }

func newComb31(mx int) *comb31 {
	c := &comb31{[]int{1}, []int{1}}
	c._grow(mx)
	return c
}

func (c *comb31) _grow(mx int) {
	n := len(c._f)
	c._f = append(make([]int, 0, mx+1), c._f...)[:mx+1]
	for i := n; i <= mx; i++ {
		c._f[i] = c._f[i-1] * i % mod31
	}
	c._invF = append(make([]int, 0, mx+1), c._invF...)[:mx+1]
	c._invF[mx] = pow31(c._f[mx], mod31-2)
	for i := mx; i > n; i-- {
		c._invF[i-1] = c._invF[i] * i % mod31
	}
}

func (c *comb31) f(n int) int {
	if n >= len(c._f) {
		c._grow(n * 2)
	}
	return c._f[n]
}

func (c *comb31) invF(n int) int {
	if n >= len(c._f) {
		c._grow(n * 2)
	}
	return c._invF[n]
}

func (c *comb31) c(n, k int) int {
	if k < 0 || k > n {
		return 0
	}
	return c.f(n) * c.invF(k) % mod31 * c.invF(n-k) % mod31
}

// 盒子有区分，球无区分，允许空盒
func (c *comb31) h(box, ball int) int {
	return c.c(box+ball-1, ball)
}

func cf1931G(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	cm := newComb31(0)

	var T, a, b, c, d int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &a, &b, &c, &d)
		if a > b {
			a, b = b, a
		}
		if b-a > 1 {
			Fprintln(out, 0)
		} else if b == 0 {
			if c > 0 && d > 0 {
				Fprintln(out, 0)
			} else {
				Fprintln(out, 1)
			}
		} else if a == b {
			Fprintln(out, (cm.h(b, c)*cm.h(b+1, d)+cm.h(b+1, c)*cm.h(b, d))%mod31)
		} else {
			Fprintln(out, cm.h(b, c)*cm.h(b, d)%mod31)
		}
	}
}

//func main() { cf1931G(os.Stdin, os.Stdout) }
