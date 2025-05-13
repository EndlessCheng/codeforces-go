package main

import (
	. "fmt"
	"io"
	"sort"
)

// https://github.com/EndlessCheng
func cf645D(in io.Reader, out io.Writer) {
	var n, m int
	Fscan(in, &n, &m)
	type edge struct{ to, i int }
	g := make([][]edge, n)
	for i := range m {
		var v, w int
		Fscan(in, &v, &w)
		g[v-1] = append(g[v-1], edge{w - 1, i})
	}

	deg := make([]int, n)
	l, r := n-1, m+1
	ans := l + sort.Search(r-l, func(m int) bool {
		m += l
		clear(deg)
		for _, es := range g {
			for _, e := range es {
				if e.i < m {
					deg[e.to]++
				}
			}
		}
		q := []int{}
		for i, d := range deg {
			if d == 0 {
				q = append(q, i)
			}
		}
		for len(q) > 0 {
			if len(q) > 1 {
				return false
			}
			v := q[0]
			q = q[1:]
			for _, e := range g[v] {
				if e.i >= m {
					continue
				}
				w := e.to
				deg[w]--
				if deg[w] == 0 {
					q = append(q, w)
				}
			}
		}
		return true
	})
	if ans > m {
		ans = -1
	}
	Fprint(out, ans)
}

//func main() { cf645D(bufio.NewReader(os.Stdin), os.Stdout) }
