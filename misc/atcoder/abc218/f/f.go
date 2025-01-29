package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://github.com/EndlessCheng
func run(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m int
	Fscan(in, &n, &m)
	type edge struct{ to, i int }
	g := make([][]edge, n)
	for i := range m {
		var v, w int
		Fscan(in, &v, &w)
		g[v-1] = append(g[v-1], edge{w - 1, i})
	}

	from := make([]edge, n)
	dis := make([]int, n)
	bfs := func(ban int) int {
		for i := range dis {
			dis[i] = -1
		}
		dis[0] = 0
		q := []int{0}
		for len(q) > 0 {
			v := q[0]
			q = q[1:]
			for _, e := range g[v] {
				w := e.to
				if e.i != ban && dis[w] < 0 {
					dis[w] = dis[v] + 1
					if ban < 0 {
						from[w] = edge{v, e.i}
					}
					q = append(q, w)
				}
			}
		}
		return dis[n-1]
	}

	res := bfs(-1)
	ans := make([]int, m)
	for i := range ans {
		ans[i] = res
	}
	if res != -1 {
		for v := n - 1; v != 0; v = from[v].to {
			ans[from[v].i] = bfs(from[v].i)
		}
	}
	for _, v := range ans {
		Fprintln(out, v)
	}
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
