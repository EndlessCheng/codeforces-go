package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
type fenwick27 []int

func (t fenwick27) add(i int) {
	for ; i < len(t); i += i & -i {
		t[i]++
	}
}

func (t fenwick27) pre(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res += t[i]
	}
	return
}

func (t fenwick27) query(l, r int) int {
	return t.pre(r) - t.pre(l-1)
}

func cf1827D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, size int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		g := make([][]int, n)
		for w := 1; w < n; w++ {
			var v int
			Fscan(in, &v)
			g[v-1] = append(g[v-1], w)
		}

		pa := make([][19]int, n)
		dep := make([]int, n)
		tin := make([]int, n)
		tout := make([]int, n)
		now := 0
		var dfs func(int, int)
		dfs = func(v, p int) {
			now++
			tin[v] = now
			pa[v][0] = p
			for _, w := range g[v] {
				if w != p {
					dep[w] = dep[v] + 1
					dfs(w, v)
				}
			}
			tout[v] = now
		}
		dfs(0, -1)

		mx := bits.Len(uint(n))
		for i := range mx - 1 {
			for v := range pa {
				p := pa[v][i]
				if p != -1 {
					pa[v][i+1] = pa[p][i]
				} else {
					pa[v][i+1] = -1
				}
			}
		}
		up := func(v, d int) int {
			for k := uint32(dep[v] - d); k > 0; k &= k - 1 {
				v = pa[v][bits.TrailingZeros32(k)]
			}
			return v
		}
		down := func(v, to int) int {
			if dep[to] <= dep[v] {
				return -1
			}
			to = up(to, dep[v]+1)
			if pa[to][0] == v {
				return to
			}
			return -1
		}

		t := make(fenwick27, n+1)
		t.add(tin[0])
		ct, maxSonSize := 0, 0
		for i := 1; i < n; i++ {
			nodes := i + 1
			t.add(tin[i])

			son := down(ct, i)
			if son >= 0 {
				size = t.query(tin[son], tout[son])
			} else {
				size = nodes - t.query(tin[ct], tout[ct])
			}
			maxSonSize = max(maxSonSize, size)

			if maxSonSize > nodes/2 {
				maxSonSize = nodes / 2
				if son >= 0 {
					ct = son
				} else {
					ct = pa[ct][0]
				}
			}

			Fprint(out, nodes-maxSonSize*2, " ")
		}
		Fprintln(out)
	}
}

//func main() { cf1827D(bufio.NewReader(os.Stdin), os.Stdout) }
