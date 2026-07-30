package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
func cf2236G(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, q, v, w, ans, t int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &q)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		g := make([][]int, n)
		for range n - 1 {
			Fscan(in, &v, &w)
			v--
			w--
			g[v] = append(g[v], w)
			g[w] = append(g[w], v)
		}

		mx := bits.Len(uint(n))
		dep := make([]int, n)
		pa := make([][17]int, n)
		type pair struct{ topDep, cnt int }
		data := make([][17]pair, n)
		lastNZ := make([]int, n)
		st := []int{}
		var dfs func(int, int, int, int, int)
		dfs = func(v, fa, or, top, last int) {
			dep[v] = len(st)
			st = append(st, v)
			for or&a[v] > 0 {
				or &^= a[st[top]]
				top++
			}
			or |= a[v]

			pa[v][0] = fa
			data[v][0] = pair{top, len(st) - top}
			for i := range mx - 1 {
				p := pa[v][i]
				if p != -1 {
					pa[v][i+1] = pa[p][i]
					d := data[p][i]
					data[v][i+1] = pair{d.topDep, data[v][i].cnt + d.cnt}
				} else {
					pa[v][i+1] = -1
				}
			}

			lastNZ[v] = last
			if a[v] > 0 {
				last = v
			}

			for _, w := range g[v] {
				if w != fa {
					dfs(w, v, or, top, last)
				}
			}
			st = st[:len(st)-1]
		}
		dfs(0, -1, 0, 0, -1)
		uptoDep := func(v, d int) int {
			for k := uint32(dep[v] - d); k > 0; k &= k - 1 {
				v = pa[v][bits.TrailingZeros32(k)]
			}
			return v
		}
		getLCA := func(v, w int) int {
			v = uptoDep(v, dep[w])
			if v == w {
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
		up := func(v, dLca int) (int, int) {
			tot := 0
			for i := mx - 1; i >= 0; i-- {
				w := pa[v][i]
				d := data[v][i]
				if w >= 0 && d.topDep > dLca {
					tot += d.cnt
					v = w
				}
			}
			sz := dep[v] - dLca + 1
			return v, tot + sz*(sz+1)/2
		}

		for range q {
			Fscan(in, &v, &w)
			v--
			w--
			if dep[v] < dep[w] {
				v, w = w, v
			}
			lca := getLCA(v, w)
			dLca := dep[lca]
			v, ans = up(v, dLca)
			if w == lca {
				Fprintln(out, ans)
				continue
			}
			w, t = up(w, dLca)
			ans += t - 1

			nodes := []int{}
			for x := w; x >= 0 && dep[x] > dLca; x = lastNZ[x] {
				nodes = append(nodes, x)
			}

			or := a[lca]
			for x := v; x >= 0 && dep[x] > dLca; x = lastNZ[x] {
				or |= a[x]
			}

			depV := dep[v]
			depW := dLca + 1
			for i := len(nodes) - 1; i >= 0; i-- {
				w := nodes[i]
				ans += (depV - dLca) * (dep[w] - depW)
				depW = dep[w]
				for v >= 0 && dep[v] > dLca && or&a[w] > 0 {
					or &^= a[v]
					depV = dep[v] - 1
					v = lastNZ[v]
				}
				or |= a[w]
				if i == 0 {
					ans += depV - dLca
				}
			}
			Fprintln(out, ans)
		}
	}
}

//func main() { cf2236G(bufio.NewReader(os.Stdin), os.Stdout) }
