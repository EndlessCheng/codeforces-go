package main

import (
	. "fmt"
	"io"
	"slices"
	"strconv"
	"strings"
)

// https://github.com/EndlessCheng
const mod45 = 1_000_000_007

type comb45 struct{ _f, _invF []int }

func (c *comb45) _grow(mx int) {
	pow := func(x, n int) int {
		res := 1
		for ; n > 0; n /= 2 {
			if n%2 > 0 {
				res = res * x % mod45
			}
			x = x * x % mod45
		}
		return res
	}
	n := len(c._f)
	c._f = slices.Grow(c._f, mx+1)[:mx+1]
	for i := n; i <= mx; i++ {
		c._f[i] = c._f[i-1] * i % mod45
	}
	c._invF = slices.Grow(c._invF, mx+1)[:mx+1]
	c._invF[mx] = pow(c._f[mx], mod45-2)
	for i := mx; i > n; i-- {
		c._invF[i-1] = c._invF[i] * i % mod45
	}
}

func (c *comb45) f(n int) int {
	if n >= len(c._f) {
		c._grow(n * 2)
	}
	return c._f[n]
}

func (c *comb45) invF(n int) int {
	if n >= len(c._f) {
		c._grow(n * 2)
	}
	return c._invF[n]
}

func (c *comb45) c(n, k int) int {
	if k < 0 || k > n {
		return 0
	}
	return c.f(n) * c.invF(k) % mod45 * c.invF(n-k) % mod45
}

func cf145C(in io.Reader, out io.Writer) {
	var n, k, tot, ans int
	var s string
	Fscan(in, &n, &k)
	cnt := map[int]int{}
	for range n {
		Fscan(in, &s)
		if strings.Count(s, "4")+strings.Count(s, "7") == len(s) {
			v, _ := strconv.Atoi(s)
			cnt[v]++
			tot++
		}
	}

	f := make([]int, min(tot, k)+1)
	f[0] = 1
	for _, c := range cnt {
		for j := len(f) - 1; j > 0; j-- {
			f[j] = (f[j] + f[j-1]*c) % mod45
		}
	}

	cm := &comb45{[]int{1}, []int{1}}
	for i, v := range f {
		ans = (ans + v*cm.c(n-tot, k-i)) % mod45
	}
	Fprint(out, ans)
}

//func main() { cf145C(bufio.NewReader(os.Stdin), os.Stdout) }
