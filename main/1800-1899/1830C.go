package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/rand"
	"time"
)

// https://space.bilibili.com/206214
const mod30 = 998244353

func (c *comb30) pow(x int64, n int) (res int64) {
	x %= mod30
	res = 1
	for ; n > 0; n >>= 1 {
		if n&1 > 0 {
			res = res * x % mod30
		}
		x = x * x % mod30
	}
	return
}

type comb30 struct{ _f, _invF []int64 }

func newComb30(mx int) *comb30 {
	c := &comb30{[]int64{1}, []int64{1}}
	c._init(mx)
	return c
}

func (c *comb30) _init(mx int) {
	n := len(c._f)
	c._f = append(make([]int64, 0, mx+1), c._f...)[:mx+1]
	for i := n; i <= mx; i++ {
		c._f[i] = c._f[i-1] * int64(i) % mod30
	}
	c._invF = append(make([]int64, 0, mx+1), c._invF...)[:mx+1]
	c._invF[mx] = c.pow(c._f[mx], mod30-2)
	for i := mx; i > n; i-- {
		c._invF[i-1] = c._invF[i] * int64(i) % mod30
	}
}

func (c *comb30) f(n int) int64 {
	if n >= len(c._f) {
		c._init(n * 2)
	}
	return c._f[n]
}

func (c *comb30) invF(n int) int64 {
	if n >= len(c._f) {
		c._init(n * 2)
	}
	return c._invF[n]
}

func (c *comb30) c(n, k int) int64 {
	if k < 0 || k > n {
		return 0
	}
	return c.f(n) * c.invF(k) % mod30 * c.invF(n-k) % mod30
}

func CF1830C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	cm := newComb30(0)
	rand.Seed(time.Now().UnixNano())

	var T, n, k, l, r int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		diff := make([]uint64, n)
		for ; k > 0; k-- {
			Fscan(in, &l, &r)
			v := rand.Uint64()
			diff[l-1] ^= v
			if r < n {
				diff[r] ^= v
			}
		}
		cnt := map[uint64]int{}
		s := uint64(0)
		for _, d := range diff {
			s ^= d
			cnt[s]++
		}
		ans := int64(1)
		for _, c := range cnt {
			if c%2 > 0 {
				Fprintln(out, 0)
				continue o
			}
			// 卡特兰数 C(c/2)
			ans = ans * (cm.c(c, c/2) - cm.c(c, c/2-1) + mod30) % mod30
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1830C(os.Stdin, os.Stdout) }
