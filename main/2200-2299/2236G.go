package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"os"
)

// https://github.com/EndlessCheng
func cf2236G(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, q, v, w, ans, t int
	mem := [21]int{}
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &q)
		a := make([]int, n+1)
		for i := 1; i <= n; i++ {
			Fscan(in, &a[i])
		}
		g := make([][]int, n+1)
		for range n - 1 {
			Fscan(in, &v, &w)
			g[v] = append(g[v], w)
			g[w] = append(g[w], v)
		}

		mx := bits.Len(uint(n))
		pa := make([][17]int, n+1)
		dep := make([]int, n+1)
		topDep := make([]int, n+1)
		sum := make([]int, n+1)
		lastNZ := make([]int, n+1)
		st := []int{}
		var dfs func(int, int, int, int, int, int)
		dfs = func(v, fa, or, idx, topD, nz int) {
			pa[v][0] = fa
			for i := range mx - 1 {
				pa[v][i+1] = pa[pa[v][i]][i]
			}

			if a[v] > 0 {
				st = append(st, v)
				for or&a[v] > 0 {
					or &^= a[st[idx]]
					topD = dep[st[idx]] + 1
					idx++
				}
				or |= a[v]
			}
			topDep[v] = topD
			dep[v] = dep[fa] + 1
			sum[v] = sum[fa] + dep[v] - topD + 1

			lastNZ[v] = nz
			if a[v] > 0 {
				nz = v
			}

			for _, w := range g[v] {
				if w != fa {
					dfs(w, v, or, idx, topD, nz)
				}
			}
			if a[v] > 0 {
				st = st[:len(st)-1]
			}
		}
		dfs(1, 0, 0, 0, 0, 0)

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
			x := v
			for i := mx - 1; i >= 0; i-- {
				p := pa[x][i]
				if topDep[p] > dLca {
					x = p
				}
			}
			if topDep[x] > dLca {
				x = pa[x][0]
			}
			sz := dep[x] - dLca + 1
			return x, sum[v] - sum[x] + sz*(sz+1)/2
		}

		for range q {
			Fscan(in, &v, &w)
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

			nodes := mem[:0]
			for x := w; dep[x] > dLca; x = lastNZ[x] {
				nodes = append(nodes, x)
			}

			or := a[lca]
			for x := v; dep[x] > dLca; x = lastNZ[x] {
				or |= a[x]
			}

			depV := dep[v]
			depW := dLca + 1
			for i := len(nodes) - 1; i >= 0; i-- {
				w := nodes[i]
				ans += (depV - dLca) * (dep[w] - depW)
				depW = dep[w]
				for dep[v] > dLca && or&a[w] > 0 {
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

func main() { cf2236G(bufio.NewReader(os.Stdin), os.Stdout) }
