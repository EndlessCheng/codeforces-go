package main

import (
	. "fmt"
	"io"
	"slices"
	"strconv"
	"strings"
)

// https://github.com/EndlessCheng
const mod = 1_000_000_007

func pow(x, n int) (res int) {
	res = 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return
}

type comb struct{ _f, _invF []int }

func (c *comb) _grow(mx int) {
	n := len(c._f)
	c._f = slices.Grow(c._f, mx+1)[:mx+1]
	for i := n; i <= mx; i++ {
		c._f[i] = c._f[i-1] * i % mod
	}
	c._invF = slices.Grow(c._invF, mx+1)[:mx+1]
	c._invF[mx] = pow(c._f[mx], mod-2)
	for i := mx; i > n; i-- {
		c._invF[i-1] = c._invF[i] * i % mod
	}
}

func (c *comb) f(n int) int {
	if n >= len(c._f) {
		c._grow(n * 2)
	}
	return c._f[n]
}

func (c *comb) invF(n int) int {
	if n >= len(c._f) {
		c._grow(n * 2)
	}
	return c._invF[n]
}

func (c *comb) c(n, k int) int {
	if k < 0 || k > n {
		return 0
	}
	return c.f(n) * c.invF(k) % mod * c.invF(n-k) % mod
}

func cf145C(in io.Reader, out io.Writer) {
	var n, k, tot int
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

	cm := &comb{[]int{1}, []int{1}}
	ans := cm.c(n-tot, k)
	f := make([]int, min(tot, k)+1)
	f[0] = 1
	for _, c := range cnt {
		for j := len(f) - 1; j > 0; j-- {
			ans += f[j-1] * c % mod * cm.c(n-tot, k-j) % mod
			f[j] = (f[j] + f[j-1]*c) % mod
		}
	}
	Fprint(out, ans%mod)
}

//func main() { cf145C(bufio.NewReader(os.Stdin), os.Stdout) }
