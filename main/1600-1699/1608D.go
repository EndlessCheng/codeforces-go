package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
const mod08 = 998244353

func pow08(x, n int) (res int) {
	res = 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod08
		}
		x = x * x % mod08
	}
	return
}

type comb08 struct{ _f, _invF []int }

func newComb08(mx int) *comb08 {
	c := &comb08{[]int{1}, []int{1}}
	c._grow(mx)
	return c
}

func (c *comb08) _grow(mx int) {
	n := len(c._f)
	c._f = slices.Grow(c._f, mx+1)[:mx+1]
	for i := n; i <= mx; i++ {
		c._f[i] = c._f[i-1] * i % mod08
	}
	c._invF = slices.Grow(c._invF, mx+1)[:mx+1]
	c._invF[mx] = pow08(c._f[mx], mod08-2)
	for i := mx; i > n; i-- {
		c._invF[i-1] = c._invF[i] * i % mod08
	}
}

func (c *comb08) f(n int) int {
	if n >= len(c._f) {
		c._grow(n * 2)
	}
	return c._f[n]
}

func (c *comb08) invF(n int) int {
	if n >= len(c._f) {
		c._grow(n * 2)
	}
	return c._invF[n]
}

func (c *comb08) c(n, k int) int {
	if k < 0 || k > n {
		return 0
	}
	return c.f(n) * c.invF(k) % mod08 * c.invF(n-k) % mod08
}

func cf1608D(in io.Reader, out io.Writer) {
	cm := newComb08(0)
	var n, q, w int
	var s string
	bad, allBW, allWB := 1, 1, 1
	Fscan(in, &n)
	for range n {
		Fscan(in, &s)
		u, v := s[0], s[1]
		if u == 'W' {
			w++
		} else if u == '?' {
			q++
		}
		if v == 'W' {
			w++
		} else if v == '?' {
			q++
		}

		// 如果没有 BB 和 WW，那么同时包含 BW 和 WB 的情况非法（减去只包含 BW 和只包含 WB 的情况）
		if u == '?' && v == '?' {
			bad = bad * 2 % mod08 // 可以变成 BW 也可以变成 WB
		} else if u == v {
			// 有 BB 和 WW，可以让同时包含 BW 和 WB 的情况合法
			// c(q, n-w) 保证了在有 BB 的情况下，一定有 WW
			bad = 0
		}
		if u == 'W' || v == 'B' {
			allBW = 0
		}
		if u == 'B' || v == 'W' {
			allWB = 0
		}
	}

	Fprint(out, (cm.c(q, n-w)-bad+allBW+allWB+mod08)%mod08)
}

//func main() { cf1608D(bufio.NewReader(os.Stdin), os.Stdout) }
