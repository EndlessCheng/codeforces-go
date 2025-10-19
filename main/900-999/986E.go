package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
func cf986E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}

	const MX int = 1e7 + 1
	lpf := [MX]int{}
	for i := 2; i < MX; i++ {
		if lpf[i] == 0 {
			for j := i; j < MX; j += i {
				if lpf[j] == 0 {
					lpf[j] = i
				}
			}
		}
	}

	const mod = 1_000_000_007
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

	var n, q int
	Fscan(in, &n)
	g := make([][]int, n)
	for range n - 1 {
		var v, w int
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}
	pa := make([][17]int, n)
	dep := make([]int, n)
	var build func(int, int)
	build = func(v, p int) {
		pa[v][0] = p
		for _, w := range g[v] {
			if w != p {
				dep[w] = dep[v] + 1
				build(w, v)
			}
		}
	}
	build(0, -1)
	mx := bits.Len(uint(n))
	for i := range mx - 1 {
		for x := range pa {
			p := pa[x][i]
			if p != -1 {
				pa[x][i+1] = pa[p][i]
			} else {
				pa[x][i+1] = -1
			}
		}
	}
	uptoDep := func(v, d int) int {
		for k := uint32(dep[v] - d); k > 0; k &= k - 1 {
			v = pa[v][bits.TrailingZeros32(k)]
		}
		return v
	}
	getLCA := func(v, w int) int {
		if dep[v] > dep[w] {
			v, w = w, v
		}
		w = uptoDep(w, dep[v])
		if w == v {
			return v
		}
		for i := mx - 1; i >= 0; i-- {
			pv, pw := pa[v][i], pa[w][i]
			if pv != pw {
				v, w = pv, pw
			}
		}
		return pa[v][0]
	}

	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	type pair struct{ x, i int }
	qs := make([][]pair, n)
	Fscan(in, &q)
	ans := make([]int, q)
	for i := range ans {
		ans[i] = 1
		var v, w, x int
		Fscan(in, &v, &w, &x)
		v--
		w--
		lca := getLCA(v, w)
		qs[v] = append(qs[v], pair{x, i})
		qs[w] = append(qs[w], pair{x, i})
		qs[lca] = append(qs[lca], pair{-x, i})
	}

	mulCnt := map[int]map[int]int{}
	var dfs func(int, int)
	dfs = func(v, fa int) {
		for x := a[v]; x > 1; {
			p := lpf[x]
			mul := p
			for x /= p; x%p == 0; x /= p {
				mul *= p
			}
			if mulCnt[p] == nil {
				mulCnt[p] = map[int]int{}
			}
			mulCnt[p][mul]++
		}

		for _, q := range qs[v] {
			res := 1
			for x := abs(q.x); x > 1; {
				p := lpf[x]
				mul := p
				for x /= p; x%p == 0; x /= p {
					mul *= p
				}
				for m, c := range mulCnt[p] {
					res = res * pow(min(m, mul), c) % mod
				}
			}
			if q.x > 0 {
				ans[q.i] = ans[q.i] * res % mod
			} else {
				ans[q.i] = ans[q.i] * pow(res, mod-3) % mod * gcd(-q.x, a[v]) % mod
			}
		}

		for _, w := range g[v] {
			if w != fa {
				dfs(w, v)
			}
		}

		for x := a[v]; x > 1; {
			p := lpf[x]
			mul := p
			for x /= p; x%p == 0; x /= p {
				mul *= p
			}
			mulCnt[p][mul]--
		}
	}
	dfs(0, -1)

	for _, v := range ans {
		Fprintln(out, v)
	}
}

//func main() { cf986E(bufio.NewReader(os.Stdin), os.Stdout) }
