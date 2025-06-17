package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf825G(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, q, op, v, w, last int
	Fscan(in, &n, &q)
	g := make([][]int, n)
	for range n - 1 {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}
	Fscan(in, &v, &v)
	mn := make([]int, n)
	var dfs func(int, int, int)
	dfs = func(v, fa, x int) {
		x = min(x, v)
		mn[v] = x
		for _, w := range g[v] {
			if w != fa {
				dfs(w, v, x)
			}
		}
	}
	dfs(v%n, -1, n)

	ans := v % n
	for range q - 1 {
		Fscan(in, &op, &v)
		v = (v + last) % n
		if op == 1 {
			ans = min(ans, mn[v])
		} else {
			last = min(ans, mn[v]) + 1
			Fprintln(out, last)
		}
	}
}

//func main() { cf825G(bufio.NewReader(os.Stdin), os.Stdout) }
