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

	var n, v, w, t int
	Fscan(in, &n)
	g := make([][]int, n)
	g[0] = append(g[0], -1)
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}
	ans := make([][2]int, n)
	var f func(int, int)
	f = func(v, fa int) {
		if len(g[v]) == 1 {
			t++
			ans[v] = [2]int{t, t}
			return
		}
		ans[v][0] = t + 1
		for _, w := range g[v] {
			if w != fa {
				f(w, v)
			}
		}
		ans[v][1] = t
	}
	f(0, -1)
	for _, p := range ans {
		Fprintln(out, p[0], p[1])
	}
}

func main() { run(os.Stdin, os.Stdout) }
