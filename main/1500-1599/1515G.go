package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1515G(in io.Reader, _w io.Writer) {
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m, now, q int
	Fscan(in, &n, &m)
	type nb struct{ to, wt int }
	g := make([][]nb, n)
	for range m {
		var v, w, wt int
		Fscan(in, &v, &w, &wt)
		g[v-1] = append(g[v-1], nb{w - 1, wt})
	}

	sccGCD := make([]int, n)
	dis := make([]int, n)
	dfn := make([]int, n)
	st := []int{}
	var tarjan func(int) int
	tarjan = func(v int) int {
		now++
		dfn[v] = now
		lowV := now
		st = append(st, v)
		for _, e := range g[v] {
			w := e.to
			if dfn[w] == 0 {
				dis[w] = dis[v] + e.wt
				lowV = min(lowV, tarjan(w))
			}
			if dfn[w] != 1e9 {
				sccGCD[v] = gcd(gcd(sccGCD[v], sccGCD[w]), dis[v]-dis[w]+e.wt)
				lowV = min(lowV, dfn[w])
			}
		}
		if dfn[v] == lowV {
			for {
				w := st[len(st)-1]
				st = st[:len(st)-1]
				sccGCD[w] = sccGCD[v]
				dfn[w] = 1e9
				if w == v {
					break
				}
			}
		}
		return lowV
	}
	for i, t := range dfn {
		if t == 0 {
			tarjan(i)
		}
	}

	Fscan(in, &q)
	for range q {
		var v, s, t int
		Fscan(in, &v, &s, &t)
		if s%gcd(sccGCD[v-1], t) == 0 {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}
}

//func main() { cf1515G(bufio.NewReader(os.Stdin), os.Stdout) }
