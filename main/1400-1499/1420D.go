package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func CF1420D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const mod = 998244353
	const mx int = 3e5
	F := [mx + 1]int64{1}
	for i := 1; i <= mx; i++ {
		F[i] = F[i-1] * int64(i) % mod
	}
	pow := func(x, n int64) (res int64) {
		res = 1
		for ; n > 0; n >>= 1 {
			if n&1 > 0 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}
	invF := [...]int64{mx: pow(F[mx], mod-2)}
	for i := mx; i > 0; i-- {
		invF[i-1] = invF[i] * int64(i) % mod
	}
	C := func(n, k int) int64 {
		if k < 0 || k > n {
			return 0
		}
		return F[n] * invF[k] % mod * invF[n-k] % mod
	}

	var n, k int
	Fscan(in, &n, &k)
	a := make([]struct{ l, r int }, n)
	for i := range a {
		Fscan(in, &a[i].l, &a[i].r)
	}

	cntX := map[int]int{}
	d := map[int]int{}
	for _, p := range a {
		cntX[p.l]++
		d[p.l]++
		d[p.r+1]--
	}
	xs := make([]int, 0, len(d))
	for k := range d {
		xs = append(xs, k)
	}
	sort.Ints(xs)
	cnt := make(map[int]int, len(xs))
	s := 0
	for _, v := range xs {
		s += d[v]
		cnt[v] = s
	}

	ans := int64(0)
	for x, c := range cntX {
		ans += C(cnt[x], k) - C(cnt[x]-c, k)
	}
	Fprint(out, (ans%mod+mod)%mod)
}

//func main() { CF1420D(os.Stdin, os.Stdout) }
