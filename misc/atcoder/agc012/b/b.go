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
	var n, m, q int
	Fscan(in, &n, &m)
	g := make([][]int, n)
	for ; m > 0; m-- {
		var v, w int
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}
	Fscan(in, &q)
	qs := make([]struct{ v, d, c int }, q)
	for i := range qs {
		Fscan(in, &qs[i].v, &qs[i].d, &qs[i].c)
		qs[i].v--
	}

	ans := make([]int, n)
	dep := make([]int, n)
	for i := q - 1; i >= 0; i-- {
		q := qs[i]
		var dfs func(int, int)
		dfs = func(v, d int) {
			if d < 0 || ans[v] > 0 && d <= dep[v] {
				return
			}
			if ans[v] == 0 {
				ans[v] = q.c
			}
			dep[v] = d
			for _, w := range g[v] {
				dfs(w, d-1)
			}
		}
		dfs(q.v, q.d)
	}
	for _, c := range ans {
		Fprintln(out, c)
	}
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
