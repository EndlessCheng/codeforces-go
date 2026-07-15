package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
const mod50 = 998244353

func pow50(x, n int) (res int) {
	res = 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod50
		}
		x = x * x % mod50
	}
	return
}

type comb50 struct{ _f, _invF []int }

func newComb50(mx int) *comb50 {
	c := &comb50{[]int{1}, []int{1}}
	c._grow(mx)
	return c
}

func (c *comb50) _grow(mx int) {
	n := len(c._f)
	c._f = slices.Grow(c._f, mx+1)[:mx+1]
	for i := n; i <= mx; i++ {
		c._f[i] = c._f[i-1] * i % mod50
	}
	c._invF = slices.Grow(c._invF, mx+1)[:mx+1]
	c._invF[mx] = pow50(c._f[mx], mod50-2)
	for i := mx; i > n; i-- {
		c._invF[i-1] = c._invF[i] * i % mod50
	}
}

func (c *comb50) f(n int) int {
	if n >= len(c._f) {
		c._grow(n * 2)
	}
	return c._f[n]
}

func (c *comb50) invF(n int) int {
	if n >= len(c._f) {
		c._grow(n * 2)
	}
	return c._invF[n]
}

func (c *comb50) c(n, k int) int {
	if k < 0 || k > n {
		return 0
	}
	return c.f(n) * c.invF(k) % mod50 * c.invF(n-k) % mod50
}

func cf2150D(in io.Reader, out io.Writer) {
	var cm = newComb50(0)
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n+1)
		s := make([]int, n+1)
		for i := 1; i <= n; i++ {
			Fscan(in, &a[i])
			a[i] = (a[i] + a[i-1]) % mod50
			s[i] = (s[i-1] + a[i]) % mod50
		}

		ans := a[n] * n % mod50
		for x := range 2 {
			for y := range 2 {
				for l := 2; l <= n-x-y; l++ {
					if (l+n-x-y)&1 > 0 {
						continue
					}
					f := cm.c((l+n-x-y-2)>>1, l-1)
					g := f * (n - x - y) % mod50 * cm.invF(l) % mod50 * cm.f(l-1) % mod50
					ans = (ans + g*(s[n]-s[l-1]-s[n-l]) + f*(x*a[n-l+1]+y*(a[n]-a[l-1]))) % mod50
				}
			}
		}
		Fprintln(out, (ans+mod50)%mod50)
	}
}

//func main() { cf2150D(bufio.NewReader(os.Stdin), os.Stdout) }
