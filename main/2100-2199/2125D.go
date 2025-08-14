package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2125D(in io.Reader, out io.Writer) {
	const mod = 998244353
	pow := func(x, n int) int {
		res := 1
		for ; n > 0; n /= 2 {
			if n%2 > 0 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}
	var n, m, l, r, p, q int
	Fscan(in, &n, &m)
	type pair struct{ l, p int }
	g := make([][]pair, m+1)
	notP := 1
	for range n {
		Fscan(in, &l, &r, &p, &q)
		p = p * pow(q, mod-2) % mod
		g[r] = append(g[r], pair{l, p})
		notP = notP * (1 - p) % mod
	}

	f := make([]int, m+1)
	f[0] = notP
	for i, ps := range g {
		for _, p := range ps {
			// [p.l, i] 只能有这一条线段，所以其余右端点在 [p.l, i] 中的区间都不能留
			f[i] = (f[i] + f[p.l-1]*p.p%mod*pow(1-p.p, mod-2)) % mod
		}
	}
	Fprint(out, (f[m]+mod)%mod)
}

//func main() { cf2125D(bufio.NewReader(os.Stdin), os.Stdout) }
