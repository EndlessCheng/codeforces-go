package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf645D(in io.Reader, out io.Writer) {
	var n, m, ans int
	Fscan(in, &n, &m)
	type edge struct{ to, i int }
	g := make([][]edge, n)
	deg := make([]int, n)
	for i := range m {
		var v, w int
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], edge{w, i})
		deg[w]++
	}

	q := []int{}
	for i, d := range deg {
		if d == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		if len(q) > 1 {
			Fprint(out, -1)
			return
		}
		v := q[0]
		q = q[1:]
		for _, e := range g[v] {
			w := e.to
			deg[w]--
			if deg[w] == 0 {
				ans = max(ans, e.i)
				q = append(q, w)
			}
		}
	}
	Fprint(out, ans+1)
}

//func main() { cf645D(bufio.NewReader(os.Stdin), os.Stdout) }
