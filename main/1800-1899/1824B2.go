package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
const mod24 = 1_000_000_007

func pow24(x int64, n int) (res int64) {
	x %= mod24
	res = 1
	for ; n > 0; n >>= 1 {
		if n&1 > 0 {
			res = res * x % mod24
		}
		x = x * x % mod24
	}
	return
}

type comb24 struct{ _f, _invF []int64 }

func newComb24(mx int) *comb24 {
	c := &comb24{[]int64{1}, []int64{1}}
	c._init(mx)
	return c
}

func (c *comb24) _init(mx int) {
	n := len(c._f)
	c._f = append(make([]int64, 0, mx+1), c._f...)[:mx+1]
	for i := n; i <= mx; i++ {
		c._f[i] = c._f[i-1] * int64(i) % mod24
	}
	c._invF = append(make([]int64, 0, mx+1), c._invF...)[:mx+1]
	c._invF[mx] = pow24(c._f[mx], mod24-2)
	for i := mx; i > n; i-- {
		c._invF[i-1] = c._invF[i] * int64(i) % mod24
	}
}

func (c *comb24) f(n int) int64 {
	if n >= len(c._f) {
		c._init(n * 2)
	}
	return c._f[n]
}

func (c *comb24) invF(n int) int64 {
	if n >= len(c._f) {
		c._init(n * 2)
	}
	return c._invF[n]
}

func (c *comb24) c(n, k int) int64 {
	if k < 0 || k > n {
		return 0
	}
	return c.f(n) * c.invF(k) % mod24 * c.invF(n-k) % mod24
}

func CF1824B2(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, k, v, w int
	Fscan(in, &n, &k)
	if k%2 > 0 {
		Fprint(out, 1)
		return
	}
	cm := newComb24(0)
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}
	ans := int64(0)
	var f func(int, int) int
	f = func(v, fa int) int {
		sz := 1
		for _, w := range g[v] {
			if w != fa {
				sz += f(w, v)
			}
		}
		ans = (ans + cm.c(sz, k/2)*cm.c(n-sz, k/2)) % mod24
		return sz
	}
	f(0, -1)
	Fprint(out, (ans*pow24(cm.c(n, k), mod24-2)+1)%mod24)
}

//func main() { CF1824B2(os.Stdin, os.Stdout) }
