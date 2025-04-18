package main

import (
	"cmp"
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
type vec66 struct{ x, y int }

func (a vec66) sub(b vec66) vec66 { return vec66{a.x - b.x, a.y - b.y} }
func (a vec66) dot(b vec66) int   { return a.x*b.x + a.y*b.y }
func (a vec66) det(b vec66) int   { return a.x*b.y - a.y*b.x }

func cf1366F(in io.Reader, out io.Writer) {
	const mod = 1_000_000_007
	var n, m, k, ans int
	Fscan(in, &n, &m, &k)
	type nb struct{ to, wt int }
	g := make([][]nb, n)
	for range m {
		var v, w, wt int
		Fscan(in, &v, &w, &wt)
		v--
		w--
		g[v] = append(g[v], nb{w, wt})
		g[w] = append(g[w], nb{v, wt})
	}

	f := make([]int, n)
	for i := 1; i < n; i++ {
		f[i] = -1e18
	}
	for range m {
		nf := make([]int, n)
		for i := range nf {
			nf[i] = -1e18
		}
		for v, fv := range f {
			if fv < 0 {
				continue
			}
			for _, e := range g[v] {
				nf[e.to] = max(nf[e.to], fv+e.wt)
			}
		}
		f = nf
		ans += slices.Max(f)
	}

	a := make([]vec66, 0, n)
	for i, fv := range f {
		if fv < 0 {
			continue
		}
		mx := 0
		for _, e := range g[i] {
			mx = max(mx, e.wt)
		}
		a = append(a, vec66{mx, fv})
	}
	slices.SortFunc(a, func(a, b vec66) int { return cmp.Or(a.x-b.x, a.y-b.y) })
	q := a[:0]
	for _, v := range a {
		for len(q) > 1 && q[len(q)-1].sub(q[len(q)-2]).det(v.sub(q[len(q)-1])) >= 0 {
			q = q[:len(q)-1]
		}
		q = append(q, v)
	}
	if len(q) > 1 && q[0].x == q[1].x {
		q = q[1:]
	}

	k -= m
	i := 1
	for len(q) > 1 {
		nxt := (q[0].y-q[1].y)/(q[1].x-q[0].x) + 1
		if nxt > k {
			break
		}
		if nxt > i {
			ans = (ans + (i+nxt-1)*(nxt-i)/2%mod*q[0].x + (nxt-i)*q[0].y) % mod
			i = nxt
		}
		q = q[1:]
	}
	ans = (ans + (i+k)*(k-i+1)/2%mod*q[0].x + (k-i+1)*q[0].y) % mod
	Fprint(out, ans)
}

//func main() { cf1366F(bufio.NewReader(os.Stdin), os.Stdout) }
