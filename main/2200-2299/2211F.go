package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2211F(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mod = 676767677
	add := func(a *int, b int) {
		*a += b
		if *a >= mod {
			*a -= mod
		}
	}

	var Q int
	Fscan(in, &Q)
	type pair struct{ n, m int }
	qs := make([]pair, Q)
	maxNM := 1
	for i := range qs {
		Fscan(in, &qs[i].n, &qs[i].m)
		qs[i].m--
		maxNM = max(maxNM, qs[i].n+qs[i].m)
	}

	mul := make([]int, maxNM+1)
	inv := make([]int, maxNM+1)
	iv := make([]int, maxNM+1)
	mul[0], inv[0], iv[0] = 1, 1, 1
	if maxNM >= 1 {
		mul[1], inv[1], iv[1] = 1, 1, 1
	}
	for i := 2; i <= maxNM; i++ {
		mul[i] = mul[i-1] * i % mod
		iv[i] = (mod - mod/i) * iv[mod%i] % mod
		inv[i] = inv[i-1] * iv[i] % mod
	}

	D := func(n, m int) int {
		return mul[n+m] * inv[n] % mod * inv[m] % mod
	}

	for _, q := range qs {
		n, m := q.n, q.m
		ans := 0

		var f func(int, int, int)
		f = func(l, r, dep int) {
			if r-l <= 1 {
				return
			}
			mid := (l + r) >> 1
			dl := mid - l
			dr := r - mid
			v := D(n, m)
			if l != 0 {
				add(&v, mod-D(n-dl, m))
			}
			if r != n+1 {
				add(&v, mod-D(n-dr, m))
			}
			if l != 0 && r != n+1 {
				add(&v, D(n-dl-dr, m))
			}
			add(&ans, v*dep%mod)
			f(l, mid, dep+1)
			f(mid, r, dep+1)
		}

		f(0, n+1, 1)
		Fprintln(out, ans)
	}
}

//func main() { cf2211F(bufio.NewReader(os.Stdin), os.Stdout) }
