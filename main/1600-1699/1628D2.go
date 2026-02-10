package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
const mod28 = 1_000_000_007

func pow28(x, n int) (res int) {
	res = 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod28
		}
		x = x * x % mod28
	}
	return
}

type comb28 struct{ _f, _invF []int }

func newComb28(mx int) *comb28 {
	c := &comb28{[]int{1}, []int{1}}
	c._grow(mx)
	return c
}

func (c *comb28) _grow(mx int) {
	n := len(c._f)
	c._f = slices.Grow(c._f, mx+1)[:mx+1]
	for i := n; i <= mx; i++ {
		c._f[i] = c._f[i-1] * i % mod28
	}
	c._invF = slices.Grow(c._invF, mx+1)[:mx+1]
	c._invF[mx] = pow28(c._f[mx], mod28-2)
	for i := mx; i > n; i-- {
		c._invF[i-1] = c._invF[i] * i % mod28
	}
}

func (c *comb28) f(n int) int {
	if n >= len(c._f) {
		c._grow(n * 2)
	}
	return c._f[n]
}

func (c *comb28) invF(n int) int {
	if n >= len(c._f) {
		c._grow(n * 2)
	}
	return c._invF[n]
}

func (c *comb28) c(n, k int) int {
	if k < 0 || k > n {
		return 0
	}
	return c.f(n) * c.invF(k) % mod28 * c.invF(n-k) % mod28
}

func cf1628D2(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	cm := newComb28(0)
	var T, n, m, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m, &k)
		if n == m {
			Fprintln(out, m*k%mod28)
			continue
		}
		ans := 0
		for i := 1; i <= m; i++ {
			ans = (ans + cm.c(n-i-1, m-i)*i%mod28*pow28(2, mod28-1+i-n)) % mod28
		}
		Fprintln(out, ans*k%mod28)
	}
}

//func main() { cf1628D2(bufio.NewReader(os.Stdin), os.Stdout) }
