package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
const mod94 = 998244353

func pow94(x int64, n int) (res int64) {
	x %= mod94
	res = 1
	for ; n > 0; n >>= 1 {
		if n&1 > 0 {
			res = res * x % mod94
		}
		x = x * x % mod94
	}
	return
}

type comb94 struct{ f, invF []int64 }

func newComb94(mx int) *comb94 {
	c := &comb94{[]int64{1}, []int64{1}}
	c.init(mx)
	return c
}

func (c *comb94) init(mx int) {
	n := len(c.f)
	c.f = append(make([]int64, 0, mx+1), c.f...)[:mx+1]
	for i := n; i <= mx; i++ {
		c.f[i] = c.f[i-1] * int64(i) % mod94
	}
	c.invF = append(make([]int64, 0, mx+1), c.invF...)[:mx+1]
	c.invF[mx] = pow94(c.f[mx], mod94-2)
	for i := mx; i > n; i-- {
		c.invF[i-1] = c.invF[i] * int64(i) % mod94
	}
}

func (c *comb94) factorial(n int) int64 {
	if n >= len(c.f) {
		c.init(n * 2)
	}
	return c.f[n]
}

func (c *comb94) invFactorial(n int) int64 {
	if n >= len(c.f) {
		c.init(n * 2)
	}
	return c.invF[n]
}

func CF1794D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const mx int = 1e6
	np := [mx + 1]bool{true, true}
	for i := 2; i <= mx; i++ {
		if !np[i] {
			for j := i * i; j <= mx; j += i {
				np[j] = true
			}
		}
	}

	var n, v int
	Fscan(in, &n)
	cnt := map[int]int{}
	for i := 0; i < n*2; i++ {
		Fscan(in, &v)
		cnt[v]++
	}

	cm := newComb94(n)
	f := make([]int64, n+1)
	f[0] = cm.factorial(n)
	for v, c := range cnt {
		for j := n; j >= 0; j-- {
			f[j] = f[j] * cm.invFactorial(c) % mod94
			if j > 0 && !np[v] {
				f[j] = (f[j] + f[j-1]*cm.invFactorial(c-1)) % mod94
			}
		}
	}
	Fprint(out, f[n])
}

//func main() { CF1794D(os.Stdin, os.Stdout) }
