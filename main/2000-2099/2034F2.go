package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
const mod34 = 998244353

func pow34(x, n int) int {
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod34
		}
		x = x * x % mod34
	}
	return res
}

type comb34 struct{ _f, _invF []int }

func newComb34(mx int) *comb34 {
	c := &comb34{[]int{1}, []int{1}}
	c._grow(mx)
	return c
}

func (c *comb34) _grow(mx int) {
	n := len(c._f)
	c._f = slices.Grow(c._f, mx+1)[:mx+1]
	for i := n; i <= mx; i++ {
		c._f[i] = c._f[i-1] * i % mod34
	}
	c._invF = slices.Grow(c._invF, mx+1)[:mx+1]
	c._invF[mx] = pow34(c._f[mx], mod34-2)
	for i := mx; i > n; i-- {
		c._invF[i-1] = c._invF[i] * i % mod34
	}
}

func (c *comb34) f(n int) int {
	if n >= len(c._f) {
		c._grow(n * 2)
	}
	return c._f[n]
}

func (c *comb34) invF(n int) int {
	if n >= len(c._f) {
		c._grow(n * 2)
	}
	return c._invF[n]
}

func (c *comb34) c(n, k int) int {
	if k < 0 || k > n {
		return 0
	}
	return c.f(n) * c.invF(k) % mod34 * c.invF(n-k) % mod34
}

func cf2034F2(in io.Reader, out io.Writer) {
	cm := newComb34(0)
	var T, n, m, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m, &k)
		type pair struct{ r, b int }
		a := make([]pair, k+1)
		for i := 1; i <= k; i++ {
			Fscan(in, &a[i].r, &a[i].b)
		}
		slices.SortFunc(a[1:], func(x, y pair) int { return x.r + x.b - y.r - y.b })

		ans := 0
		f := make([]int, k+1)
		f[0] = 1
		for i := 1; i <= k; i++ {
			p := a[i]
			for j, q := range a[:i] {
				x := p.r - q.r
				y := p.b - q.b
				if x >= 0 && y >= 0 {
					f[i] = (f[i] + f[j]*cm.c(x+y, x)) % mod34
				}
			}
			nn := n - p.r
			mm := m - p.b
			ans = (ans + f[i]*(nn*2+mm)%mod34*cm.c(nn+mm, nn)) % mod34
		}
		ans = (ans*pow34(cm.c(n+m, n), mod34-2) + n*2 + m) % mod34
		Fprintln(out, ans)
	}
}

//func main() { cf2034F2(bufio.NewReader(os.Stdin), os.Stdout) }
