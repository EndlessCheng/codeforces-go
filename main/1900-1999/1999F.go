package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
const mod99 = 1_000_000_007

func pow99(x, n int) (res int) {
	res = 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod99
		}
		x = x * x % mod99
	}
	return
}

type comb99 struct{ _f, _invF []int }

func newComb99(mx int) *comb99 {
	c := &comb99{[]int{1}, []int{1}}
	c._grow(mx)
	return c
}

func (c *comb99) _grow(mx int) {
	n := len(c._f)
	c._f = slices.Grow(c._f, mx+1)[:mx+1]
	for i := n; i <= mx; i++ {
		c._f[i] = c._f[i-1] * i % mod99
	}
	c._invF = slices.Grow(c._invF, mx+1)[:mx+1]
	c._invF[mx] = pow99(c._f[mx], mod99-2)
	for i := mx; i > n; i-- {
		c._invF[i-1] = c._invF[i] * i % mod99
	}
}

func (c *comb99) f(n int) int {
	if n >= len(c._f) {
		c._grow(n * 2)
	}
	return c._f[n]
}

func (c *comb99) invF(n int) int {
	if n >= len(c._f) {
		c._grow(n * 2)
	}
	return c._invF[n]
}

func (c *comb99) c(n, k int) int {
	if k < 0 || k > n {
		return 0
	}
	return c.f(n) * c.invF(k) % mod99 * c.invF(n-k) % mod99
}

func cf1999F(in io.Reader, out io.Writer) {
	cm := newComb99(0)
	var T, n, k, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		c1 := 0
		for range n {
			Fscan(in, &v)
			c1 += v
		}
		ans := 0
		for i := range k/2 + 1 {
			ans = (ans + cm.c(n-c1, i)*cm.c(c1, k-i)) % mod99
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1999F(bufio.NewReader(os.Stdin), os.Stdout) }
