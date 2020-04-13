package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1217D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, v, w int
	Fscan(in, &n, &m)
	g := make([][][2]int, n)
	for i := 0; i < m; i++ {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], [2]int{w, i})
	}

	hasCycle := false
	ans := make([]interface{}, m)
	vis := make([]int8, n)
	var f func(int)
	f = func(v int) {
		vis[v] = 1
		for _, e := range g[v] {
			w, i := e[0], e[1]
			if b := vis[w]; b == 0 {
				f(w)
				ans[i] = 1
			} else if b == 2 {
				ans[i] = 1
			} else {
				hasCycle = true
				ans[i] = 2
			}
		}
		vis[v] = 2
	}
	for i, b := range vis {
		if b == 0 {
			f(i)
		}
	}
	if hasCycle {
		Fprintln(out, 2)
	} else {
		Fprintln(out, 1)
	}
	Fprint(out, ans...)
}

//func main() { CF1217D(os.Stdin, os.Stdout) }
