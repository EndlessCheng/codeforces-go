package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, v, w int
	Fscan(in, &n, &m)
	mk := make([][2]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &v, &w)
		if v > 0 {
			mk[i][0] = v
		}
		if w > 0 {
			mk[i][1] = w
		}
	}
	rel := make([][2]int, m)
	for i := range rel {
		Fscan(in, &v, &w)
		w--
		rel[i] = [2]int{v, w}
		mk[v][w] = -mk[v][w]
	}
	g := make([][]int, n+1)
	for v, vs := range mk {
		for _, w := range vs {
			if w > 0 {
				g[v] = append(g[v], w)
				g[w] = append(g[w], v)
			}
		}
	}

	ans := make([]int, n+1)
	t := -1
	vis := make([]bool, n+1)
	var f func(int)
	f = func(v int) {
		ans[v] = t
		vis[v] = true
		for _, w := range g[v] {
			if !vis[w] {
				f(w)
			}
		}
	}
	f(1)
	for t = m - 1; t >= 0; t-- {
		r := rel[t]
		v := r[0]
		w := -mk[v][r[1]]
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
		// 一端未访问另一端访问的情况符合「松手」的描述
		if !vis[v] && vis[w] {
			f(v)
		} else if !vis[w] && vis[v] {
			f(w)
		}
	}
	for i := 1; i <= n; i++ {
		Fprintln(out, ans[i])
	}
}

func main() { run(os.Stdin, os.Stdout) }
