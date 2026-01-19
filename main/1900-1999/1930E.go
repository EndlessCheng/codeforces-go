package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
const mod30 = 998244353

func pow30(x, n int) (res int) {
	res = 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod30
		}
		x = x * x % mod30
	}
	return
}

type comb30 struct{ _f, _invF []int }

func newComb30(mx int) *comb30 {
	c := &comb30{[]int{1}, []int{1}}
	c._grow(mx)
	return c
}

func (c *comb30) _grow(mx int) {
	n := len(c._f)
	c._f = slices.Grow(c._f, mx+1)[:mx+1]
	for i := n; i <= mx; i++ {
		c._f[i] = c._f[i-1] * i % mod30
	}
	c._invF = slices.Grow(c._invF, mx+1)[:mx+1]
	c._invF[mx] = pow30(c._f[mx], mod30-2)
	for i := mx; i > n; i-- {
		c._invF[i-1] = c._invF[i] * i % mod30
	}
}

func (c *comb30) f(n int) int {
	if n >= len(c._f) {
		c._grow(n * 2)
	}
	return c._f[n]
}

func (c *comb30) invF(n int) int {
	if n >= len(c._f) {
		c._grow(n * 2)
	}
	return c._invF[n]
}

func (c *comb30) c(n, k int) int {
	if k < 0 || k > n {
		return 0
	}
	return c.f(n) * c.invF(k) % mod30 * c.invF(n-k) % mod30
}

func cf1930E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	cm := newComb30(0)
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		for k := 1; k*2 < n; k++ {
			ans := 1
			for m := 1; m*k*2 < n; m++ {
				ans += cm.c(n, m*k*2) - cm.c(n-(m-1)*k*2-1, k*2-1)
			}
			Fprint(out, (ans%mod30+mod30)%mod30, " ")
		}
		Fprintln(out)
	}
}

//func main() { cf1930E(bufio.NewReader(os.Stdin), os.Stdout) }
