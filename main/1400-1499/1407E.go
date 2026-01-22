package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1407E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m int
	Fscan(in, &n, &m)
	type edge struct{ to, tp int }
	g := make([][]edge, n)
	for range m {
		var v, w, tp int
		Fscan(in, &v, &w, &tp)
		g[w-1] = append(g[w-1], edge{v - 1, tp})
	}

	dis := make([]int, n)
	ans := make([]int, n)
	for i := range n {
		dis[i] = -1
		ans[i] = -1
	}
	dis[n-1] = 0

	q := []int{n - 1}
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		for _, e := range g[v] {
			w, tp := e.to, e.tp
			if ans[w] != tp {
				ans[w] = tp ^ 1
			} else if dis[w] < 0 {
				dis[w] = dis[v] + 1
				q = append(q, w)
			}
		}
	}

	Fprintln(out, dis[0])
	for _, v := range ans {
		Fprint(out, max(v, 0))
	}
}

//func main() { cf1407E(bufio.NewReader(os.Stdin), os.Stdout) }
