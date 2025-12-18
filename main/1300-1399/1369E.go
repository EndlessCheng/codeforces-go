package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1369E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m int
	Fscan(in, &n, &m)
	w := make([]int, n)
	for i := range w {
		Fscan(in, &w[i])
	}
	type pair struct{ to, i int }
	g := make([][]pair, n)
	deg := make([]int, n)
	for i := range m {
		var x, y int
		Fscan(in, &x, &y)
		x--
		y--
		g[x] = append(g[x], pair{y, i})
		g[y] = append(g[y], pair{x, i})
		deg[x]++
		deg[y]++
	}

	q := []int{}
	for i, d := range deg {
		if d <= w[i] {
			q = append(q, i)
		}
	}

	ans := []int{}
	vis := make([]bool, m)
	for len(q) > 0 {
		x := q[0]
		q = q[1:]
		for _, e := range g[x] {
			if vis[e.i] {
				continue
			}
			vis[e.i] = true
			ans = append(ans, e.i)
			y := e.to
			deg[y]--
			if deg[y] == w[y] {
				q = append(q, y)
			}
		}
	}

	if len(ans) < m {
		Fprint(out, "DEAD")
	} else {
		Fprintln(out, "ALIVE")
		for i := m - 1; i >= 0; i-- {
			Fprint(out, ans[i]+1, " ")
		}
	}
}

//func main() { cf1369E(bufio.NewReader(os.Stdin), os.Stdout) }
