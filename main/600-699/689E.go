package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF689E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const mod int64 = 1e9 + 7
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

	var n, k, l, r, s int
	Fscan(in, &n, &k)
	mx := int64(n)
	F := make([]int64, mx+1)
	F[0] = 1
	for i := int64(1); i <= mx; i++ {
		F[i] = F[i-1] * i % mod
	}
	invF := make([]int64, mx+1)
	invF[mx] = pow(F[mx], mod-2)
	for i := mx; i > 0; i-- {
		invF[i-1] = invF[i] * i % mod
	}
	C := func(n, k int) int64 { return F[n] * invF[k] % mod * invF[n-k] % mod }

	d := map[int]int{}
	for ; n > 0; n-- {
		Fscan(in, &l, &r)
		d[l]++
		d[r+1]--
	}
	type pair struct{ p, d int }
	pd := make([]pair, 0, len(d))
	for p, d := range d {
		pd = append(pd, pair{p, d})
	}
	sort.Slice(pd, func(i, j int) bool { return pd[i].p < pd[j].p })
	ans := int64(0)
	for i, p := range pd {
		if s >= k {
			ans = (ans + C(s, k)*int64(p.p-pd[i-1].p)) % mod
		}
		s += p.d
	}
	Fprint(out, ans)
}

//func main() { CF689E(os.Stdin, os.Stdout) }
